package main

import (
	"example/bookstore/database"
	"example/bookstore/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.Connect()

	routes.RegisterBookRoutes(router)

	router.Run(":8081")
}
