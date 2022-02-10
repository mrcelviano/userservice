package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mrcelviano/userservice/commons"
	"github.com/mrcelviano/userservice/internal/http"
	"github.com/mrcelviano/userservice/internal/logic"
	"github.com/mrcelviano/userservice/internal/repository"
	"github.com/mrcelviano/userservice/internal/rpc"
)

func main() {
	env := commons.GetEnvVar()
	commons.ConfigInit("configs/" + env + "_setting.json")
	pg := commons.InitGocraftDBRConnectionPG()

	//repo
	var (
		userRepo     = repository.NewUserRepository()
		notification = rpc.NewNotificationGRPCRepository()
	)

	//logic
	var (
		userLogic = logic.NewUserLogic(userRepo, notification)
	)

	//http
	e := echo.New()
	e.Pre(
		middleware.AddTrailingSlash(),
		commons.HTTPDBRSessionPG(pg),
	)
	http.NewUserHandlers(e.Group("api"), userLogic)

	commons.NewSignalHandler(e)
	commons.StartHttp(e, 8080)
	select {}
}
