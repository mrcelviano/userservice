package commons

import (
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
)

func HTTPDBRSessionPG(db *dbr.Connection) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx := req.Context()
			ctx = DBSessionNewContext(ctx, db)
			c.SetRequest(req.WithContext(ctx))
			return next(c)
		}
	}
}
