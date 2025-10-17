package main

import (
	"fmt"
	"procaskill/db"
	_ "procaskill/docs"
	"procaskill/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db.ConnectDB()

	router := gin.Default()

	// Route list:
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.MainRouter(router)
	routes.UserRoutes(router)
	routes.TaskRoutes(router)

	var baseUri = fmt.Sprintf("%v:%v", os.Getenv("BASE_URI"), os.Getenv("PORT"))
	fmt.Println("Servidor ejecutandose en: " + baseUri)
	server := router.Run(baseUri)

	log.Fatal(server)
}
