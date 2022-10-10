package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BubbleSort(arr []int, ch chan []int) {
	fmt.Println("Executing BubbleSort...")
	numbers := DuplicateArray(arr)
	startTime := time.Now()
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < (len(numbers) - i - 1); j++ {
			if numbers[j] > numbers[j+1] {
				temp := numbers[j]
				numbers[j] = numbers[j+1]
				numbers[j+1] = temp
			}
		}
	}
	endTime := time.Now()
	fmt.Printf("\tDuración del Bubble Sort: %v\n", endTime.Sub(startTime))
	ch <- numbers
}

func SelectionSort(arr []int, ch chan []int) {
	fmt.Println("Executing SelectionSort...")
	numbers := DuplicateArray(arr)
	arrayLength := len(numbers)
	startTime := time.Now()
	for i := 0; i < arrayLength; i++ {
		indexMin := i
		for j := indexMin; j < arrayLength; j++ {
			if numbers[j] < numbers[indexMin] {
				indexMin = j
			}
		}
		if numbers[i] > numbers[indexMin] {
			temporal := numbers[i]
			numbers[i] = numbers[indexMin]
			numbers[indexMin] = temporal
		}
	}
	endTime := time.Now()
	fmt.Printf("\tDuración del Selection Sort: %v\n", endTime.Sub(startTime))
	ch <- numbers
}

func InsertionSort(arr []int, ch chan []int) {
	fmt.Println("Executing InsertionSort...")
	numbers := DuplicateArray(arr)
	arrayLength := len(numbers)
	startTime := time.Now()
	for i := 1; i < arrayLength; i++ {
		key := numbers[i]
		j := i - 1

		for j >= 0 && numbers[j] > key {
			numbers[j+1] = numbers[j]
			j = j - 1
		}

		numbers[j+1] = key
	}
	endTime := time.Now()
	fmt.Printf("\tDuración del Insertion Sort: %v\n", endTime.Sub(startTime))
	ch <- numbers
}

func DuplicateArray(numbers []int) []int {
	newArray := make([]int, len(numbers))
	copy(newArray, numbers)
	return newArray
}

func main() {
	fmt.Println("====== 100 NUMBERS =======")
	numbers := rand.Perm(100)

	channel := make(chan []int)

	go BubbleSort(numbers, channel)
	go InsertionSort(numbers, channel)
	go SelectionSort(numbers, channel)

	var lecture [][]int

	for i := 0; i < 3; i++ {
		lecture = append(lecture, <-channel)
	}

	fmt.Println("====== 1000 NUMBERS =======")

	numbers1 := rand.Perm(1000)

	go BubbleSort(numbers1, channel)
	go InsertionSort(numbers1, channel)
	go SelectionSort(numbers1, channel)

	for i := 0; i < 3; i++ {
		lecture = append(lecture, <-channel)
	}

	fmt.Println("====== 10000 NUMBERS =======")

	numbers2 := rand.Perm(10000)

	go BubbleSort(numbers2, channel)
	go InsertionSort(numbers2, channel)
	go SelectionSort(numbers2, channel)

	for i := 0; i < 3; i++ {
		lecture = append(lecture, <-channel)
	}
}
