package Models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	ProductID    int `json:"product_id"`
	CustomerId   string  `json:"customer_id"`
	Quantity     int `json:"quantity"`
	Status       string `json:"status"`
}
func (b *Order) TableName() string {
	return "order"
}

