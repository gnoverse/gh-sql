package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Issue holds the schema definition for the Issue entity.
type Issue struct {
	ent.Schema
}

// Fields of the Issue.
func (Issue) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("node_id"),
		field.String("url").
			Comment("URL for the issue"),
		field.String("repository_url"),
		field.String("labels_url"),
		field.String("comments_url"),
		field.String("events_url"),
		field.String("html_url"),
		field.Int64("number").
			Comment("Number uniquely identifying the issue within its repository"),
		field.String("state").
			Comment("State of the issue; either 'open' or 'closed'"),
		field.Enum("state_reason").
			Optional().
			Nillable().
			Values("completed", "reopened", "not_planned"),
		field.String("title").
			Comment("Title of the issue"),
		field.String("body").
			Optional().
			Nillable(),
		field.Bool("locked"),
		field.String("active_lock_reason").
			Optional().
			Nillable(),
		// field.Int64("comments"), removed in favour of comments edge
		field.Time("closed_at").
			Optional().
			Nillable(),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Bool("draft"),
		// TODO: unify with IssueComment
		field.Enum("author_association").Values(
			"COLLABORATOR",
			"CONTRIBUTOR",
			"FIRST_TIMER",
			"FIRST_TIME_CONTRIBUTOR",
			"MANNEQUIN",
			"MEMBER",
			"NONE",
			"OWNER",
		),
		// TODO: better type
		field.JSON("reactions", map[string]any{}),
	}
}

// Edges of the Issue.
func (Issue) Edges() []ent.Edge {
	// edge: labels
	// edge: pull_request
	// edge: milestone (#/components/schemas/nullable-milestone)
	// edge: performed_via_github_app (#/components/schemas/nullable-integration)
	return []ent.Edge{
		edge.From("repository", Repository.Type).
			Ref("issues").
			Unique().
			Required(),
		edge.From("user", User.Type).
			Ref("issues_created").
			Unique(),
		edge.To("assignees", User.Type),
		edge.To("closed_by", User.Type).
			Unique(),
		edge.To("comments", IssueComment.Type),
	}
}
