// Package schema specifies the ent schema types for gh-sql's database.
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gnoverse/gh-sql/pkg/model"
)

// Repository holds the schema definition for the Repository entity.
type Repository struct {
	ent.Schema
}

// Fields of the Repository.
func (Repository) Fields() []ent.Field {
	return []ent.Field{
		// Order should match the OpenAPI description (not what is actually returned by the API),
		// see misc/schema2schema for help generating a schema like this.

		field.Int64("id"),
		field.String("node_id"),
		field.String("name"),
		field.String("full_name").Unique(),
		field.Bool("private"),
		field.String("html_url"),
		field.String("description").
			Optional(),
		field.Bool("fork"),
		field.String("url"),
		field.String("archive_url"),
		field.String("assignees_url"),
		field.String("blobs_url"),
		field.String("branches_url"),
		field.String("collaborators_url"),
		field.String("comments_url"),
		field.String("commits_url"),
		field.String("compare_url"),
		field.String("contents_url"),
		field.String("contributors_url"),
		field.String("deployments_url"),
		field.String("downloads_url"),
		field.String("events_url"),
		field.String("forks_url"),
		field.String("git_commits_url"),
		field.String("git_refs_url"),
		field.String("git_tags_url"),
		field.String("git_url"),
		field.String("issue_comment_url"),
		field.String("issue_events_url"),
		field.String("issues_url"),
		field.String("keys_url"),
		field.String("labels_url"),
		field.String("languages_url"),
		field.String("merges_url"),
		field.String("milestones_url"),
		field.String("notifications_url"),
		field.String("pulls_url"),
		field.String("releases_url"),
		field.String("ssh_url"),
		field.String("stargazers_url"),
		field.String("statuses_url"),
		field.String("subscribers_url"),
		field.String("subscription_url"),
		field.String("tags_url"),
		field.String("teams_url"),
		field.String("trees_url"),
		field.String("clone_url"),
		field.String("mirror_url").
			Optional().
			Nillable(),
		field.String("hooks_url"),
		field.String("svn_url"),
		field.String("homepage").
			Optional(),
		field.String("language").
			Optional(),
		field.Int64("forks_count"),
		field.Int64("stargazers_count"),
		field.Int64("watchers_count"),
		field.Int64("size").
			Comment("The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0."),
		field.String("default_branch"),
		field.Int64("open_issues_count"),
		field.Bool("is_template"),
		field.Strings("topics"),
		field.Bool("has_issues_enabled").StructTag(`json:"has_issues"`),
		field.Bool("has_projects"),
		field.Bool("has_wiki"),
		field.Bool("has_pages"),
		field.Bool("has_downloads"),
		field.Bool("has_discussions"),
		field.Bool("archived"),
		field.Bool("disabled").
			Comment("Returns whether or not this repository disabled."),
		field.Enum("visibility").
			Nillable().
			Optional().
			Values("public", "private", "internal"),
		field.Time("pushed_at"),
		field.Time("created_at"),
		field.Time("updated_at"),
		// ignored: permissions: not required
		// ignored: allow_rebase_merge: not required
		// ignored: temp_clone_token: not required
		// ignored: allow_squash_merge: not required
		// ignored: allow_auto_merge: not required
		// ignored: delete_branch_on_merge: not required
		// ignored: allow_merge_commit: not required
		// ignored: allow_update_branch: not required
		// ignored: use_squash_pr_title_as_default: not required
		// ignored: squash_merge_commit_title: not required
		// ignored: squash_merge_commit_message: not required
		// ignored: merge_commit_title: not required
		// ignored: merge_commit_message: not required
		// ignored: allow_forking: not required
		// ignored: web_commit_signoff_required: not required
		field.Int64("subscribers_count"),
		field.Int64("network_count"),
		field.Int64("forks"),
		// ignored: master_branch: not required
		field.Int64("open_issues"),
		field.Int64("watchers"),
		field.JSON("license", &model.License{}).Optional(),
		// ignored: anonymous_access_enabled: not required
		// ignored: custom_properties: not required
	}
}

// Edges of the Repository.
func (Repository) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("repositories").
			Unique(),
		edge.To("issues", Issue.Type),
		edge.To("pull_requests", PullRequest.Type),
		// edge: template_repository (#/components/schemas/nullable-repository)
		// edge: organization (#/components/schemas/nullable-simple-user)
		// edge: parent (#/components/schemas/repository)
		// edge: source (#/components/schemas/repository)
		// edge: code_of_conduct (#/components/schemas/code-of-conduct-simple)
		// edge: security_and_analysis (#/components/schemas/security-and-analysis)
	}
}
