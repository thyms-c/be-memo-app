package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thyms-c/be-memo-app/internal/requests"
	"github.com/thyms-c/be-memo-app/internal/services"
	"github.com/thyms-c/be-memo-app/internal/utils"
)

type MemoHandler interface {
	GetAllMemos(c echo.Context) error
	CreateMemo(c echo.Context) error
	GetMemoByUserType(c echo.Context) error
}

type memoHandlerImpl struct {
	memoService services.MemoService
}

func NewMemoHandler(memoService services.MemoService) MemoHandler {
	return &memoHandlerImpl{
		memoService: memoService,
	}
}

func (h *memoHandlerImpl) GetAllMemos(c echo.Context) error {
	memos, err := h.memoService.GetAllMemos(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, memos)
}

func (h *memoHandlerImpl) CreateMemo(c echo.Context) error {
	req := new(requests.MemoRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request",
		})
	}

	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
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

	memo, err := h.memoService.CreateMemo(c.Request().Context(), req, userType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to create memo",
		})
	}
	return c.JSON(http.StatusCreated, memo)
}

func (h *memoHandlerImpl) GetMemoByUserType(c echo.Context) error {

	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
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

	memos, err := h.memoService.GetMemoByUserType(c.Request().Context(), string(userType))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to retrieve memos",
		})
	}
	return c.JSON(http.StatusOK, memos)
}
