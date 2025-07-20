package testing

type StubService[T any] struct {
	Response T
	Error    error
}

func (s *StubService[T]) Execute() (T, error) {
	return s.Response, s.Error
}

// CRUDStubService provides default implementations for a CRUD-style service.
// Each function field can be overridden in tests to customize behavior.
type CRUDStubService[E any] struct {
	CreateFn func(*E) error
	UpdateFn func(string, *E) error
	GetFn    func(string) (*E, error)
	ListFn   func(map[string]interface{}) ([]*E, error)
	DeleteFn func(string) error
}

func (s *CRUDStubService[E]) Create(e *E) error {
	if s.CreateFn != nil {
		return s.CreateFn(e)
	}
	return nil
}

func (s *CRUDStubService[E]) Update(id string, e *E) error {
	if s.UpdateFn != nil {
		return s.UpdateFn(id, e)
	}
	return nil
}

func (s *CRUDStubService[E]) GetByID(id string) (*E, error) {
	if s.GetFn != nil {
		return s.GetFn(id)
	}
	var zero E
	return &zero, nil
}

func (s *CRUDStubService[E]) List(f map[string]interface{}) ([]*E, error) {
	if s.ListFn != nil {
		return s.ListFn(f)
	}
	return []*E{}, nil
}

func (s *CRUDStubService[E]) Delete(id string) error {
	if s.DeleteFn != nil {
		return s.DeleteFn(id)
	}
	return nil
}
