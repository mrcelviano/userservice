package main

import (
	"github.com/labstack/echo"
	"social-tech/userservice/commons"
	"social-tech/userservice/internal/http"
	"social-tech/userservice/internal/logic"
	"social-tech/userservice/internal/repository"
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
