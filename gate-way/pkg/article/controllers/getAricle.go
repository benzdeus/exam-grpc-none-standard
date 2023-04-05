package controllers

import (
	"gate-way/pkg/utils"

	"github.com/labstack/echo/v4"
)

func (c *controller) GetArticle(ctx echo.Context) error {
	// get token from header
	token := ctx.Request().Header.Get("Authorization")

	if token == "" {
		return ctx.JSON(401, map[string]string{
			"error": "token is required",
		})
	}

	// checkPermission
	permission := utils.CheckPermission(token)

	return ctx.JSON(200, map[string]string{
		"permission": permission,
	})

}
