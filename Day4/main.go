package main

import (
	"fmt"
	"github.com/freshers-bootcamp/Day4/Config"
	"github.com/freshers-bootcamp/Day4/Models"
	"github.com/freshers-bootcamp/Day4/Routes"
	"github.com/jinzhu/gorm"
)
var err error
func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Product{},&Models.Order{},&Models.Customer{})
	r := Routes.SetupRouter()
	r.Run()
}