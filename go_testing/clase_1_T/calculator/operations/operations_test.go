package operations

import (
	"testing"

	"github.com/stretchr/testify/assert" // se construye sobre el package testing, por lo que es importante importarlo también
)

// go test -v ====== para ver la info de cuántos test pasó y cuántos no

func TestAddGood(t *testing.T) {
	// Arrange = valores de entrada
	num1 := 10
	num2 := 30
	await := 40
	// Act = ejecutar la función
	result, err := Add(num1, num2)
	// Assert
	assert.Equal(t, await, result, "Number expected: %v | Number obteined: %v", await, result)
	assert.NotNil(t, err, err.Error())
	/*if result != await {
		t.Errorf("Number expected: %v | Number obteined: %v", await, result)
	}*/
}

func TestAddBad(t *testing.T) {
	// arrange
	// act
	// assert
}
