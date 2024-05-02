package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sample-crud-app/controllers/user"
)

type Route struct {
	Router *gin.Engine
}

var userController *user.UserController

func (r Route) Routes() {
	r.Router.POST("/users", userController.Create)
	r.Router.GET("/users/:id", userController.QueryByID)
	r.Router.GET("/users", userController.QueryAll)
	r.Router.DELETE("/users/:id", userController.Delete)
	r.Router.PUT("/users/:id", userController.Update)
}

func NewRoutes(router *gin.Engine) *Route {
	return &Route{
		Router: router,
	}
}
