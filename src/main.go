package main

import "fmt"

func main() {
  lookup := make(map[string]int)
  lookup["goku"] = 9001
  power, exists := lookup["vegeta"]
  power2, exists2 := lookup["goku"]

  // prints 0, false
  // 0 is the default value for an integer
  fmt.Println(power, exists)
  fmt.Println(power2, exists2)
}