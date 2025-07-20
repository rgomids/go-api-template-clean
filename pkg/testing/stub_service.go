package testing

type StubService[T any] struct {
	Response T
	Error    error
}

func (s *StubService[T]) Execute() (T, error) {
	return s.Response, s.Error
}
