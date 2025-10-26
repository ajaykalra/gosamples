package main

import (
	"fmt"
)

func main() {
	var a int = 5
	defer deferredFunction(a)
	fmt.Printf("after setting defer call: %d\n", a)
	a = 10
	fmt.Printf("value before panic %d\n", a)
	// Uncommenting the next line will cause a panic
	panic("Something went wrong!")
	a = 100 // This line will never be executed as panic occurs before it

	//fmt.Println("This will not run if there's a panic.")
}

func deferredFunction(x int) {
	fmt.Println("deferredFunction called.. Value of a in deferredFunction:", x)

	if r := recover(); r != nil {
		fmt.Println("Recovered in deferredFunction", r)
	}
}
