package main

import (
	"blogsite_server/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.Default()

	routes.SetupRoutes(router)
	router.Run(":8000")
}
