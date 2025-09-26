// handlers/users_handler.go
package handlers

import (
	"blogsite_server/database"
	"blogsite_server/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

// RegisterUser is a placeholder for the user registration logic.
func RegisterUser(c *gin.Context) {
	
	var body struct{
		Username string
		Email string
		Password string
	}
	
	if c.Bind(&body)!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}
	hash,err:=bcrypt.GenerateFromPassword([]byte(body.Password),10)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Failed to hash password"})
		return 
	}
	user:=models.User{
		Username: body.Username,
		Email: body.Email,
		Password: string(hash),
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
	var body struct{
		Email string
		Password string
	}

	if c.Bind(&body)!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Failed to read request"})
		return
	}
	var user models.User
	database.DB.First(&user,"email=?",body.Email)

	if user.ID==0{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(body.Password))

	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"sub":user.ID,
		"exp":time.Now().Add(time.Hour*24*30).Unix(),
	})
	
	tokenString,err:=token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})



	// c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in"})
}

func Validate(c *gin.Context) {
	// Get the user from the context (attached by RequireAuth middleware)
	user, _ := c.Get("user")

	// Respond with the user's information
	c.JSON(http.StatusOK, gin.H{
		"message": "I am logged in",
		"user":    user,
	})
}

// GetUserProfile is a placeholder for fetching a user's profile.
func GetUserProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Handling fetching user profile"})
}