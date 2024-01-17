package main

import (
	"Rest_Api_Go/db"
	"Rest_Api_Go/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	fmt.Println("db successfully initialized!!!")
	server := gin.Default() //configures http server
	//pass pointer to server to register routes, saves realestate
	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost 8080

}
