package main

import (
	"database/sql"
	"fmt"
	"log"
)

type product struct {
	ID int
	Name string
	Price float64
}

func main() {
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	createTable(db)

	addProduct(db, "Product1", 10.99)
	addProduct(db, "Product2", 20.99)

	products := getProducts(db)
	fmt.Println("List Of Product")
	for _, p := range products {
		fmt.Printf("ID : %d, Name : %s, Price : $%.2f\n", p.ID, p.Name, p.Price)
	}
}

func createTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS product (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(100)
			price DECIMAL(10,2)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func addProduct(db *sql.DB, name string, price float64) {
	_, err := db.Exec("INSERT INTO products (name, price) VALUES (?, ?)", name, price)
	if err != nil {
		log.Fatal(err)
	}
}

func getProducts(db *sql.DB) []product{
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []product
	for rows.Next() {
		var p product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, p)
	}
	return products
}