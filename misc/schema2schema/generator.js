// Script to help generate an ent schema from the GitHub
// graphql schema.
// Example:
// node generator.js components.schemas.full-repository
// node generator.js components.schemas.full-repository 1>/dev/null # filters for errors

import { promises as fs } from "fs";

const rawSchemaURL =
  "https://github.com/github/rest-api-description/raw/main/descriptions/api.github.com/api.github.com.2022-11-28.json";

if (process.argv.length < 3) {
  console.log(
    `usage: ${process.argv.slice(0, 2).join(" ")} <path.to.object.schema>`
  );
  process.exit(1);
}

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
const path = process.argv[2].split(".");
for (const el of path) {
  sel = sel[el];
}

const typeMap = {
  string: "String",
  boolean: "Bool",
  integer: "Int",
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
    (k) => !handledPropKeys.includes(k)
  );
  if (unhandled.length > 0) {
    console.error(
      `// XXX: ${key}: unhandled properties: ${unhandled.join(" ")}`
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
