package organisation

import (
	"github.com/gin-gonic/gin"
	"github.com/sample-crud-app/controllers/organisation"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Route struct {
	Router *gin.Engine
}

var organisationController *organisation.OrganisationController

func (r Route) Routes() {
	r.Router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Router.POST("/organizations", organisationController.Create)
	r.Router.GET("/organizations/:Organisation-ID", organisationController.QueryByID)
	r.Router.GET("/organization", organisationController.QueryByName)
	r.Router.GET("/organizations", organisationController.QueryAll)
	r.Router.DELETE("/organizations/:Organisation-ID", organisationController.Delete)
	r.Router.PUT("/organizations/:Organisation-ID", organisationController.Update)
}

func NewRoutes(router *gin.Engine) *Route {
	return &Route{
		Router: router,
	}
}
