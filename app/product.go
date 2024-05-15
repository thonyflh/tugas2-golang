package app

import (
	"encoding/xml"
)

type Product struct {
	Base
	Price    int64    `json:"price" xml:"price"`
	year     int      // Private access, will not be serialized to JSON/XML
	Category Category `json:"category" xml:"category"`
	XMLName  xml.Name `xml:"product" json:"-"`
}

func NewProduct(id int, name string, price int64, year int, category Category) *Product {
	return &Product{
		Base: Base{
			ID:   id,
			Name: name,
		},
		Price:    price,
		year:     year,
		Category: category,
	}
}

func (p *Product) GetYear() int {
	return p.year
}

func (p *Product) SetYear(year int) {
	p.year = year
}
