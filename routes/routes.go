package routes

import (
	"blogsite_server/handlers"
	"blogsite_server/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/",handlers.Ping)

	api:=router.Group("/api")
	{
		api.POST("/users/register",handlers.RegisterUser)
		api.POST("/users/login",handlers.LoginUser)
		api.GET("/users/:id",handlers.GetUserProfile)

		api.GET("/validate", middleware.RequireAuth, handlers.Validate)

		api.GET("/posts",handlers.GetAllPosts)
		api.POST("/posts",middleware.RequireAuth,handlers.CreatePost)
		api.GET("/posts/:id",handlers.GetPostByID)
		api.PUT("/posts/:id",middleware.RequireAuth,handlers.UpdatePost)
		api.DELETE("/posts/:id",middleware.RequireAuth,handlers.DeletePost)

		api.GET("/search",handlers.SearchPosts)
	}
}