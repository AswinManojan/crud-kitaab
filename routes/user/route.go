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
	r.Router.R.POST("/createuser", r.Handler.Create)
	r.Router.R.GET("/getuserbyid/:id", r.Handler.GetByID)
	r.Router.R.GET("/getuserbyname", r.Handler.GetByName)
	r.Router.R.GET("/getallusers", r.Handler.GetAll)

}

func NewRoutes(router *utils.ServerStruct, handler *user.Handler) *Route {
	return &Route{
		Router:  router,
		Handler: handler,
	}
}
