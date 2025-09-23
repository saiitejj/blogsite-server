// handlers/users_handler.go
package handlers

import (
	"blogsite_server/database"
	"blogsite_server/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

// RegisterUser is a placeholder for the user registration logic.
func RegisterUser(c *gin.Context) {
	
	var body struct{
		Username string
		Email string
		Password string
	}
	
	if c.Bind(&body)!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request"})
		return
	}
	user:=models.User{
		Username: body.Username,
		Email: body.Email,
		Password: body.Password,
	}

	result:=database.DB.Create(&user)
	if result.Error!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"Handing user login"})

}

// LoginUser is a placeholder for the user login logic.
func LoginUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Handling user login"})
}

// GetUserProfile is a placeholder for fetching a user's profile.
func GetUserProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Handling fetching user profile"})
}