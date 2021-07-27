package main

import (
	"fmt"
	"github.com/freshers-bootcamp/Day3/Exercise2/Config"
	"github.com/freshers-bootcamp/Day3/Exercise2/Models"
	"github.com/freshers-bootcamp/Day3/Exercise2/Routes"
	"github.com/jinzhu/gorm"
)
var err error
func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	r.Run()
}