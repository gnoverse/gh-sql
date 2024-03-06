package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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

		field.Int("id"),
		field.String("node_id"),
		field.String("name"),
		field.String("full_name"),
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
			Optional(),
		field.String("hooks_url"),
		field.String("svn_url"),
		field.String("homepage").
			Optional(),
		field.String("language").
			Optional(),
		field.Int("forks_count"),
		field.Int("stargazers_count"),
		field.Int("watchers_count"),
		field.Int("size").
			Comment("The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0."),
		field.String("default_branch"),
		field.Int("open_issues_count"),
		field.Bool("is_template"),
		field.Strings("topics"),
		field.Bool("has_issues"),
		field.Bool("has_projects"),
		field.Bool("has_wiki"),
		field.Bool("has_pages"),
		field.Bool("has_downloads"),
		field.Bool("has_discussions"),
		field.Bool("archived"),
		field.Bool("disabled").
			Comment("Returns whether or not this repository disabled."),
		field.Enum("visibility").Values("public", "private", "internal"),
		field.String("pushed_at"),
		field.String("created_at"),
		field.String("updated_at"),
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
		field.Int("subscribers_count"),
		field.Int("network_count"),
		field.Int("forks"),
		// ignored: master_branch: not required
		field.Int("open_issues"),
		field.Int("watchers"),
		// ignored: anonymous_access_enabled: not required
		// ignored: custom_properties: not required
	}
}

// Edges of the Repository.
func (Repository) Edges() []ent.Edge {
	return []ent.Edge{
		// edge: owner (#/components/schemas/simple-user)
		// edge: template_repository (#/components/schemas/nullable-repository)
		// edge: license (#/components/schemas/nullable-license-simple)
		// edge: organization (#/components/schemas/nullable-simple-user)
		// edge: parent (#/components/schemas/repository)
		// edge: source (#/components/schemas/repository)
		// edge: code_of_conduct (#/components/schemas/code-of-conduct-simple)
		// edge: security_and_analysis (#/components/schemas/security-and-analysis)
	}
}
