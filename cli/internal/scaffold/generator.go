package scaffold

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
	"time"

	"github.com/rgomids/go-api-template-clean/cli/pkg/file"
	"github.com/rgomids/go-api-template-clean/cli/pkg/formatter"
)

var templateDir string

func init() {
	if _, f, _, ok := runtime.Caller(0); ok {
		templateDir = filepath.Join(filepath.Dir(f), "templates")
	} else {
		templateDir = "cli/internal/scaffold/templates"
	}
}

type fieldTemplate struct {
	Name    string
	Column  string
	GoType  string
	SQLType string
}

type entityTemplate struct {
	Entity      string
	EntitySnake string
	Fields      []fieldTemplate
}

// Generate creates the scaffold files for the given spec.
func Generate(spec *ScaffoldSpec) error {
	if spec == nil {
		return fmt.Errorf("spec is nil")
	}

	data := buildTemplateData(spec)

	files := map[string]string{
		"entity.tmpl":          filepath.Join("internal/domain/entity", formatter.ToSnake(spec.EntityName)+".go"),
		"repository.tmpl":      filepath.Join("internal/domain/repository", formatter.ToSnake(spec.EntityName)+"_repository.go"),
		"service.tmpl":         filepath.Join("internal/domain/service", formatter.ToSnake(spec.EntityName)+"_service.go"),
		"usecase.tmpl":         filepath.Join("internal/domain/usecase", formatter.ToSnake(spec.EntityName)+"_usecase.go"),
		"handler.tmpl":         filepath.Join("internal/handler/http", formatter.ToSnake(spec.EntityName)+"_handler.go"),
		"handler_test.tmpl":    filepath.Join("internal/handler/http", formatter.ToSnake(spec.EntityName)+"_handler_test.go"),
		"usecase_test.tmpl":    filepath.Join("internal/domain/usecase", formatter.ToSnake(spec.EntityName)+"_usecase_test.go"),
		"mock_repository.tmpl": filepath.Join("mocks", formatter.ToSnake(spec.EntityName)+"_repository_mock.go"),
	}

	for tmpl, dest := range files {
		if err := generateFile(tmpl, dest, data); err != nil {
			return err
		}
	}

	if err := generateMigrations(spec, data); err != nil {
		return err
	}

	return nil
}

func generateFile(tmplName, dest string, data entityTemplate) error {
	path := filepath.Join(templateDir, tmplName)
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	t, err := template.New(tmplName).Parse(string(b))
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return err
	}
	return file.WriteFile(dest, buf.Bytes())
}

func generateMigrations(spec *ScaffoldSpec, data entityTemplate) error {
	ts := time.Now().Unix()
	name := fmt.Sprintf("%d_create_%s_table", ts, formatter.ToSnake(spec.EntityName)+"s")
	up := filepath.Join("db/migrations", name+".up.sql")
	down := filepath.Join("db/migrations", name+".down.sql")
	if err := generateFile("migration_up.tmpl", up, data); err != nil {
		return err
	}
	if err := generateFile("migration_down.tmpl", down, data); err != nil {
		return err
	}
	return nil
}

func buildTemplateData(spec *ScaffoldSpec) entityTemplate {
	d := entityTemplate{
		Entity:      toPascal(spec.EntityName),
		EntitySnake: formatter.ToSnake(spec.EntityName),
	}
	for _, f := range spec.Fields {
		d.Fields = append(d.Fields, fieldTemplate{
			Name:    toPascal(f.Name),
			Column:  formatter.ToSnake(f.Name),
			GoType:  goType(f),
			SQLType: sqlType(f),
		})
	}
	return d
}

func toPascal(s string) string {
	if s == "" {
		return ""
	}
	parts := strings.FieldsFunc(s, func(r rune) bool { return r == '_' || r == '-' || r == ' ' })
	for i, p := range parts {
		parts[i] = strings.Title(p)
	}
	return strings.Join(parts, "")
}

func goType(f FieldSpec) string {
	switch f.Type {
	case "string", "enum":
		return "string"
	case "int":
		return "int"
	case "float":
		return "float64"
	case "bool":
		return "bool"
	case "time":
		return "time.Time"
	case "json":
		return "[]byte"
	case "array":
		elem := goType(FieldSpec{Type: f.Subtype})
		if elem == "" {
			elem = "string"
		}
		return "[]" + elem
	default:
		return "string"
	}
}

func sqlType(f FieldSpec) string {
	switch f.Type {
	case "string", "enum":
		return "text"
	case "int":
		return "integer"
	case "float":
		return "numeric"
	case "bool":
		return "boolean"
	case "time":
		return "timestamp"
	case "json":
		return "jsonb"
	case "array":
		elem := sqlType(FieldSpec{Type: f.Subtype})
		if elem == "" {
			elem = "text"
		}
		return elem + "[]"
	default:
		return "text"
	}
}
