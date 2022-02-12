package delivery

import (
	"github.com/labstack/echo"
	"github.com/mrcelviano/userservice/internal/domain"
	"github.com/mrcelviano/userservice/pkg/logger"
	"net/http"
	"strconv"
)

type responseError struct {
	Message string `json:"message"`
}

type userHandlers struct {
	logic domain.UserService
}

func NewUserHandlers(rg *echo.Group, logic domain.UserService) *echo.Group {
	u := userHandlers{logic: logic}

	rg.POST("/", u.Create)
	rg.GET("/", u.GetList)
	rg.GET("/:id/", u.GetByID)
	rg.PUT("/:id/", u.Update)
	rg.DELETE("/:id/", u.Delete)

	return rg
}

func (u *userHandlers) Create(c echo.Context) error {
	ctx := c.Request().Context()
	user := domain.User{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	user, err = u.logic.Create(ctx, user)
	if err != nil {
		return c.JSON(getStatusCode(err), responseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}

func (u *userHandlers) GetList(c echo.Context) error {
	ctx := c.Request().Context()
	request := domain.GetUserListRequest{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	resp, err := u.logic.GetList(ctx, request)
	if err != nil {
		return c.JSON(getStatusCode(err), responseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (u *userHandlers) GetByID(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.ParseInt(c.Param("id"), 0, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	field, err := u.logic.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), responseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, field)
}

func (u *userHandlers) Update(c echo.Context) error {
	ctx := c.Request().Context()
	user := domain.User{}
	id, err := strconv.ParseInt(c.Param("id"), 0, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	user.ID = id
	user, err = u.logic.Update(ctx, user)
	if err != nil {
		return c.JSON(getStatusCode(err), responseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func (u *userHandlers) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.ParseInt(c.Param("id"), 0, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	err = u.logic.Delete(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), responseError{Message: err.Error()})
	}
	return c.NoContent(http.StatusOK)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logger.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrBadParamInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
