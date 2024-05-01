package routes

import (
	"github.com/sample-crud-app/controllers/organisation"
	"github.com/sample-crud-app/utils"
)

type Route struct {
	Router  *utils.ServerStruct
	Handler *controllers.OrganizationHandler
}

func (r Route) Routes() {
	r.Router.R.POST("/createorganization", r.Handler.CreateOrganizationHandler)
	r.Router.R.GET("/getorganizationbyid/:id", r.Handler.GetOrganizationByIDHandler)
	r.Router.R.GET("/getorganizationbyname", r.Handler.GetOrganizationByNameHandler)
	r.Router.R.DELETE("/deleteorganizationbyid/:id", r.Handler.DeleteOrganizaionByIDHandler)
	r.Router.R.PATCH("/updateorganization/:id", r.Handler.UpdateOrganizationHandler)
}

func NewRoutes(router *utils.ServerStruct, handler *controllers.OrganizationHandler) *Route {
	return &Route{
		Router:  router,
		Handler: handler,
	}
}
