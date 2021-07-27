package Routes
import (
	"github.com/freshers-bootcamp/Day4/Controllers"
	"github.com/gin-gonic/gin"
)
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/retail-service")
	{
		grp1.GET("product", Controllers.GetProduct)
		grp1.POST("product", Controllers.AddProduct )
		grp1.GET("product/:product_id", Controllers.GetProductByID )
		grp1.PATCH("product/:product_id", Controllers.PatchProduct )
		grp1.DELETE("product/:product_id", Controllers.DeleteProduct)
		grp1.GET("order",Controllers.GetOrders )
	}
	grp2:=r.Group("/customer")
	{
		grp2.GET("product", Controllers.GetProduct)
		grp2.POST("/",Controllers.AddCustomer)
		grp2.POST("/order",Controllers.PlaceOrder )
		grp2.GET("/:customer_id", Controllers.GetCustomerByID )
		grp2.DELETE("/:customer_id", Controllers.GetCustomerByID)
		grp2.GET("order/:customer_id",Controllers.GetOrderByID)
	}
	return r
}