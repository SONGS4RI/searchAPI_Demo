package customerror

import "errors"

type CustomError struct {
	Code int
	Err  error
}

func (e *CustomError) Error() string {
    return e.Err.Error()
}

var ErrNotFound = &CustomError{
    Code:  404,
    Err: errors.New("NOT FOUND"),
}