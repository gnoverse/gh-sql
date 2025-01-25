package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gnoverse/gh-sql/pkg/model"
)

// PullRequest holds the schema definition for the PullRequest entity.
type PullRequest struct {
	ent.Schema
}

// Fields of the PullRequest.
func (PullRequest) Fields() []ent.Field {
	return []ent.Field{
		field.String("url"),
		field.Int64("id"),
		field.String("node_id"),
		field.String("html_url"),
		field.String("diff_url"),
		field.String("patch_url"),
		field.String("issue_url"),
		field.String("commits_url"),
		field.String("review_comments_url"),
		field.String("review_comment_url"),
		field.String("comments_url"),
		field.String("statuses_url"),
		field.Int64("number").
			Comment("Number uniquely identifying the pull request within its repository."),
		field.Enum("state").
			Values("open", "closed").
			Comment("State of this Pull Request. Either `open` or `closed`."),
		field.Bool("locked"),
		field.String("title").
			Comment("The title of the pull request."),
		field.String("body").
			Optional(),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Time("closed_at").
			Optional(),
		field.Time("merged_at").
			Optional(),
		field.String("merge_commit_sha").
			Optional(),
		// These fields are only available in the full pull request object.
		// Seeing as a lot of them are redundant, we skip these so we can
		// directly use the /pulls endpoint, saving a lot of requests.
		// field.String("active_lock_reason").
		// 	Optional().
		// 	Nillable(),
		// field.Bool("merged"),
		// field.Bool("mergeable").
		// 	Optional(),
		// field.String("mergeable_state"),
		// field.Int64("comments"),
		// field.Int64("review_comments"),
		// field.Bool("maintainer_can_modify").
		// 	Comment("Indicates whether maintainers can modify the pull request."),
		// field.Int64("commits"),
		// field.Int64("additions"),
		// field.Int64("deletions"),
		// field.Int64("changed_files"),
		field.JSON("head", model.PRBranch{}),
		field.JSON("base", model.PRBranch{}),
		field.Bool("draft").
			Optional().
			Nillable(),
		field.Enum("author_association").GoType(model.AuthorAssociation("")),
	}
}

// Edges of the PullRequest.
func (PullRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("repository", Repository.Type).
			Ref("pull_requests").
			Unique().
			Required(),
		edge.From("issue", Issue.Type).
			Ref("pull_request").
			Unique().
			Required(),
		edge.From("user", User.Type).
			Ref("prs_created").
			Unique(),
		edge.To("assignees", User.Type),
		edge.To("requested_reviewers", User.Type),
		// edge: requested_teams (#/components/schemas/team-simple)
		// edge: labels
		// edge: milestone (#/components/schemas/nullable-milestone)
		// edge: auto_merge (#/components/schemas/auto-merge)
	}
}
