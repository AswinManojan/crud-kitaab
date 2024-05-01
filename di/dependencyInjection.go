package di

import (
	controllers "github.com/sample-crud-app/controllers/organisation"
	usercontrollers "github.com/sample-crud-app/controllers/user"
	"github.com/sample-crud-app/repositories/organisation"
	usr "github.com/sample-crud-app/repositories/user"
	routes "github.com/sample-crud-app/routes/organisation"
	userroutes "github.com/sample-crud-app/routes/user"
	orgsvc "github.com/sample-crud-app/services/organisation"
	usrsvc "github.com/sample-crud-app/services/user"
	"github.com/sample-crud-app/utils"
)

func Init() *utils.ServerStruct {
	// utils.DBConnect()
	// OrgnRepo := orgnrepo.NewRepoImpl()
	// OrgnSVC := orgnsvc.NewSVCImpl(OrgnRepo)
	// OrgnHandler := controllers.NewOrganizationHandler(OrgnSVC)
	// srvr := utils.NewServer()
	// OrgnRoutes := routes.NewRoutes(srvr, OrgnHandler)
	// OrgnRoutes.Routes()
	// return srvr
	utils.DBConnect()
	OrgnRepo := organisation.NewRepoImpl()
	UserRepo := usr.NewRepoImpl()
	OrgnSVC := orgsvc.NewSVCImpl(OrgnRepo)
	UserSVC := usrsvc.NewSVCImpl(UserRepo)
	orgnHandler := controllers.NewOrganizationHandler(OrgnSVC)
	userHandler := usercontrollers.NewUserHandler(UserSVC)
	srvr := utils.NewServer()
	OrgnRoutes := routes.NewRoutes(srvr, orgnHandler)
	UserRoutes := userroutes.NewRoutes(srvr, userHandler)
	OrgnRoutes.Routes()
	UserRoutes.Routes()
	return srvr
}
