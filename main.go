package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sample-crud-app/routes/organisation"
	"github.com/sample-crud-app/routes/user"
	"github.com/sample-crud-app/utils"
	_ "github.com/sample-crud-app/docs"
)

var Route organisation.Route

// @title Kitaab User and Organisation Services API
// @version 1.0
// @description Kitaab training API in Go and Gin Framework

// @host localhost:8080
// @BasePath 
func main() {
	server := gin.Default()
	utils.DBConnect()
	OrgnRoutes := organisation.NewRoutes(server)
	UserRoutes := user.NewRoutes(server)
	OrgnRoutes.Routes()
	UserRoutes.Routes()
	server.Run(":8080")
}
