package main

import (
	"github.com/labstack/echo"
	"github.com/mrcelviano/userservice/commons"
	"github.com/mrcelviano/userservice/internal/http"
	"github.com/mrcelviano/userservice/internal/logic"
	"github.com/mrcelviano/userservice/internal/repository"
)

func main() {
	env := commons.GetEnvVar()
	commons.ConfigInit("configs/" + env + "_setting.json")

	//repo
	var (
		userRepo = repository.NewUserRepository()
	)

	//logic
	var (
		userLogic = logic.NewUserLogic(userRepo)
	)

	//http
	e := echo.New()
	http.NewUserHandlers(e.Group("api"), userLogic)

	commons.NewSignalHandler(e)
	commons.StartHttp(e, 8080)
	select {}
}
