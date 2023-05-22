package login

import (
	"github.com/ahadiwasti/reacting-auth/pkg/api/domain/account"
	"github.com/ahadiwasti/reacting-auth/pkg/api/model"
)

// VerifyPassword : verify password by salt
func VerifyPassword(password string, userModel model.User) bool {
	if pwd, err := account.HashPassword(password, userModel.Salt); err == nil && pwd == userModel.Password {
		return true
	}
	return false
}
