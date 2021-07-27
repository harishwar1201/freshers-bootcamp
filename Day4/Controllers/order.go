package Controllers

import (
	"fmt"
	"github.com/freshers-bootcamp/Day4/Config"
	"github.com/freshers-bootcamp/Day4/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetOrders(c *gin.Context) {
var user []Models.Order
err := Models.GetAllOrders(&user)
if err != nil {
c.AbortWithStatus(http.StatusNotFound)
} else {
c.JSON(http.StatusOK, user)
}
}
func CreateOrder(c *gin.Context){
	var user Models.Order
	c.BindJSON(&user)
	err := Models.CreateOrder(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
func GetOrderByID(c *gin.Context) {
	id := c.Params.ByName("order_id")
	var user Models.Order
	err := Models.GetOrderByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
func PlaceOrder(c *gin.Context){
	var order Models.Order
	err := c.BindJSON(&order)
	if err != nil {
		return
	}
	possible := IsOrderPossible(order,c)
	if possible == false {
		c.JSON(http.StatusOK, gin.H{
			"status":"order failed",
		})
		return
	}
	cooldown := isCoolDownOver()
	if cooldown == false{
		c.JSON(http.StatusOK,gin.H{
			"status":"cannot make another order request for next 1 minute",
		})
		return
	}
	order.Status="order successfully placed"
	err = Models.CreateOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id": order.ID,
			"product_id": order.ProductID,
			"quantity": order.Quantity,
			"status": "order  placed successfully",
		})
	}

}
func IsOrderPossible(order Models.Order, c *gin.Context) bool{
	id:=order.ProductID
	var product Models.Product
	err :=Models.GetProductByID(&product,strconv.Itoa(id))
	if err !=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}
	if product.Quantity < order.Quantity{
		return false
	}
	product.Quantity -= order.Quantity
	err = Models.PatchProduct(&product,strconv.Itoa(id))
	if err !=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}
	return true
}
func isCoolDownOver() bool{
	var order Models.Order
	Config.DB.Model(&order).Last(&order)
	if order.ID == 0{
		return true
	}
	currTime := time.Now()
	diffInTime := currTime.Sub(order.CreatedAt).Seconds()
	if diffInTime <= 60{
		return false
	}
	return true
}