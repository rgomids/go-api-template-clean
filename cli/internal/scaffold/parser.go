package scaffold

import (
	"fmt"
	"strings"
)

// Field represents a simple field in the entity.
type Field struct {
	Name string
	Type string
}

// Relationship defines associations with other entities.
type Relationship struct {
	Name string
	Type string
}

// ScaffoldSpec captures all information required to scaffold a module.
type ScaffoldSpec struct {
	Entity        string
	Fields        []Field
	Relationships []Relationship
}

// Parse interprets CLI arguments into a ScaffoldSpec.
func Parse(args []string) (*ScaffoldSpec, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("no entity provided")
	}
	spec := &ScaffoldSpec{Entity: args[0]}
	for _, a := range args[1:] {
		parts := strings.SplitN(a, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid argument %q", a)
		}
		name, typ := parts[0], parts[1]
		if isRelationshipType(typ) {
			spec.Relationships = append(spec.Relationships, Relationship{Name: name, Type: typ})
		} else {
			spec.Fields = append(spec.Fields, Field{Name: name, Type: typ})
		}
	}
	return spec, nil
}

func isRelationshipType(t string) bool {
	return t == "belongsTo"
}
