package routes
import (
	"github.com/sample-crud-app/controllers/user"
	"github.com/sample-crud-app/utils"
)

type Route struct {
	Router  *utils.ServerStruct
	Handler *handler.UserHandler
}

func (r Route) Routes() {
	r.Router.R.POST("/createuser", r.Handler.CreateUserHandler)
	r.Router.R.GET("/getuserbyid/:id", r.Handler.GetUserByIDHandler)
	r.Router.R.GET("/getuserbyname", r.Handler.GetUserByNameHandler)
	r.Router.R.GET("/getallusers", r.Handler.GetAllUsersHandler)

}

func NewRoutes(router *utils.ServerStruct, handler *handler.UserHandler) *Route {
	return &Route{
		Router:  router,
		Handler: handler,
	}
}
