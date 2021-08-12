package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	go process()
	time.Sleep(time.Millisecond * 1000) // Esto es malo, no lo hagas!
	fmt.Println("done")
}

func process() {
	fmt.Println("processing")
}
