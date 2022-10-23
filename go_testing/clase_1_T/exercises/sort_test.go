package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	numbers, expected := []int{5, 7, 1, 10, 3}, []int{1, 3, 5, 7, 10}

	result := InsertionSort(numbers)

	assert.Equal(t, expected, result)
	assert.Equal(t, len(expected), len(result))
}
