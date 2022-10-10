package main

import (
	"errors"
	"fmt"
)

type Matrix struct {
	Values      [][]float64
	Height      int
	Weight      int
	isQuadratic bool
	MaxValue    float64
}

func (m *Matrix) Set(numbers ...float64) error {
	if len(numbers) != (m.Height * m.Weight) {
		return errors.New("Las dimensiones no coinciden")
	}
	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Weight; j++ {

		}
	}
	return nil
}

func (m *Matrix) Print() {
	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Weight; j++ {
			fmt.Printf("%v", m.Values[i][j])
		}
	}
}

func main() {

}
