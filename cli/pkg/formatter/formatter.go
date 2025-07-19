package formatter

import (
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

// EntityMeta aggregates formatted names for an entity.
type EntityMeta struct {
	EntityName   string // PascalCase singular
	PluralPascal string // PascalCase plural
	PluralSnake  string // snake_case plural
	PluralKebab  string // kebab-case plural
}

// BuildEntityMeta formats the entity name and infers its plural form using the
// inflection library.
func BuildEntityMeta(name string) EntityMeta {
	plural := inflection.Plural(name)

	return EntityMeta{
		EntityName:   strcase.ToCamel(name),
		PluralPascal: strcase.ToCamel(plural),
		PluralSnake:  strcase.ToSnake(plural),
		PluralKebab:  strcase.ToKebab(plural),
	}
}

// ToSnake converts a CamelCase string into snake_case.
func ToSnake(s string) string {
	var out []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			out = append(out, '_')
		}
		out = append(out, r)
	}
	return strings.ToLower(string(out))
}
