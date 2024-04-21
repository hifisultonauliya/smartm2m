package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"smartm2m/src/controllers"
	"smartm2m/src/helper"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	client, err := helper.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
		return
	}
	defer client.Disconnect(context.Background())

	authController := controllers.NewAuthController()
	itemController := controllers.NewItemController()

	// Routes
	v1 := r.Group("/api/v1")
	{
		// User routes
		users := v1.Group("/users")
		{
			users.POST("/register", authController.Register)
			users.POST("/login", authController.Login)
		}

		// Item routes
		items := v1.Group("/items").Use(AuthMiddleware())
		{
			items.POST("/", itemController.CreateItem)
			items.POST("/:id/purchase", itemController.PurchaseItem)

			items.GET("/", itemController.GetItems)
			items.GET("/:id", itemController.GetItem)

			items.PUT("/:id", itemController.UpdateItem)

			items.DELETE("/:id", itemController.DeleteItem)

		}
	}

	r.Run(":8080")
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		claims, err := helper.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		fmt.Println("User ID:", claims.UserID)

		c.Next()
	}
}
