package scaffold

import (
	"fmt"
	"strings"
)

// FieldSpec represents a single field of the entity.
// Type stores the base type (string, float, enum, array, json, ...).
// Subtype holds any nested information such as array element or enum values.
type FieldSpec struct {
	Name    string
	Type    string
	Subtype string
	Raw     string
}

// RelationshipSpec captures an association with another entity.
type RelationshipSpec struct {
	Name         string
	Relationship string
}

// ScaffoldSpec aggregates all parsed information about an entity.
type ScaffoldSpec struct {
	EntityName    string
	Fields        []FieldSpec
	Relationships []RelationshipSpec
}

// Parse interprets CLI arguments into a ScaffoldSpec.
// It expects at least the entity name followed by field or relationship
// declarations in the format "name:type".
func Parse(args []string) (*ScaffoldSpec, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("no entity provided")
	}

	spec := &ScaffoldSpec{EntityName: args[0]}

	for _, a := range args[1:] {
		f, r, err := parseArg(a)
		if err != nil {
			return nil, err
		}
		if f != nil {
			spec.Fields = append(spec.Fields, *f)
		} else if r != nil {
			spec.Relationships = append(spec.Relationships, *r)
		}
	}

	return spec, nil
}

// parseArg analyses a single CLI argument and decides whether it describes a
// field or a relationship. It returns the resulting specification or an error
// when the syntax is invalid.
func parseArg(arg string) (*FieldSpec, *RelationshipSpec, error) {
	parts := strings.SplitN(arg, ":", 2)
	if len(parts) != 2 {
		return nil, nil, fmt.Errorf("invalid argument %q", arg)
	}

	name := parts[0]
	typ := parts[1]

	if isRelationshipType(typ) {
		return nil, &RelationshipSpec{Name: name, Relationship: typ}, nil
	}

	fieldType, subtype, err := splitType(typ)
	if err != nil {
		return nil, nil, err
	}

	f := &FieldSpec{
		Name:    name,
		Type:    fieldType,
		Subtype: subtype,
		Raw:     typ,
	}

	return f, nil, nil
}

// splitType breaks a complex type string like "enum[a,b]" or "array[string]" in
// its base type and subtype. When no subtype is present the subtype will be an
// empty string.
func splitType(t string) (string, string, error) {
	if idx := strings.Index(t, "["); idx != -1 {
		if !strings.HasSuffix(t, "]") {
			return "", "", fmt.Errorf("invalid type %q", t)
		}
		return t[:idx], t[idx+1 : len(t)-1], nil
	}
	return t, "", nil
}

// isRelationshipType reports whether the given type token represents a
// relationship declaration instead of a field.
func isRelationshipType(t string) bool {
	switch t {
	case "belongsTo", "hasMany", "hasOne", "manyToMany":
		return true
	default:
		return false
	}
}
