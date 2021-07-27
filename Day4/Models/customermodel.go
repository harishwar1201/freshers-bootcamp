package Models
type Customer struct{
	CustomerId  string `json:"customer_id" gorm:"primary_key" `
	FirstName   string `json:"first_name"`
	LastName    string  `json:"last_name"`
	Number      string  `json:"number"`
}
func (b *Customer) TableName() string{
	return "customer"
}