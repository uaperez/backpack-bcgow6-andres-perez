package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	ctx := context.Background()

	ctx2 := context.WithValue(ctx, "Clave", "Valor")

	//MyFunction(ctx2, 10)

	//fmt.Printf("%+v", ctx)

	//ctx3, _ := context.WithDeadline(ctx2, time.Now().Add(5*time.Second))

	//<-ctx3.Done()

	//fmt.Println(ctx3.Err().Error())

	//ctx3, cancel := context.WithTimeout(ctx2, 5*time.Second)

	MyFunction(ctx2, 10)

	<-ctx2.Done()

	fmt.Println(ctx2.Err().Error())
}

func MyFunction(ctx context.Context, dato int) {
	valor := ctx.Value("Clave")

	str, ok := valor.(string)
	if !ok {
		fmt.Println("no es un string")
		os.Exit(1)
	}

	fmt.Printf("%T", str)

	_, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()
}
