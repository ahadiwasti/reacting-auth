package login

import (
	"github.com/ahadiwasti/reacting-auth/pkg/api/domain/account"
	"github.com/ahadiwasti/reacting-auth/pkg/api/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

type verifyCases struct {
	sourcePwd string
	targetPwd string
	salt      string
	expected  bool
}

func TestVerifyPassword(t *testing.T) {
	salt, _ := account.MakeSalt()
	pwd, _ := account.HashPassword("zeus", salt)
	for _, cs := range []verifyCases{{
		"zeus", pwd, salt, true,
	}, {
		"zeus0", pwd, salt, false,
	}, {
		"zeus", pwd, "wrong", false,
	}, {
		"zeus", "wrong", salt, false,
	},
	} {
		assert.Equal(t, cs.expected, VerifyPassword(cs.sourcePwd, model.User{
			Password: cs.targetPwd,
			Salt:     cs.salt,
		}), "Cases of password verification")
	}
}
