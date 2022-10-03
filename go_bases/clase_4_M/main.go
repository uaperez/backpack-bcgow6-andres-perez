package main

import (
	"errors"
	"fmt"
	"os"
)

type MyCustomError struct {
	StatusCode int
	Message    string
}

func (err *MyCustomError) Error() string {
	return fmt.Sprintf("%s (%d)", err.Message, err.StatusCode)
}

func GetError(status int) (code int, err error) {
	if status >= 400 {
		return 500, &MyCustomError{
			StatusCode: 500,
			Message:    "Algo salió mal... :/",
		}
	}
	return 200, nil
}

func main() {

	statusCode, err := GetError(300)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Status %d, todo ha salido bien!", statusCode)

	err1 := fmt.Errorf("First error :/")

	// Continua ejecucion
	// 50 lineas despues

	err2 := fmt.Errorf("second error: %w", err1)

	err3 := errors.New("Mi tercer error :/")

	fmt.Println(err1)
	fmt.Println(err2)
	fmt.Println(err3)
}
