package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rgomids/go-api-template-clean/cli/pkg/file"
	"github.com/rgomids/go-api-template-clean/cli/pkg/formatter"
)

// Update modifies existing project files to register new components.
func Update(spec *ScaffoldSpec) error {
	if spec == nil {
		return fmt.Errorf("spec is nil")
	}

	if err := updateRoutes(spec); err != nil {
		return err
	}
	if err := updateContainer(spec); err != nil {
		return err
	}

	return nil
}

func updateRoutes(spec *ScaffoldSpec) error {
	path := filepath.Join("internal/handler/http/routes", "routes.go")
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	content := string(b)
	marker := "// [AUTO-GENERATED-ROUTES]"

	route := fmt.Sprintf("/%ss", formatter.ToSnake(spec.EntityName))
	if strings.Contains(content, route) {
		return nil
	}

	if !strings.Contains(content, marker) {
		return fmt.Errorf("routes marker not found")
	}

	snippet := fmt.Sprintf("\trouter.Route(\"%s\", func(r chi.Router) {\n", route)
	snippet += "\t\tr.Post(\"/\", handler.Create)\n"
	snippet += "\t\tr.Get(\"/\", handler.List)\n"
	snippet += "\t\tr.Get(\"/{id}\", handler.GetByID)\n"
	snippet += "\t\tr.Put(\"/{id}\", handler.Update)\n"
	snippet += "\t\tr.Delete(\"/{id}\", handler.Delete)\n"
	snippet += "\t})\n"

	content = strings.Replace(content, marker, snippet+"\n\t"+marker, 1)

	return file.WriteFile(path, []byte(content))
}

func updateContainer(spec *ScaffoldSpec) error {
	path := filepath.Join("internal/app", "container.go")
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	content := string(b)
	marker := "// [AUTO-GENERATED-CONTAINER]"

	pascal := toPascal(spec.EntityName)
	camel := strings.ToLower(pascal[:1]) + pascal[1:]

	if strings.Contains(content, pascal+"Handler") {
		return nil
	}

	parts := strings.Split(content, marker)
	if len(parts) < 4 {
		return fmt.Errorf("container markers missing")
	}

	fieldSnippet := fmt.Sprintf("\t%[1]sRepository domainrepo.%[1]sRepository\n\t%[1]sService service.%[1]sService\n\t%[1]sHandler *httphandler.%[1]sHandler\n", pascal)
	varSnippet := fmt.Sprintf("\t%[2]sRepo := New%[1]sRepository()\n\t%[2]sService := New%[1]sService(%[2]sRepo)\n\t%[2]sHandler := New%[1]sHandler(%[2]sService)\n", pascal, camel)
	returnSnippet := fmt.Sprintf("\t\t%[1]sRepository: %[2]sRepo,\n\t\t%[1]sService:   %[2]sService,\n\t\t%[1]sHandler:   %[2]sHandler,\n", pascal, camel)

	content = parts[0] + fieldSnippet + marker + parts[1] + varSnippet + marker + parts[2] + returnSnippet + marker + parts[3]

	return file.WriteFile(path, []byte(content))
}
