package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("error n1")

func x() error {
	return fmt.Errorf("Info extra: %w", ErrNotFound)
}

type MyError struct {
	StatusCode int
	Message    string
}

func (err *MyError) Error() string {
	return fmt.Sprintf("%s (%d)", err.Message, err.StatusCode)
}

func GetMyError(status int) (code int, err error) {
	if status >= 400 {
		return 500, &MyError{
			StatusCode: 500,
			Message:    "Algo salió mal... :/",
		}
	}
	return 200, nil
}

func main() {
	err1 := &MyError{
		StatusCode: 502,
		Message:    "Soy el error 1, distinto al 2",
	}

	err2 := &MyError{
		StatusCode: 406,
		Message:    "Soy el error 2, distinto al 1",
	}

	bothErrorsAreEqual := errors.As(err1, &err2)

	fmt.Println(bothErrorsAreEqual)
}
