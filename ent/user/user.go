// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLogin holds the string denoting the login field in the database.
	FieldLogin = "login"
	// FieldNodeID holds the string denoting the node_id field in the database.
	FieldNodeID = "node_id"
	// FieldAvatarURL holds the string denoting the avatar_url field in the database.
	FieldAvatarURL = "avatar_url"
	// FieldGravatarID holds the string denoting the gravatar_id field in the database.
	FieldGravatarID = "gravatar_id"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldHTMLURL holds the string denoting the html_url field in the database.
	FieldHTMLURL = "html_url"
	// FieldFollowersURL holds the string denoting the followers_url field in the database.
	FieldFollowersURL = "followers_url"
	// FieldFollowingURL holds the string denoting the following_url field in the database.
	FieldFollowingURL = "following_url"
	// FieldGistsURL holds the string denoting the gists_url field in the database.
	FieldGistsURL = "gists_url"
	// FieldStarredURL holds the string denoting the starred_url field in the database.
	FieldStarredURL = "starred_url"
	// FieldSubscriptionsURL holds the string denoting the subscriptions_url field in the database.
	FieldSubscriptionsURL = "subscriptions_url"
	// FieldOrganizationsURL holds the string denoting the organizations_url field in the database.
	FieldOrganizationsURL = "organizations_url"
	// FieldReposURL holds the string denoting the repos_url field in the database.
	FieldReposURL = "repos_url"
	// FieldEventsURL holds the string denoting the events_url field in the database.
	FieldEventsURL = "events_url"
	// FieldReceivedEventsURL holds the string denoting the received_events_url field in the database.
	FieldReceivedEventsURL = "received_events_url"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldSiteAdmin holds the string denoting the site_admin field in the database.
	FieldSiteAdmin = "site_admin"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCompany holds the string denoting the company field in the database.
	FieldCompany = "company"
	// FieldBlog holds the string denoting the blog field in the database.
	FieldBlog = "blog"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldHireable holds the string denoting the hireable field in the database.
	FieldHireable = "hireable"
	// FieldBio holds the string denoting the bio field in the database.
	FieldBio = "bio"
	// FieldPublicRepos holds the string denoting the public_repos field in the database.
	FieldPublicRepos = "public_repos"
	// FieldPublicGists holds the string denoting the public_gists field in the database.
	FieldPublicGists = "public_gists"
	// FieldFollowers holds the string denoting the followers field in the database.
	FieldFollowers = "followers"
	// FieldFollowing holds the string denoting the following field in the database.
	FieldFollowing = "following"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeRepositories holds the string denoting the repositories edge name in mutations.
	EdgeRepositories = "repositories"
	// EdgeIssuesCreated holds the string denoting the issues_created edge name in mutations.
	EdgeIssuesCreated = "issues_created"
	// EdgeCommentsCreated holds the string denoting the comments_created edge name in mutations.
	EdgeCommentsCreated = "comments_created"
	// EdgeIssuesAssigned holds the string denoting the issues_assigned edge name in mutations.
	EdgeIssuesAssigned = "issues_assigned"
	// EdgeIssuesClosed holds the string denoting the issues_closed edge name in mutations.
	EdgeIssuesClosed = "issues_closed"
	// Table holds the table name of the user in the database.
	Table = "users"
	// RepositoriesTable is the table that holds the repositories relation/edge.
	RepositoriesTable = "repositories"
	// RepositoriesInverseTable is the table name for the Repository entity.
	// It exists in this package in order to avoid circular dependency with the "repository" package.
	RepositoriesInverseTable = "repositories"
	// RepositoriesColumn is the table column denoting the repositories relation/edge.
	RepositoriesColumn = "user_repositories"
	// IssuesCreatedTable is the table that holds the issues_created relation/edge.
	IssuesCreatedTable = "issues"
	// IssuesCreatedInverseTable is the table name for the Issue entity.
	// It exists in this package in order to avoid circular dependency with the "issue" package.
	IssuesCreatedInverseTable = "issues"
	// IssuesCreatedColumn is the table column denoting the issues_created relation/edge.
	IssuesCreatedColumn = "user_issues_created"
	// CommentsCreatedTable is the table that holds the comments_created relation/edge.
	CommentsCreatedTable = "issue_comments"
	// CommentsCreatedInverseTable is the table name for the IssueComment entity.
	// It exists in this package in order to avoid circular dependency with the "issuecomment" package.
	CommentsCreatedInverseTable = "issue_comments"
	// CommentsCreatedColumn is the table column denoting the comments_created relation/edge.
	CommentsCreatedColumn = "user_comments_created"
	// IssuesAssignedTable is the table that holds the issues_assigned relation/edge. The primary key declared below.
	IssuesAssignedTable = "issue_assignees"
	// IssuesAssignedInverseTable is the table name for the Issue entity.
	// It exists in this package in order to avoid circular dependency with the "issue" package.
	IssuesAssignedInverseTable = "issues"
	// IssuesClosedTable is the table that holds the issues_closed relation/edge.
	IssuesClosedTable = "issues"
	// IssuesClosedInverseTable is the table name for the Issue entity.
	// It exists in this package in order to avoid circular dependency with the "issue" package.
	IssuesClosedInverseTable = "issues"
	// IssuesClosedColumn is the table column denoting the issues_closed relation/edge.
	IssuesClosedColumn = "issue_closed_by"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldLogin,
	FieldNodeID,
	FieldAvatarURL,
	FieldGravatarID,
	FieldURL,
	FieldHTMLURL,
	FieldFollowersURL,
	FieldFollowingURL,
	FieldGistsURL,
	FieldStarredURL,
	FieldSubscriptionsURL,
	FieldOrganizationsURL,
	FieldReposURL,
	FieldEventsURL,
	FieldReceivedEventsURL,
	FieldType,
	FieldSiteAdmin,
	FieldName,
	FieldCompany,
	FieldBlog,
	FieldLocation,
	FieldEmail,
	FieldHireable,
	FieldBio,
	FieldPublicRepos,
	FieldPublicGists,
	FieldFollowers,
	FieldFollowing,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// IssuesAssignedPrimaryKey and IssuesAssignedColumn2 are the table columns denoting the
	// primary key for the issues_assigned relation (M2M).
	IssuesAssignedPrimaryKey = []string{"issue_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByLogin orders the results by the login field.
func ByLogin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLogin, opts...).ToFunc()
}

// ByNodeID orders the results by the node_id field.
func ByNodeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNodeID, opts...).ToFunc()
}

// ByAvatarURL orders the results by the avatar_url field.
func ByAvatarURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatarURL, opts...).ToFunc()
}

// ByGravatarID orders the results by the gravatar_id field.
func ByGravatarID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGravatarID, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByHTMLURL orders the results by the html_url field.
func ByHTMLURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHTMLURL, opts...).ToFunc()
}

// ByFollowersURL orders the results by the followers_url field.
func ByFollowersURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFollowersURL, opts...).ToFunc()
}

// ByFollowingURL orders the results by the following_url field.
func ByFollowingURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFollowingURL, opts...).ToFunc()
}

// ByGistsURL orders the results by the gists_url field.
func ByGistsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGistsURL, opts...).ToFunc()
}

// ByStarredURL orders the results by the starred_url field.
func ByStarredURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStarredURL, opts...).ToFunc()
}

// BySubscriptionsURL orders the results by the subscriptions_url field.
func BySubscriptionsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubscriptionsURL, opts...).ToFunc()
}

// ByOrganizationsURL orders the results by the organizations_url field.
func ByOrganizationsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrganizationsURL, opts...).ToFunc()
}

// ByReposURL orders the results by the repos_url field.
func ByReposURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReposURL, opts...).ToFunc()
}

// ByEventsURL orders the results by the events_url field.
func ByEventsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEventsURL, opts...).ToFunc()
}

// ByReceivedEventsURL orders the results by the received_events_url field.
func ByReceivedEventsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReceivedEventsURL, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// BySiteAdmin orders the results by the site_admin field.
func BySiteAdmin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSiteAdmin, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCompany orders the results by the company field.
func ByCompany(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCompany, opts...).ToFunc()
}

// ByBlog orders the results by the blog field.
func ByBlog(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBlog, opts...).ToFunc()
}

// ByLocation orders the results by the location field.
func ByLocation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocation, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByHireable orders the results by the hireable field.
func ByHireable(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHireable, opts...).ToFunc()
}

// ByBio orders the results by the bio field.
func ByBio(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBio, opts...).ToFunc()
}

// ByPublicRepos orders the results by the public_repos field.
func ByPublicRepos(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublicRepos, opts...).ToFunc()
}

// ByPublicGists orders the results by the public_gists field.
func ByPublicGists(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublicGists, opts...).ToFunc()
}

// ByFollowers orders the results by the followers field.
func ByFollowers(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFollowers, opts...).ToFunc()
}

// ByFollowing orders the results by the following field.
func ByFollowing(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFollowing, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByRepositoriesCount orders the results by repositories count.
func ByRepositoriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRepositoriesStep(), opts...)
	}
}

// ByRepositories orders the results by repositories terms.
func ByRepositories(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRepositoriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByIssuesCreatedCount orders the results by issues_created count.
func ByIssuesCreatedCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIssuesCreatedStep(), opts...)
	}
}

// ByIssuesCreated orders the results by issues_created terms.
func ByIssuesCreated(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIssuesCreatedStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCommentsCreatedCount orders the results by comments_created count.
func ByCommentsCreatedCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCommentsCreatedStep(), opts...)
	}
}

// ByCommentsCreated orders the results by comments_created terms.
func ByCommentsCreated(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommentsCreatedStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByIssuesAssignedCount orders the results by issues_assigned count.
func ByIssuesAssignedCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIssuesAssignedStep(), opts...)
	}
}

// ByIssuesAssigned orders the results by issues_assigned terms.
func ByIssuesAssigned(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIssuesAssignedStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByIssuesClosedCount orders the results by issues_closed count.
func ByIssuesClosedCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIssuesClosedStep(), opts...)
	}
}

// ByIssuesClosed orders the results by issues_closed terms.
func ByIssuesClosed(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIssuesClosedStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRepositoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RepositoriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RepositoriesTable, RepositoriesColumn),
	)
}
func newIssuesCreatedStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IssuesCreatedInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, IssuesCreatedTable, IssuesCreatedColumn),
	)
}
func newCommentsCreatedStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommentsCreatedInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CommentsCreatedTable, CommentsCreatedColumn),
	)
}
func newIssuesAssignedStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IssuesAssignedInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, IssuesAssignedTable, IssuesAssignedPrimaryKey...),
	)
}
func newIssuesClosedStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IssuesClosedInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, IssuesClosedTable, IssuesClosedColumn),
	)
}
