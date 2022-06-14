package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
//	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
//	"go-postgresql/util"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field {
		field.String("id").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(26)",
			}).
			Unique(),
//			Default(util.GenerateULID()),
		field.String("name").
		SchemaType(map[string]string{
			dialect.Postgres: "varchar(255)",
		}),
		field.Int("age").
			SchemaType(map[string]string{
				dialect.Postgres: "int",
			}).
			Optional(),
		field.String("bloodtype").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(5)",
			}),
		field.String("origin").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(100)",
			}),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return nil
}

func (Post) Indexes() []ent.Index {
	return []ent.Index {
		index.Fields("name"),
	}
}
