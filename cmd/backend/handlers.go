package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// GetUserHandler will return user based on user http input
func GetUserHandler(c *gin.Context) {
	var (
		user User
		err  error
	)

	// Get the user query parameter(s)
	user.ID = c.Param("id")

	// Get the user from the database
	user, err = user.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

// AddUserHandler adds a new User to the User list
func AddUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

// DeleteUserHandler will delete a specified user based on user http input
func DeleteUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// UpdateUserHandler will update a user
func UpdateUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
