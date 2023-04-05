package controllers

import (
	"gate-way/pkg/pb"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *controller) Register(ctx echo.Context) error {

	type request struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		Permission string `json:"permission"`
	}

	req := new(request)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	res, err := c.AuthService.Register(ctx.Request().Context(), &pb.RegisterRequest{
		Username:   req.Username,
		Password:   req.Password,
		Permission: req.Permission,
	})

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(int(res.Code), res)

}
