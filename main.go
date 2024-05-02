package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sample-crud-app/routes/organisation"
	"github.com/sample-crud-app/routes/user"
	"github.com/sample-crud-app/utils"
)

var Route organisation.Route

func main() {
	server := gin.Default()
	utils.DBConnect()
	OrgnRoutes := organisation.NewRoutes(server)
	UserRoutes := user.NewRoutes(server)
	OrgnRoutes.Routes()
	UserRoutes.Routes()
	server.Run(":8080")
}
