package Controllers
import (
	"fmt"
	"github.com/freshers-bootcamp/Day4/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProduct(c *gin.Context) {
	var user []Models.Product
	err := Models.GetAllProducts(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func AddProduct(c *gin.Context) {
	var user Models.Product
	c.BindJSON(&user)
	err := Models.AddProduct(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.Product
	err := Models.GetProductByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func PatchProduct(c *gin.Context) {
	var user Models.Product
	id := c.Params.ByName("id")
	err := Models.GetProductByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound,user)
	}
	err = c.BindJSON(&user)
	if err !=nil {
		return
	}
	err = Models.PatchProduct(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteProduct(c *gin.Context) {
	var user Models.Product
	id := c.Params.ByName("id")
	err := Models.DeleteProduct(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
