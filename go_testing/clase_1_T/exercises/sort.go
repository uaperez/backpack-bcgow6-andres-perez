package exercises

func InsertionSort(numbersArray []int) []int {
	numbers := make([]int, len(numbersArray))
	copy(numbers, numbersArray)
	arrayLength := len(numbers)
	for i := 1; i < arrayLength; i++ {
		key := numbers[i]
		j := i - 1

		for j >= 0 && numbers[j] > key {
			numbers[j+1] = numbers[j]
			j = j - 1
		}

		numbers[j+1] = key
	}
	return numbers
}
