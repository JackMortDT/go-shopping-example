package main

import "fmt"

func main() {
	scores := []int{1, 2, 3, 4, 5}
	scores=  removeAtIndex(scores, 2)
	fmt.Println(scores) // [1, 2, 5, 4]
}

// No preserva el orden
func removeAtIndex(source[]int, index int) []int {
	lastIndex := len(source) - 1
	// Intercambia el último valor y el valor que queremos eliminar
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}