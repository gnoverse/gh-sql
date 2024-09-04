package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gnolang/gh-sql/pkg/model"
)

// IssueComment holds the schema definition for the IssueComment entity.
type IssueComment struct {
	ent.Schema
}

// Fields of the IssueComment.
func (IssueComment) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Comment("Unique identifier of the issue comment"),
		field.String("node_id"),
		field.String("url").
			Comment("URL for the issue comment"),
		field.String("body"),
		field.String("html_url"),
		field.String("created_at"),
		field.String("updated_at"),
		field.String("issue_url"),
		field.Enum("author_association").GoType(model.AuthorAssociation("")),
		// TODO: better type
		field.JSON("reactions", map[string]any{}),
	}
}

// Edges of the IssueComment.
func (IssueComment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("issue", Issue.Type).
			Ref("comments").
			Unique(),
		edge.From("user", User.Type).
			Ref("comments_created").
			Unique(),
	}
}
