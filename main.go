package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	go func() {
		fmt.Println("processing")
	}()
	time.Sleep(time.Millisecond * 1000) // Esto es malo, no lo hagas!
	fmt.Println("done")
}
