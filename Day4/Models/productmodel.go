package Models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ProductName   string `json:"product_name"`
	Price  int `json:"price"`
	Quantity   int `json:"quantity"`
}
func (b *Product) TableName() string {
	return "product"
}