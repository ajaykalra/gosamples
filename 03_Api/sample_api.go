package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AddDto struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

func Add(a int, b int) int {
	return a + b
}

func main() {
	// Register a handler function for the "/hello" endpoint.
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/add", createAddHandler)

	// Start the HTTP server on port 8080.
	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// helloHandler is the function that handles requests to the "/hello" endpoint.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to indicate JSON response.
	w.Header().Set("Content-Type", "application/json")

	// Write a simple JSON response.
	fmt.Fprintf(w, `{"message": "Hello, World!"}`)
}

func createAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newPost AddDto
	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// In a real application, you would save newPost to a database
	// For this example, we'll just print it and return it
	fmt.Printf("Received add: %+v\n", newPost)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // Indicate successful creation
	//calculator.add(newPost.Num1, newPost.Num2))
	json.NewEncoder(w).Encode(Add(newPost.Num1, newPost.Num2))
}
