package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("login").Unique(),
		field.Int64("id"),
		field.String("node_id"),
		field.String("avatar_url"),
		field.String("gravatar_id").
			Optional(),
		field.String("url"),
		field.String("html_url"),
		field.String("followers_url"),
		field.String("following_url"),
		field.String("gists_url"),
		field.String("starred_url"),
		field.String("subscriptions_url"),
		field.String("organizations_url"),
		field.String("repos_url"),
		field.String("events_url"),
		field.String("received_events_url"),
		field.String("type"),
		field.Bool("site_admin"),
		field.String("name").
			Optional(),
		field.String("company").
			Optional(),
		field.String("blog").
			Optional(),
		field.String("location").
			Optional(),
		field.String("email").
			Optional(),
		field.Bool("hireable").
			Optional(),
		field.String("bio").
			Optional(),
		field.Int64("public_repos"),
		field.Int64("public_gists"),
		field.Int64("followers"),
		field.Int64("following"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("repositories", Repository.Type),
		edge.To("issues_created", Issue.Type),
		edge.To("comments_created", IssueComment.Type),
		edge.From("issues_assigned", Issue.Type).
			Ref("assignees"),
		edge.From("timeline_events_created", TimelineEvent.Type).
			Ref("actor"),
	}
}
