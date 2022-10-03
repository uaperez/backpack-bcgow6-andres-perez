package main

import (
	"fmt"
	"time"
)

func processing(i int, channel chan int) {
	fmt.Println("Iniciando proceso con", i)
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Finalizando proceso con", i)
	channel <- i
}

func main() {
	channel := make(chan int)
	now := time.Now()
	fmt.Println("Init program...")
	for i := 0; i < 10; i++ {
		go processing(i, channel)
	}
	for i := 0; i < 10; i++ {
		read := <-channel
		fmt.Println("Lectura del channel:", read)
	}
	//time.Sleep(2000 * time.Millisecond)
	fmt.Println("Finished program...")
	fmt.Printf("El programa demoró %d ms\n", time.Now().Sub(now).Milliseconds())
}
