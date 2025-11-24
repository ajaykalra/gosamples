package main

import (
	"database/sql"
	"fmt"
	"log"

	// "encoding/json"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "products"
)

type Product struct {
	ID          int
	ProductName string
	Price       float64
	Quantity    int
}

func main() {
	http.HandleFunc("/products", productsHandler)
	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Hello, World!"}`)

	// var products []Product
	// var err error
	// products, err = getProducts()
	// fmt.Println("Products:", products, "error:", err)
}

func getProducts() ([]Product, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Use the connection string to open a DB handle.
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("sql.Open: %v", err)
		return nil, err
	}
	defer db.Close()

	// Verify connection with Ping.
	if err := db.Ping(); err != nil {
		log.Fatalf("db.Ping: %v", err)
		return nil, err
	}

	fmt.Println("Connected to Postgres")

	rows, err := db.Query("SELECT product_id , product_name, price, stock_quantity FROM products")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.ProductName, &product.Price, &product.Quantity); err != nil {
			log.Fatal(err)
		}
		fmt.Println(product.ID, product.ProductName, product.Price, product.Quantity)
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println("Products:", products)

	return products, nil
}
