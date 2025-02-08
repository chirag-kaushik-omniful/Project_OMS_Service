package routes

import (
	"oms/controllers"

	"github.com/gin-gonic/gin"
	"github.com/omniful/go_commons/http"
)

func GetRouter(Router *http.Server) {
	api := Router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Hello"})
		})

		// Orders Group
		orders := api.Group("/orders")
		{
			orders.GET("/view", controllers.ViewOrders)
			orders.POST("/createBulk", controllers.CreateBulkOrder)
		}
	}
}
