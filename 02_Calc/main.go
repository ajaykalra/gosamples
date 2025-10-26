package main

import (
	calculator "calculator/Calculator"
	"fmt"
)

func main() {
	fmt.Println("Addition:", calculator.Add(2, 3))
	fmt.Println("Subtraction:", calculator.Subtract(5, 2))
	fmt.Println("Multiplication:", calculator.Multiply(4, 3))
}
