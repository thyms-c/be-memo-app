package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thyms-c/be-memo-app/internal/services"
	"github.com/thyms-c/be-memo-app/internal/utils"
)

type CounterHandler interface {
	GetCounter(c echo.Context) error
}
type counterHandlerImpl struct {
	counterService services.CounterService
}

func NewCounterHandler(counterService services.CounterService) CounterHandler {
	return &counterHandlerImpl{
		counterService: counterService,
	}
}
func (h *counterHandlerImpl) GetCounter(c echo.Context) error {
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		log.Println("error", err)
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "Unauthorized",
		})
	}

	userType, err := utils.GetUserTypeByToken(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "Unauthorized",
		})
	}

	counter, err := h.counterService.GetCounterByUserRole(c.Request().Context(), string(userType))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, counter)
}
