package utils

import (
	"github.com/thyms-c/be-memo-app/internal/configs"
	"github.com/thyms-c/be-memo-app/internal/customerror"
	"github.com/thyms-c/be-memo-app/internal/models"
)

func GetUserTypeByToken(token string) (models.Role, error) {
	var userType models.Role
	configs := configs.NewConfig()

	if token == configs.UserToken {
		userType = models.UserRole
	} else if token == configs.AdminToken {
		userType = models.AdminRole
	} else {
		return "", customerror.ErrInvalidToken
	}

	return userType, nil
}
