package user
import (
	"github.com/sample-crud-app/controllers/user"
	"github.com/sample-crud-app/utils"
)

type Route struct {
	Router  *utils.ServerStruct
	Handler *user.Handler
}

func (r Route) Routes() {
	r.Router.R.POST("/user", r.Handler.Create)
	r.Router.R.GET("/user/:id", r.Handler.QueryByID)
	r.Router.R.GET("/user", r.Handler.QueryByName)
	r.Router.R.GET("/users", r.Handler.QueryAll)
	r.Router.R.DELETE("/user/:id",r.Handler.Delete)
	r.Router.R.PUT("/users/:id",r.Handler.Update)
}

func NewRoutes(router *utils.ServerStruct, handler *user.Handler) *Route {
	return &Route{
		Router:  router,
		Handler: handler,
	}
}
