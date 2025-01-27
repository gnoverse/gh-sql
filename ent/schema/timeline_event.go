package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gnoverse/gh-sql/pkg/model"
)

// TimelineEvent holds the schema definition for the TimelineEvent entity.
type TimelineEvent struct {
	ent.Schema
}

// Fields of the TimelineEvent.
func (TimelineEvent) Fields() []ent.Field {
	return []ent.Field{
		// Not all events contain an ID or node_id. This id serves to ensure
		// "insertion order" correctness.
		field.Int64("id").
			StructTag(`json:"-"`).
			Annotations(model.NoCopyTemplate{}),
		field.Int64("numeric_id").StructTag(`json:"id"`),
		field.String("node_id"),
		field.String("url"),
		field.String("event"),
		field.String("commit_id").
			Optional().
			Nillable(),
		field.String("commit_url").
			Optional().
			Nillable(),
		field.Time("created_at"),
		field.JSON("data", model.TimelineEventWrapper{}).StructTag(`json:"-"`),
	}
}

// Edges of the TimelineEvent.
func (TimelineEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("actor", User.Type).
			Unique(),
		edge.From("issue", Issue.Type).
			Ref("timeline").
			Unique(),
	}
}
