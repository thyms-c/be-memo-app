package utils

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/thyms-c/be-memo-app/internal/customerror"
)

func GetTokenFromHeader(c echo.Context) (string, error) {
	bearer := c.Request().Header.Get("Authorization")

	bearer = strings.TrimSpace(bearer)
	splittedToken := strings.Split(bearer, "Bearer ")
	if len(splittedToken) != 2 {
		return "", customerror.ErrNoToken
	}

	token := splittedToken[1]

	return token, nil
}
