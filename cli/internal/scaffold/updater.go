package scaffold

import "fmt"

// Update modifies existing project files to register new components.
func Update(spec *ScaffoldSpec) error {
	if spec == nil {
		return fmt.Errorf("spec is nil")
	}
	// TODO: implement updates to routes and container
	return nil
}
