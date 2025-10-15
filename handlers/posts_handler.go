// handlers/posts_handler.go
package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"blogsite_server/database"
	"blogsite_server/models"
)

// Ping is a simple health-check handler.
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Server is alive!"})
}

// CreatePost is a placeholder for creating a new post.
func CreatePost(c *gin.Context) {

	var body struct{
		Title string
		Content string
	}
	if c.Bind(&body)!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}
	user,exists:=c.Get("user")
	if !exists{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	post := models.Post{
		Title:   body.Title,
		Content: body.Content,
		UserID:  user.(models.User).ID, // Get the user ID from the authenticated user
	}

	// 4. Save to the database
	result := database.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create post"})
		return
	}

	// 5. Return the created post
	c.JSON(http.StatusOK, gin.H{"post": post})
}

// GetAllPosts is a placeholder for getting all posts.
func GetAllPosts(c *gin.Context) {
	var posts []models.Post

	result := database.DB.Preload("User").Find(&posts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

// GetPostByID is a placeholder for getting a single post by its ID.
func GetPostByID(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	result := database.DB.Preload("User").First(&post, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}
// UpdatePost is a placeholder for updating a post.
func UpdatePost(c *gin.Context) {
	id:=c.Param("id")
	var body struct{
		Title string
		Content string
	}
	c.Bind(&body)

	var post models.Post
	database.DB.First(&post,id)
	if post.ID==0{
		c.JSON(http.StatusNotFound,gin.H{"error":"Post not found"})
		return
	}
	user,_:=c.Get("user")
	if post.UserID!=user.(models.User).ID{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"You are not authorized"})
		return
	}
	database.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Content: body.Content,
	})

	c.JSON(http.StatusOK, gin.H{"post":post})
}

// DeletePost is a placeholder for deleting a post.
func DeletePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	database.DB.First(&post, id)
	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	user, _ := c.Get("user")

	if post.UserID != user.(models.User).ID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to delete this post"})
		return
	}

	database.DB.Delete(&models.Post{}, id)

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}


// SearchPosts is a placeholder for searching posts.
func SearchPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Handling post search"})
}