package hystrix

import "fmt"

type badRequestError struct {
	err error
}

func NewBadRequestError(s string) *badRequestError {
	return &badRequestError{
		err: CircuitError{"bad request error: " + s},
	}
}

func WarpBadRequestError(err error) *badRequestError {
	return &badRequestError{
		err: fmt.Errorf("hystrix: bad request error: %w", err),
	}
}

func (b *badRequestError) Error() string {
	return b.err.Error()
}

func isBadRequestError(err error) bool {
	_, ok := err.(*badRequestError)
	return ok
}
