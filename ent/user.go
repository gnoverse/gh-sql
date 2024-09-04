// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gnolang/gh-sql/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Login holds the value of the "login" field.
	Login string `json:"login"`
	// NodeID holds the value of the "node_id" field.
	NodeID string `json:"node_id"`
	// AvatarURL holds the value of the "avatar_url" field.
	AvatarURL string `json:"avatar_url"`
	// GravatarID holds the value of the "gravatar_id" field.
	GravatarID string `json:"gravatar_id"`
	// URL holds the value of the "url" field.
	URL string `json:"url"`
	// HTMLURL holds the value of the "html_url" field.
	HTMLURL string `json:"html_url"`
	// FollowersURL holds the value of the "followers_url" field.
	FollowersURL string `json:"followers_url"`
	// FollowingURL holds the value of the "following_url" field.
	FollowingURL string `json:"following_url"`
	// GistsURL holds the value of the "gists_url" field.
	GistsURL string `json:"gists_url"`
	// StarredURL holds the value of the "starred_url" field.
	StarredURL string `json:"starred_url"`
	// SubscriptionsURL holds the value of the "subscriptions_url" field.
	SubscriptionsURL string `json:"subscriptions_url"`
	// OrganizationsURL holds the value of the "organizations_url" field.
	OrganizationsURL string `json:"organizations_url"`
	// ReposURL holds the value of the "repos_url" field.
	ReposURL string `json:"repos_url"`
	// EventsURL holds the value of the "events_url" field.
	EventsURL string `json:"events_url"`
	// ReceivedEventsURL holds the value of the "received_events_url" field.
	ReceivedEventsURL string `json:"received_events_url"`
	// Type holds the value of the "type" field.
	Type string `json:"type"`
	// SiteAdmin holds the value of the "site_admin" field.
	SiteAdmin bool `json:"site_admin"`
	// Name holds the value of the "name" field.
	Name string `json:"name"`
	// Company holds the value of the "company" field.
	Company string `json:"company"`
	// Blog holds the value of the "blog" field.
	Blog string `json:"blog"`
	// Location holds the value of the "location" field.
	Location string `json:"location"`
	// Email holds the value of the "email" field.
	Email string `json:"email"`
	// Hireable holds the value of the "hireable" field.
	Hireable bool `json:"hireable"`
	// Bio holds the value of the "bio" field.
	Bio string `json:"bio"`
	// PublicRepos holds the value of the "public_repos" field.
	PublicRepos int64 `json:"public_repos"`
	// PublicGists holds the value of the "public_gists" field.
	PublicGists int64 `json:"public_gists"`
	// Followers holds the value of the "followers" field.
	Followers int64 `json:"followers"`
	// Following holds the value of the "following" field.
	Following int64 `json:"following"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"-"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Repositories holds the value of the repositories edge.
	Repositories []*Repository `json:"repositories,omitempty"`
	// IssuesCreated holds the value of the issues_created edge.
	IssuesCreated []*Issue `json:"issues_created,omitempty"`
	// CommentsCreated holds the value of the comments_created edge.
	CommentsCreated []*IssueComment `json:"comments_created,omitempty"`
	// IssuesAssigned holds the value of the issues_assigned edge.
	IssuesAssigned []*Issue `json:"issues_assigned,omitempty"`
	// TimelineEventsCreated holds the value of the timeline_events_created edge.
	TimelineEventsCreated []*TimelineEvent `json:"timeline_events_created,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// RepositoriesOrErr returns the Repositories value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RepositoriesOrErr() ([]*Repository, error) {
	if e.loadedTypes[0] {
		return e.Repositories, nil
	}
	return nil, &NotLoadedError{edge: "repositories"}
}

// IssuesCreatedOrErr returns the IssuesCreated value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) IssuesCreatedOrErr() ([]*Issue, error) {
	if e.loadedTypes[1] {
		return e.IssuesCreated, nil
	}
	return nil, &NotLoadedError{edge: "issues_created"}
}

// CommentsCreatedOrErr returns the CommentsCreated value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CommentsCreatedOrErr() ([]*IssueComment, error) {
	if e.loadedTypes[2] {
		return e.CommentsCreated, nil
	}
	return nil, &NotLoadedError{edge: "comments_created"}
}

// IssuesAssignedOrErr returns the IssuesAssigned value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) IssuesAssignedOrErr() ([]*Issue, error) {
	if e.loadedTypes[3] {
		return e.IssuesAssigned, nil
	}
	return nil, &NotLoadedError{edge: "issues_assigned"}
}

// TimelineEventsCreatedOrErr returns the TimelineEventsCreated value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TimelineEventsCreatedOrErr() ([]*TimelineEvent, error) {
	if e.loadedTypes[4] {
		return e.TimelineEventsCreated, nil
	}
	return nil, &NotLoadedError{edge: "timeline_events_created"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldSiteAdmin, user.FieldHireable:
			values[i] = new(sql.NullBool)
		case user.FieldID, user.FieldPublicRepos, user.FieldPublicGists, user.FieldFollowers, user.FieldFollowing:
			values[i] = new(sql.NullInt64)
		case user.FieldLogin, user.FieldNodeID, user.FieldAvatarURL, user.FieldGravatarID, user.FieldURL, user.FieldHTMLURL, user.FieldFollowersURL, user.FieldFollowingURL, user.FieldGistsURL, user.FieldStarredURL, user.FieldSubscriptionsURL, user.FieldOrganizationsURL, user.FieldReposURL, user.FieldEventsURL, user.FieldReceivedEventsURL, user.FieldType, user.FieldName, user.FieldCompany, user.FieldBlog, user.FieldLocation, user.FieldEmail, user.FieldBio:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int64(value.Int64)
		case user.FieldLogin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field login", values[i])
			} else if value.Valid {
				u.Login = value.String
			}
		case user.FieldNodeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field node_id", values[i])
			} else if value.Valid {
				u.NodeID = value.String
			}
		case user.FieldAvatarURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_url", values[i])
			} else if value.Valid {
				u.AvatarURL = value.String
			}
		case user.FieldGravatarID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gravatar_id", values[i])
			} else if value.Valid {
				u.GravatarID = value.String
			}
		case user.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				u.URL = value.String
			}
		case user.FieldHTMLURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field html_url", values[i])
			} else if value.Valid {
				u.HTMLURL = value.String
			}
		case user.FieldFollowersURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field followers_url", values[i])
			} else if value.Valid {
				u.FollowersURL = value.String
			}
		case user.FieldFollowingURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field following_url", values[i])
			} else if value.Valid {
				u.FollowingURL = value.String
			}
		case user.FieldGistsURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gists_url", values[i])
			} else if value.Valid {
				u.GistsURL = value.String
			}
		case user.FieldStarredURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field starred_url", values[i])
			} else if value.Valid {
				u.StarredURL = value.String
			}
		case user.FieldSubscriptionsURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subscriptions_url", values[i])
			} else if value.Valid {
				u.SubscriptionsURL = value.String
			}
		case user.FieldOrganizationsURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field organizations_url", values[i])
			} else if value.Valid {
				u.OrganizationsURL = value.String
			}
		case user.FieldReposURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field repos_url", values[i])
			} else if value.Valid {
				u.ReposURL = value.String
			}
		case user.FieldEventsURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field events_url", values[i])
			} else if value.Valid {
				u.EventsURL = value.String
			}
		case user.FieldReceivedEventsURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field received_events_url", values[i])
			} else if value.Valid {
				u.ReceivedEventsURL = value.String
			}
		case user.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				u.Type = value.String
			}
		case user.FieldSiteAdmin:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field site_admin", values[i])
			} else if value.Valid {
				u.SiteAdmin = value.Bool
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldCompany:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field company", values[i])
			} else if value.Valid {
				u.Company = value.String
			}
		case user.FieldBlog:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field blog", values[i])
			} else if value.Valid {
				u.Blog = value.String
			}
		case user.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				u.Location = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldHireable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field hireable", values[i])
			} else if value.Valid {
				u.Hireable = value.Bool
			}
		case user.FieldBio:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bio", values[i])
			} else if value.Valid {
				u.Bio = value.String
			}
		case user.FieldPublicRepos:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field public_repos", values[i])
			} else if value.Valid {
				u.PublicRepos = value.Int64
			}
		case user.FieldPublicGists:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field public_gists", values[i])
			} else if value.Valid {
				u.PublicGists = value.Int64
			}
		case user.FieldFollowers:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field followers", values[i])
			} else if value.Valid {
				u.Followers = value.Int64
			}
		case user.FieldFollowing:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field following", values[i])
			} else if value.Valid {
				u.Following = value.Int64
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryRepositories queries the "repositories" edge of the User entity.
func (u *User) QueryRepositories() *RepositoryQuery {
	return NewUserClient(u.config).QueryRepositories(u)
}

// QueryIssuesCreated queries the "issues_created" edge of the User entity.
func (u *User) QueryIssuesCreated() *IssueQuery {
	return NewUserClient(u.config).QueryIssuesCreated(u)
}

// QueryCommentsCreated queries the "comments_created" edge of the User entity.
func (u *User) QueryCommentsCreated() *IssueCommentQuery {
	return NewUserClient(u.config).QueryCommentsCreated(u)
}

// QueryIssuesAssigned queries the "issues_assigned" edge of the User entity.
func (u *User) QueryIssuesAssigned() *IssueQuery {
	return NewUserClient(u.config).QueryIssuesAssigned(u)
}

// QueryTimelineEventsCreated queries the "timeline_events_created" edge of the User entity.
func (u *User) QueryTimelineEventsCreated() *TimelineEventQuery {
	return NewUserClient(u.config).QueryTimelineEventsCreated(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("login=")
	builder.WriteString(u.Login)
	builder.WriteString(", ")
	builder.WriteString("node_id=")
	builder.WriteString(u.NodeID)
	builder.WriteString(", ")
	builder.WriteString("avatar_url=")
	builder.WriteString(u.AvatarURL)
	builder.WriteString(", ")
	builder.WriteString("gravatar_id=")
	builder.WriteString(u.GravatarID)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(u.URL)
	builder.WriteString(", ")
	builder.WriteString("html_url=")
	builder.WriteString(u.HTMLURL)
	builder.WriteString(", ")
	builder.WriteString("followers_url=")
	builder.WriteString(u.FollowersURL)
	builder.WriteString(", ")
	builder.WriteString("following_url=")
	builder.WriteString(u.FollowingURL)
	builder.WriteString(", ")
	builder.WriteString("gists_url=")
	builder.WriteString(u.GistsURL)
	builder.WriteString(", ")
	builder.WriteString("starred_url=")
	builder.WriteString(u.StarredURL)
	builder.WriteString(", ")
	builder.WriteString("subscriptions_url=")
	builder.WriteString(u.SubscriptionsURL)
	builder.WriteString(", ")
	builder.WriteString("organizations_url=")
	builder.WriteString(u.OrganizationsURL)
	builder.WriteString(", ")
	builder.WriteString("repos_url=")
	builder.WriteString(u.ReposURL)
	builder.WriteString(", ")
	builder.WriteString("events_url=")
	builder.WriteString(u.EventsURL)
	builder.WriteString(", ")
	builder.WriteString("received_events_url=")
	builder.WriteString(u.ReceivedEventsURL)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(u.Type)
	builder.WriteString(", ")
	builder.WriteString("site_admin=")
	builder.WriteString(fmt.Sprintf("%v", u.SiteAdmin))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("company=")
	builder.WriteString(u.Company)
	builder.WriteString(", ")
	builder.WriteString("blog=")
	builder.WriteString(u.Blog)
	builder.WriteString(", ")
	builder.WriteString("location=")
	builder.WriteString(u.Location)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("hireable=")
	builder.WriteString(fmt.Sprintf("%v", u.Hireable))
	builder.WriteString(", ")
	builder.WriteString("bio=")
	builder.WriteString(u.Bio)
	builder.WriteString(", ")
	builder.WriteString("public_repos=")
	builder.WriteString(fmt.Sprintf("%v", u.PublicRepos))
	builder.WriteString(", ")
	builder.WriteString("public_gists=")
	builder.WriteString(fmt.Sprintf("%v", u.PublicGists))
	builder.WriteString(", ")
	builder.WriteString("followers=")
	builder.WriteString(fmt.Sprintf("%v", u.Followers))
	builder.WriteString(", ")
	builder.WriteString("following=")
	builder.WriteString(fmt.Sprintf("%v", u.Following))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		*Alias
		UserEdges
	}{
		Alias:     (*Alias)(u),
		UserEdges: u.Edges,
	})
}

// Users is a parsable slice of User.
type Users []*User
