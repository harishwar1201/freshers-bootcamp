package Models

import (
	"github.com/freshers-bootcamp/Day4/Config"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllCustomers(customer *[]Customer ) (err error) {
	if err = Config.DB.Find(customer).Error; err != nil {
		return err
	}
	return nil
}

func AddCustomer(customer *Customer ) (err error) {
	if err = Config.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

func GetCustomerByID(customer *Customer , id string) (err error) {
	if err = Config.DB.Where("customer_id = ?", id).First(customer).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCustomer(customer *Customer  , id string) (err error) {
	Config.DB.Where("customer_id = ?", id).Delete(customer)
	return nil
}