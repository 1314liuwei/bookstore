package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Int("amount"),
		field.Enum("status").Values("completed", "to_be_paid", "transferring"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type),
		edge.To("book", Book.Type),
	}
}
