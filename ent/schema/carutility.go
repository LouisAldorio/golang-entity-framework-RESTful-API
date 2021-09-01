package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// CarUtility holds the schema definition for the CarUtility entity.
type CarUtility struct {
	ent.Schema
}

// Fields of the CarUtility.
func (CarUtility) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("unknown"),
	}
}

// Edges of the CarUtility.
func (CarUtility) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "owner" of type `User`
		// and reference it to the "cars" edge (in User schema)
		// explicitly using the `Ref` method.
		edge.From("cars", Car.Type).
			Ref("car_utilities"),
	}
}
