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
	r.Router.R.POST("/organization", r.Handler.Create)
	r.Router.R.GET("/organization/:id", r.Handler.QueryByID)
	r.Router.R.GET("/organization", r.Handler.QueryByName)
	r.Router.R.GET("/organisations", r.Handler.QueryAll)
	r.Router.R.DELETE("/organization/:id", r.Handler.Delete)
	r.Router.R.PUT("/organization/:id", r.Handler.Update)
}

func NewRoutes(router *utils.ServerStruct, handler *organisation.Handler) *Route {
	return &Route{
		Router:  router,
		Handler: handler,
	}
}
