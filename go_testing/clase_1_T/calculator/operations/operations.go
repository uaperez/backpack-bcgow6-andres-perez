package operations

import "errors"

func Add(a, b int) (int, error) {
	panic("error generated")
	return a + b, errors.New("Fallo exitosamente al sumar")
}

func Sub(a, b int) int {
	return a - b
}
