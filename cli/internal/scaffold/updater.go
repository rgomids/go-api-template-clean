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
	if err := updateMain(spec); err != nil {
		return err
	}

	if err := updateRoutesTest(spec); err != nil {
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

	pascal := toPascal(spec.EntityName)
	camel := strings.ToLower(pascal[:1]) + pascal[1:]

	param := fmt.Sprintf("%sHandler *httphandler.%sHandler", camel, pascal)
	if !strings.Contains(content, param) {
		old := "*httphandler.UserHandler)"
		newSig := fmt.Sprintf("*httphandler.UserHandler, %s)", param)
		content = strings.Replace(content, old, newSig, 1)
	}

	snippet := fmt.Sprintf("\trouter.Route(\"%s\", func(r chi.Router) {\n", route)
	snippet += fmt.Sprintf("\t\tr.Post(\"/\", %sHandler.Create)\n", camel)
	snippet += fmt.Sprintf("\t\tr.Get(\"/\", %sHandler.List)\n", camel)
	snippet += fmt.Sprintf("\t\tr.Get(\"/{id}\", %sHandler.GetByID)\n", camel)
	snippet += fmt.Sprintf("\t\tr.Put(\"/{id}\", %sHandler.Update)\n", camel)
	snippet += fmt.Sprintf("\t\tr.Delete(\"/{id}\", %sHandler.Delete)\n", camel)
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

func updateMain(spec *ScaffoldSpec) error {
	path := filepath.Join("cmd", "main.go")
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	content := string(b)

	pascal := toPascal(spec.EntityName)
	if strings.Contains(content, "c."+pascal+"Handler") {
		return nil
	}

	old := "httproutes.RegisterRoutes(r, c.UserHandler)"
	newCall := fmt.Sprintf("httproutes.RegisterRoutes(r, c.UserHandler, c.%sHandler)", pascal)
	content = strings.Replace(content, old, newCall, 1)

	return file.WriteFile(path, []byte(content))
}

func updateRoutesTest(spec *ScaffoldSpec) error {
	path := filepath.Join("internal/handler/http/routes", "routes_test.go")
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	content := string(b)

	pascal := toPascal(spec.EntityName)
	camel := strings.ToLower(pascal[:1]) + pascal[1:]
	if strings.Contains(content, camel+" :=") {
		return nil
	}

	insert := fmt.Sprintf("\t%[1]s := httphandler.New%[2]sHandler(nil)\n", camel, pascal)
	content = strings.Replace(content, "r := chi.NewRouter()", insert+"\tr := chi.NewRouter()", 1)
	content = strings.Replace(content, "RegisterRoutes(r, h)", fmt.Sprintf("RegisterRoutes(r, h, %s)", camel), 1)

	return file.WriteFile(path, []byte(content))
}
