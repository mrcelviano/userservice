package http

import (
	"github.com/labstack/echo"
	"github.com/mrcelviano/userservice/internal/app"
)

type userHandlers struct {
	logic app.UserLogic
}

func NewUserHandlers(rg *echo.Group, logic app.UserLogic) *echo.Group {
	u := userHandlers{logic: logic}

	rg.POST("/", u.Create)
	rg.GET("/", u.GetList)
	rg.GET("/:id", u.GetByID)
	rg.PUT("/:id", u.Update)
	rg.DELETE("/:id", u.Delete)

	return rg
}

func (u *userHandlers) Create(c echo.Context) error {
	return nil
}

func (u *userHandlers) GetList(c echo.Context) error {
	return nil
}

func (u *userHandlers) GetByID(c echo.Context) error {
	return nil
}

func (u *userHandlers) Update(c echo.Context) error {
	return nil
}

func (u *userHandlers) Delete(c echo.Context) error {
	return nil
}
