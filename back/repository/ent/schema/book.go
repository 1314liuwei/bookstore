package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.Int("price"),
		field.Int("surplus_catch"),
		field.String("author"),
		field.String("describe"),
		field.String("ebook"),
		field.String("cover"),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("category", Category.Type),
	}
}
