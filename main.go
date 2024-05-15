package main

import (
	"encoding/json"
	"net/http"
	"sesi-3/app"
)

func productsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	category1 := app.Category{}
	category1.SetIDAndName(1, "Electronics")

	category2 := app.Category{}
	category2.SetIDAndName(2, "Books")

	product1 := app.NewProduct(1, "Laptop", 1000000, 2023, category1)
	product2 := app.NewProduct(2, "Book", 20000, 2023, category2)

	products := []app.Product{*product1, *product2}

	json.NewEncoder(w).Encode(products)
}

func main() {
	http.HandleFunc("/api/v1/products.json", productsHandler)
	http.ListenAndServe(":8081", nil)
}
