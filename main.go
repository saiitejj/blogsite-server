package main

import (
	"blogsite_server/routes"
	"blogsite_server/database"
	"log"
	"os" // for env variables


	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	err:=godotenv.Load()
	if err!=nil{
		log.Fatal("Error loading .env file")
	}
	database.ConnectToDB()
	router:=gin.Default()

	routes.SetupRoutes(router)
	port:=os.Getenv("PORT")
	if port==""{
		port="8080"
	}
	router.Run(":"+port)
}
