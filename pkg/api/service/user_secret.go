package service

import (
	"github.com/ahadiwasti/reacting-auth/pkg/api/dao"
	"github.com/ahadiwasti/reacting-auth/pkg/api/dto"
	"github.com/ahadiwasti/reacting-auth/pkg/api/log"
	"github.com/ahadiwasti/reacting-auth/pkg/api/model"
)

var userSecretDao = dao.UserSecretDao{}

// DomainService
type UserSecretService struct {
}

// InfoOfId - get role info by id
func (us UserSecretService) InfoOfId(dto dto.GeneralGetDto) model.UserSecret {
	return userSecretDao.Get(dto.Id)
}

// Create - create a new domain
func (us UserSecretService) Create(dto dto.UserSecretCreateDto) model.UserSecret {
	userSecret := model.UserSecret{
		User_id:      dto.User_id,
		Account_name: dto.Account_name,
		Secret:       dto.Secret,
	}
	c := userSecretDao.Create(&userSecret)
	if c.Error != nil {
		log.Error(c.Error.Error())
	}
	return userSecret
}
