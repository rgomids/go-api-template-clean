package scaffold

import "fmt"

// Generate creates the scaffold files for the given spec.
func Generate(spec *ScaffoldSpec) error {
	if spec == nil {
		return fmt.Errorf("spec is nil")
	}
	// TODO: implement file generation using templates
	return nil
}
