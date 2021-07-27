package Models

import (
	"github.com/freshers-bootcamp/Day4/Config"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllOrders(order *[]Order ) (err error) {
	if err = Config.DB.Find(order).Error; err != nil {
		return err
	}
	return nil
}

func CreateOrder(order *Order)(err error){
	if err = Config.DB.Model(&order).Create(order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrdersOfCustomer(order *Order,id string) (err error){
	if err=Config.DB.Model(&order).Where("customer_id",id).Find(&order).Error ; err!=nil {
		return err
	}
	return nil
}


func GetOrderByID (order *Order, id string)(err error){
		if err = Config.DB.Model(&order).Where("id = ?", id).First(order).Error; err != nil {
			return err
		}
		return nil
	}

