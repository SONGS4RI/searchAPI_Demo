package customerror

import "errors"

type CustomError struct {
	Code int
	Cerror  error
}

func (e *CustomError) Error() string {
    return e.Cerror.Error()
}

var ErrNotFound = &CustomError{
    Code:  404,
    Cerror: errors.New("NOT FOUND"),
}