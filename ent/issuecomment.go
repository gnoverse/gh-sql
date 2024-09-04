// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gnolang/gh-sql/ent/issue"
	"github.com/gnolang/gh-sql/ent/issuecomment"
	"github.com/gnolang/gh-sql/ent/user"
	"github.com/gnolang/gh-sql/pkg/model"
)

// IssueComment is the model entity for the IssueComment schema.
type IssueComment struct {
	config `json:"-"`
	// ID of the ent.
	// Unique identifier of the issue comment
	ID int64 `json:"id,omitempty"`
	// NodeID holds the value of the "node_id" field.
	NodeID string `json:"node_id"`
	// URL for the issue comment
	URL string `json:"url"`
	// Body holds the value of the "body" field.
	Body string `json:"body"`
	// HTMLURL holds the value of the "html_url" field.
	HTMLURL string `json:"html_url"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt string `json:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt string `json:"updated_at"`
	// IssueURL holds the value of the "issue_url" field.
	IssueURL string `json:"issue_url"`
	// AuthorAssociation holds the value of the "author_association" field.
	AuthorAssociation model.AuthorAssociation `json:"author_association"`
	// Reactions holds the value of the "reactions" field.
	Reactions map[string]interface{} `json:"reactions"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the IssueCommentQuery when eager-loading is set.
	Edges                 IssueCommentEdges `json:"-"`
	issue_comments        *int64
	user_comments_created *int64
	selectValues          sql.SelectValues
}

// IssueCommentEdges holds the relations/edges for other nodes in the graph.
type IssueCommentEdges struct {
	// Issue holds the value of the issue edge.
	Issue *Issue `json:"issue,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// IssueOrErr returns the Issue value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e IssueCommentEdges) IssueOrErr() (*Issue, error) {
	if e.Issue != nil {
		return e.Issue, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: issue.Label}
	}
	return nil, &NotLoadedError{edge: "issue"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e IssueCommentEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*IssueComment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case issuecomment.FieldReactions:
			values[i] = new([]byte)
		case issuecomment.FieldID:
			values[i] = new(sql.NullInt64)
		case issuecomment.FieldNodeID, issuecomment.FieldURL, issuecomment.FieldBody, issuecomment.FieldHTMLURL, issuecomment.FieldCreatedAt, issuecomment.FieldUpdatedAt, issuecomment.FieldIssueURL, issuecomment.FieldAuthorAssociation:
			values[i] = new(sql.NullString)
		case issuecomment.ForeignKeys[0]: // issue_comments
			values[i] = new(sql.NullInt64)
		case issuecomment.ForeignKeys[1]: // user_comments_created
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the IssueComment fields.
func (ic *IssueComment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case issuecomment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ic.ID = int64(value.Int64)
		case issuecomment.FieldNodeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field node_id", values[i])
			} else if value.Valid {
				ic.NodeID = value.String
			}
		case issuecomment.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				ic.URL = value.String
			}
		case issuecomment.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				ic.Body = value.String
			}
		case issuecomment.FieldHTMLURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field html_url", values[i])
			} else if value.Valid {
				ic.HTMLURL = value.String
			}
		case issuecomment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ic.CreatedAt = value.String
			}
		case issuecomment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ic.UpdatedAt = value.String
			}
		case issuecomment.FieldIssueURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field issue_url", values[i])
			} else if value.Valid {
				ic.IssueURL = value.String
			}
		case issuecomment.FieldAuthorAssociation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author_association", values[i])
			} else if value.Valid {
				ic.AuthorAssociation = model.AuthorAssociation(value.String)
			}
		case issuecomment.FieldReactions:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field reactions", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ic.Reactions); err != nil {
					return fmt.Errorf("unmarshal field reactions: %w", err)
				}
			}
		case issuecomment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field issue_comments", value)
			} else if value.Valid {
				ic.issue_comments = new(int64)
				*ic.issue_comments = int64(value.Int64)
			}
		case issuecomment.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_comments_created", value)
			} else if value.Valid {
				ic.user_comments_created = new(int64)
				*ic.user_comments_created = int64(value.Int64)
			}
		default:
			ic.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the IssueComment.
// This includes values selected through modifiers, order, etc.
func (ic *IssueComment) Value(name string) (ent.Value, error) {
	return ic.selectValues.Get(name)
}

// QueryIssue queries the "issue" edge of the IssueComment entity.
func (ic *IssueComment) QueryIssue() *IssueQuery {
	return NewIssueCommentClient(ic.config).QueryIssue(ic)
}

// QueryUser queries the "user" edge of the IssueComment entity.
func (ic *IssueComment) QueryUser() *UserQuery {
	return NewIssueCommentClient(ic.config).QueryUser(ic)
}

// Update returns a builder for updating this IssueComment.
// Note that you need to call IssueComment.Unwrap() before calling this method if this IssueComment
// was returned from a transaction, and the transaction was committed or rolled back.
func (ic *IssueComment) Update() *IssueCommentUpdateOne {
	return NewIssueCommentClient(ic.config).UpdateOne(ic)
}

// Unwrap unwraps the IssueComment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ic *IssueComment) Unwrap() *IssueComment {
	_tx, ok := ic.config.driver.(*txDriver)
	if !ok {
		panic("ent: IssueComment is not a transactional entity")
	}
	ic.config.driver = _tx.drv
	return ic
}

// String implements the fmt.Stringer.
func (ic *IssueComment) String() string {
	var builder strings.Builder
	builder.WriteString("IssueComment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ic.ID))
	builder.WriteString("node_id=")
	builder.WriteString(ic.NodeID)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(ic.URL)
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(ic.Body)
	builder.WriteString(", ")
	builder.WriteString("html_url=")
	builder.WriteString(ic.HTMLURL)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ic.CreatedAt)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ic.UpdatedAt)
	builder.WriteString(", ")
	builder.WriteString("issue_url=")
	builder.WriteString(ic.IssueURL)
	builder.WriteString(", ")
	builder.WriteString("author_association=")
	builder.WriteString(fmt.Sprintf("%v", ic.AuthorAssociation))
	builder.WriteString(", ")
	builder.WriteString("reactions=")
	builder.WriteString(fmt.Sprintf("%v", ic.Reactions))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (ic *IssueComment) MarshalJSON() ([]byte, error) {
	type Alias IssueComment
	return json.Marshal(&struct {
		*Alias
		IssueCommentEdges
	}{
		Alias:             (*Alias)(ic),
		IssueCommentEdges: ic.Edges,
	})
}

// IssueComments is a parsable slice of IssueComment.
type IssueComments []*IssueComment
