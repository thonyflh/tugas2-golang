package main

import (
	"encoding/json"
	"encoding/xml"
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

func ProductHandlerXML(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/xml")

	var products []app.Product

	ecommerceCategory := app.Category{}
	ecommerceCategory.SetIDAndName(2, "Ecommerce")

	projectManagementCategory := app.Category{}
	projectManagementCategory.SetIDAndName(1, "Project management")

	product1 := app.NewProduct(1, "Lakugan", 1000000000, 2023, ecommerceCategory)
	product2 := app.NewProduct(2, "My score card", 2000000000, 2024, projectManagementCategory)
	products = append(products, *product1, *product2)

	xmlData, err := xml.MarshalIndent(products, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(xml.Header))
	w.Write(xmlData)
}

func AddProductHandlerJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newProduct app.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseData, err := json.Marshal(newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(responseData)
}

func main() {
	http.HandleFunc("/api/v1/products.json", productsHandler)
	http.HandleFunc("/api/v1/products.xml", ProductHandlerXML)
	http.HandleFunc("/api/v1/add-products.json", AddProductHandlerJSON)
	http.ListenAndServe(":8081", nil)
}
