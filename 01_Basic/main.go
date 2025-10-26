package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func main() {
	fmt.Println("Addition:", add(5, 3))
	fmt.Println("Subtraction:", subtract(5, 3))
}
