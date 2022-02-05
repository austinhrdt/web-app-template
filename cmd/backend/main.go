package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// set up database connection
	db := Init(os.Getenv("DATABASE_URL"))
	defer db.Close()

	// set up our router
	router := gin.Default()
	router.Use(cors.Default())

	// set up our routes
	router.GET("/health", HealthCheckHandler)
	router.GET("users/", GetUsersHandler)
	router.GET("users/:id", GetUserHandler)
	router.POST("users/", AddUserHandler)
	router.DELETE("users/:id", DeleteUserHandler)
	router.PUT("users/:id", UpdateUserHandler)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
