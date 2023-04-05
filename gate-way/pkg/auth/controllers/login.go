package controllers

import (
	"gate-way/pkg/pb"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *controller) Login(ctx echo.Context) error {

	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	req := new(request)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	res, err := c.AuthService.Login(ctx.Request().Context(), &pb.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(int(res.Code), res)

}
