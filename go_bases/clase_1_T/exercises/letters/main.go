package main

import "fmt"

func main() {
	word := "Bootcamp"
	countLetters := len(word)
	fmt.Printf("The word '%s' contains %v letters\n", word, countLetters)

	for i := 0; i < countLetters; i++ {
		fmt.Println(string(word[i]))
	}
}
