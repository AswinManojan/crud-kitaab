package organisation

import (
	"github.com/gin-gonic/gin"
	"github.com/sample-crud-app/controllers/organisation"
)

type Route struct {
	Router *gin.Engine
}

var organisationController *organisation.OrganisationController

func (r Route) Routes() {
	r.Router.POST("/organizations", organisationController.Create)
	r.Router.GET("/organizations/:id", organisationController.QueryByID)
	r.Router.GET("/organization", organisationController.QueryByName)
	r.Router.GET("/organizations", organisationController.QueryAll)
	r.Router.DELETE("/organizations/:id", organisationController.Delete)
	r.Router.PUT("/organizations/:id", organisationController.Update)
}

func NewRoutes(router *gin.Engine) *Route {
	return &Route{
		Router: router,
	}
}
