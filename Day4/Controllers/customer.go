package Controllers

import (
	"fmt"
	"github.com/freshers-bootcamp/Day4/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCustomer(c *gin.Context) {
	var user []Models.Customer
	err := Models.GetAllCustomers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func AddCustomer(c *gin.Context) {
	var user Models.Customer
	c.BindJSON(&user)
	err := Models.AddCustomer(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetCustomerByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.Customer
	err := Models.GetCustomerByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
func DeleteCustomer(c *gin.Context) {
	var user Models.Customer
	id := c.Params.ByName("customer_id")
	err := Models.DeleteCustomer(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
