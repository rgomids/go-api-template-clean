package app

// container.go centralizes dependency injection.
// Example functions illustrate how services and handlers would be constructed.

// NewExampleService provides an instance of ExampleService with injected dependencies.
func NewExampleService() interface{} {
	// TODO: construct service dependencies
	return nil
}

// NewExampleHandler provides an instance of ExampleHandler using explicit injection.
func NewExampleHandler(service interface{}) interface{} {
	// TODO: construct handler with service
	return nil
}
