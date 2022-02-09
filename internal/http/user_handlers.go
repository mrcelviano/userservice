package http

import (
	"github.com/labstack/echo"
	"github.com/mrcelviano/userservice/internal/app"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
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
	ctx := c.Request().Context()
	user := app.User{}
	err := c.Bind(&user)
	if err != nil {
		return errors.Wrap(err, "cant bind json body")
	}
	user, err = u.logic.Create(ctx, user)
	if err != nil {
		return errors.Wrap(err, "cant create new user")
	}
	return c.JSON(http.StatusCreated, user)
}

func (u *userHandlers) GetList(c echo.Context) error {
	ctx := c.Request().Context()
	p := app.Pagination{}
	err := c.Bind(&p)
	if err != nil {
		return errors.Wrap(err, "cant read query params")
	}
	resp, err := u.logic.GetList(ctx, p.WithDefaultSortKey("id"))
	if err != nil {
		return errors.Wrap(err, "cant get list")
	}
	return c.JSON(http.StatusOK, resp)
}

func (u *userHandlers) GetByID(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.ParseInt(c.Param("id"), 0, 64)
	if err != nil {
		return errors.Wrap(err, "cant convert id to int64")
	}
	field, err := u.logic.GetByID(ctx, id)
	if err != nil {
		return errors.Wrap(err, "cant get user")
	}
	return c.JSON(http.StatusOK, field)
}

func (u *userHandlers) Update(c echo.Context) error {
	ctx := c.Request().Context()
	user := app.User{}
	id, err := strconv.ParseInt(c.Param("id"), 0, 64)
	if err != nil {
		return errors.Wrap(err, "cant convert id to int64")
	}
	err = c.Bind(&user)
	if err != nil {
		return errors.Wrap(err, "cant bind json body")
	}
	user.ID = id
	user, err = u.logic.Update(ctx, user)
	if err != nil {
		return errors.Wrap(err, "cant update user")
	}
	return c.JSON(http.StatusOK, user)
}

func (u *userHandlers) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.ParseInt(c.Param("id"), 0, 64)
	if err != nil {
		return errors.Wrap(err, "cant convert id to int64")
	}
	err = u.logic.Delete(ctx, id)
	if err != nil {
		return errors.Wrap(err, "cant delete user")
	}
	return c.NoContent(http.StatusOK)
}
