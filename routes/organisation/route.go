package organisation

import (
	"github.com/sample-crud-app/controllers/organisation"
	"github.com/sample-crud-app/utils"
)

type Route struct {
	Router  *utils.ServerStruct
	Handler *organisation.Handler
}

func (r Route) Routes() {
	r.Router.R.POST("/createorganization", r.Handler.Create)
	r.Router.R.GET("/getorganizationbyid/:id", r.Handler.GetByID)
	r.Router.R.GET("/getorganizationbyname", r.Handler.GetByName)
	r.Router.R.DELETE("/deleteorganizationbyid/:id", r.Handler.DeleteByID)
	r.Router.R.PUT("/updateorganization/:id", r.Handler.Update)
}

func NewRoutes(router *utils.ServerStruct, handler *organisation.Handler) *Route {
	return &Route{
		Router:  router,
		Handler: handler,
	}
}
