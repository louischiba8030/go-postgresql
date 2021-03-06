// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-postgresql/ent/post"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Post is the model entity for the Post schema.
type Post struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Bloodtype holds the value of the "bloodtype" field.
	Bloodtype string `json:"bloodtype,omitempty"`
	// Origin holds the value of the "origin" field.
	Origin string `json:"origin,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Post) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case post.FieldAge:
			values[i] = new(sql.NullInt64)
		case post.FieldID, post.FieldName, post.FieldBloodtype, post.FieldOrigin:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Post", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Post fields.
func (po *Post) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case post.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				po.ID = value.String
			}
		case post.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				po.Name = value.String
			}
		case post.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				po.Age = int(value.Int64)
			}
		case post.FieldBloodtype:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bloodtype", values[i])
			} else if value.Valid {
				po.Bloodtype = value.String
			}
		case post.FieldOrigin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field origin", values[i])
			} else if value.Valid {
				po.Origin = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Post.
// Note that you need to call Post.Unwrap() before calling this method if this Post
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Post) Update() *PostUpdateOne {
	return (&PostClient{config: po.config}).UpdateOne(po)
}

// Unwrap unwraps the Post entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Post) Unwrap() *Post {
	tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Post is not a transactional entity")
	}
	po.config.driver = tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Post) String() string {
	var builder strings.Builder
	builder.WriteString("Post(")
	builder.WriteString(fmt.Sprintf("id=%v", po.ID))
	builder.WriteString(", name=")
	builder.WriteString(po.Name)
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", po.Age))
	builder.WriteString(", bloodtype=")
	builder.WriteString(po.Bloodtype)
	builder.WriteString(", origin=")
	builder.WriteString(po.Origin)
	builder.WriteByte(')')
	return builder.String()
}

// Posts is a parsable slice of Post.
type Posts []*Post

func (po Posts) config(cfg config) {
	for _i := range po {
		po[_i].config = cfg
	}
}
