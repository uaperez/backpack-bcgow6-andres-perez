package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {
	// arrange
	num1 := 2
	num2 := 4
	expected := -2
	// act
	result := Sub(num1, num2)
	// assert
	assert.Equal(t, expected, result, "Número esperado: %v | Número obtenido: %v", expected, result)
}

// Punto 1
func TestDivideReceiveCorrectResult(t *testing.T) {
	// arrange
	num1 := 10
	num2 := 2
	expected := 5
	// act
	result, err := Divide(num1, num2)
	// asert
	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

// Punto 3
func TestDivideReceiveError(t *testing.T) {
	num1, num2 := 15, 0
	expected := 0

	result, err := Divide(num1, num2)

	assert.NotNil(t, err, "You have an error")
	assert.Equal(t, expected, result)
}
