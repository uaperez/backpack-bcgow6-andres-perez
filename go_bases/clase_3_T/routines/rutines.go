package main

import (
	"fmt"
	"time"
)

func process(i int) {
	fmt.Println("Iniciando proceso con", i)
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Finalizando proceso con", i)
}

func main() {
	now := time.Now()
	fmt.Println("Init program...")
	for i := 0; i < 100000; i++ {
		go process(i)
	}
	//time.Sleep(2000 * time.Millisecond)
	fmt.Println("Finished program...")
	fmt.Printf("El programa demoró %d ms\n", time.Now().Sub(now).Milliseconds())
}
