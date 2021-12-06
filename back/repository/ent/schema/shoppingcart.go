package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// ShoppingCart holds the schema definition for the ShoppingCart entity.
type ShoppingCart struct {
	ent.Schema
}

// Fields of the ShoppingCart.
func (ShoppingCart) Fields() []ent.Field {
	return []ent.Field{
		field.Int("amount"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ShoppingCart.
func (ShoppingCart) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("book", Book.Type).Ref("shopping_cart").Unique(),
		edge.From("user", User.Type).Ref("shopping_cart").Unique(),
	}
}
