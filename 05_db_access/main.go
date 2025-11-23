package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "products"
)

type Product struct {
	ID          int
	ProductName string
	Price       float64
	Quantity    int
}

// Simple Postgres example. Connection parameters can be provided via
// environment variables (PGHOST, PGPORT, PGUSER, PGPASSWORD, PGDB).
// Defaults are provided for local testing.
func main() {
	// host := getenv("PGHOST", "localhost")
	// port := getenv("PGPORT", "5432")
	// user := getenv("PGUSER", "admin")
	// password := getenv("PGPASSWORD", "admin")
	// dbname := getenv("PGDB", "products")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Use the connection string to open a DB handle.
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("sql.Open: %v", err)
	}
	defer db.Close()

	// Verify connection with Ping.
	if err := db.Ping(); err != nil {
		log.Fatalf("db.Ping: %v", err)
	}

	fmt.Println("Connected to Postgres")

	// Simple query to demonstrate use of the DB handle.
	var now string
	if err := db.QueryRow("SELECT NOW()::text").Scan(&now); err != nil {
		log.Fatalf("query NOW(): %v", err)
	}
	fmt.Printf("Postgres time: %s\n", now)

	rows, err := db.Query("SELECT product_id , product_name FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.ProductName); err != nil {
			log.Fatal(err)
		}
		fmt.Println(product.ID, product.ProductName)
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}

// func getenv(key, def string) string {
// 	if v := os.Getenv(key); v != "" {
// 		return v
// 	}
// 	return def
// }
