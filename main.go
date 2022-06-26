package main

import (
	"os"

	"github.com/bosokoli/go_restaurant_management/database"
	"github.com/bosokoli/go_restaurant_management/middleware"
	"github.com/bosokoli/go_restaurant_management/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("port")
	if port == "" {
		port = "8000"

		router := gin.New()
		router.Use(gin.Logger())

		routes.UserRoutes(router)
		router.Use(middleware.Authentication())

		routes.FoodRoutes(router)
		routes.OrderRoutes(router)
		routes.TableRoutes(router)
		routes.MenuRoutes(router)
		routes.OrderItemRoutes(router)
		routes.InvoiceRoutes(router)

		router.Run(":" + port)

	}
}
