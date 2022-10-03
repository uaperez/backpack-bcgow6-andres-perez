package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello world!")

	value, err := io.Copy(os.Stdout, reader)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(value))

}
