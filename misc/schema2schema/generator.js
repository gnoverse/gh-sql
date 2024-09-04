// Script to help generate an ent schema from the GitHub
// OpenAPI schema.
// Example:
// node generator.js components.schemas.full-repository
// node generator.js components.schemas.full-repository 1>/dev/null # filters for errors

import { promises as fs } from "fs";

// FLAG PARSING
// ----------------------------------------------------------------------------
const OutputTargets = {
  Ent: "ent",
  EventStruct: "event",
};

let flagOutputTarget = OutputTargets.Ent;

const flagHandlers = {
  target(outputTarget) {
    if (!Object.values(OutputTargets).includes(outputTarget)) {
      console.error(`invalid target: ${outputTarget}`);
      process.exit(1);
    }
    flagOutputTarget = outputTarget;
  },
};
let args = process.argv.slice(2);
while (args.length > 0 && args[0].startsWith("--")) {
  // pop current arg.
  const flag = args[0];
  args = args.slice(1);
  if (flag == "--") {
    break;
  }
  const handler = flagHandlers[flag.slice(2)];
  if (typeof handler == "undefined") {
    console.error(`invalid flag: ${flag}`);
    process.exit(1);
  }
  let arg;
  if (args.length > 0) {
    arg = args[0];
    args = args.slice(1);
  }
  handler(arg);
}

if (args.length < 1) {
  console.log(
    `usage: ${process.argv.slice(0, 2).join(" ")} <path.to.object.schema>`,
  );
  process.exit(1);
}

// SCHEMA DOWNLOAD, FIND SELECTED KEY
// ----------------------------------------------------------------------------
const rawSchemaURL =
  "https://github.com/github/rest-api-description/raw/main/descriptions/api.github.com/api.github.com.2022-11-28.json";

// Download schema.json if it doesn't exist.
try {
  fs.stat("schema.json");
} catch (e) {
  const resp = await fetch(rawSchemaURL);
  const data = await resp.blob();
  fs.writeFile("schema.json", data.stream());
}

const obj = JSON.parse((await fs.readFile("schema.json")).toString("utf8"));

let sel = obj;
const path = args[0].split(".");
for (const el of path) {
  sel = sel[el];
}

// TARGETS
const forEnt = () => {
  const typeMap = {
    string: "String",
    boolean: "Bool",
    integer: "Int64",
    array: "Strings",
  };
  const handledPropKeys = [
    "type",
    "default",
    "example",
    "nullable",
    "items",
    "format", // TODO
    "description",
    "enum",
  ];

  const edges = [];

  for (const key of Object.keys(sel.properties)) {
    const prop = sel.properties[key];
    if (prop.$ref) {
      edges.push({ key: key, prop: prop });
      continue;
    }
    if (!sel.required.includes(key)) {
      console.error(`// XXX: ${key}: not required`);
      continue;
    }
    let dstType = typeMap[prop.type];
    if (!dstType) {
      console.error(`// XXX: ${key}: unhandled type: ${prop.type}`);
      continue;
    }
    if (dstType == "Strings") {
      console.error(`// XXX: ${key}: double-check the type`);
    }
    let enumValues = "";
    if (prop.enum) {
      if (dstType != "String") {
        console.error(`// XXX: unexpected enum type: ${dstType}`);
      } else {
        dstType = "Enum";
        enumValues = JSON.stringify(prop.enum);
        enumValues = "(" + enumValues.slice(1, -1) + ")";
      }
    }
    const unhandled = Object.keys(prop).filter(
      (k) => !handledPropKeys.includes(k),
    );
    if (unhandled.length > 0) {
      console.error(
        `// XXX: ${key}: unhandled properties: ${unhandled.join(" ")}`,
      );
    }
    const parts = [
      `field.${dstType}(${JSON.stringify(key)})`,
      prop.nullable ? "Optional()" : "",
      enumValues ? `Values${enumValues}` : "",
      prop.description ? `Comment(${JSON.stringify(prop.description)})` : "",
    ].filter((v) => v != "");
    console.log(parts.join(".\n\t") + ",");

    // TODO: filter out regex.
  }

  for (const edge of edges) {
    console.log(`// edge: ${edge.key} (${edge.prop.$ref})`);
  }
};

const snakeToPascalCase = (str, separator = "_") => {
  return str
    .split(separator)
    .map((word) => word.charAt(0).toUpperCase() + word.slice(1).toLowerCase())
    .join("");
};

function goType(prop) {
  if (prop.$ref !== undefined) {
    return snakeToPascalCase(prop.$ref.split("/").at(-1), "-");
  }
  switch (prop.type) {
    case "string":
      if (prop.format == "date-time") {
        return "time.Time";
      }
      return "string";
    case "boolean":
      return "bool";
    case "integer":
      return "int64";
    case "object":
      const sfields = toStructFields(Object.entries(prop.properties));
      return `struct {\n  ${sfields.join("\n  ")}\n}`;
    case "array":
      return `[]` + goType(prop.items);
  }
  console.error("can't convert type:", prop);
}
function toStructFields(entries) {
  return entries.map(
    (prop) =>
      `${snakeToPascalCase(prop[0])} ${goType(prop[1])} \`json:"${prop[0]}"\``,
  );
}

const forEvent = () => {
  const lastPath = path.at(-1);
  const pascalCaseType = snakeToPascalCase(lastPath, "-");
  const defaultProps = [
    "event",
    "actor",
    "id",
    "node_id",
    "url",
    "commit_id",
    "commit_url",
    "created_at",
    "performed_via_github_app",
  ];
  const props = Object.entries(sel.properties)
    // ignore default props
    .filter((prop) => !defaultProps.includes(prop[0]))
    // create Go struct field
    .map(
      (prop) =>
        `${snakeToPascalCase(prop[0])} ${goType(prop[1])} \`json:"${prop[0]}"\``,
    );
  const result = `type ${pascalCaseType} struct {
  timelineEvent

  ${props.join("\n  ")}
}

func (${pascalCaseType}) Name() string { return "CHANGEME" }`;
  // the openapi schema doesn't seem to have the event name, so we'll have to fill it manually.
  console.log(result);
};

switch (flagOutputTarget) {
  case OutputTargets.Ent:
    forEnt();
    break;
  case OutputTargets.EventStruct:
    forEvent();
    break;
}
