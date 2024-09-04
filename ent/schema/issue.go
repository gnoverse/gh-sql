package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gnolang/gh-sql/pkg/model"
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
			GoType(model.StateReason("")),
		field.String("title").
			Comment("Title of the issue"),
		field.String("body").
			Optional().
			Nillable(),
		field.Bool("locked"),
		field.String("active_lock_reason").
			Optional().
			Nillable(),
		field.Int64("comments_count").StructTag(`json:"comments"`),
		field.Time("closed_at").
			Optional().
			Nillable(),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Bool("draft"),
		field.Enum("author_association").GoType(model.AuthorAssociation("")),
		field.JSON("reactions", model.ReactionRollup(nil)),
	}
}

// Edges of the Issue.
func (Issue) Edges() []ent.Edge {
	// edge: labels
	// edge: pull_request
	// edge: milestone (#/components/schemas/nullable-milestone)
	// edge: closed_by (#/components/schemas/nullable-user) (redundant with timeline)
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
		edge.To("comments", IssueComment.Type),
		edge.To("timeline", TimelineEvent.Type),
	}
}
