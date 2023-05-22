package dao

import (
	"github.com/ahadiwasti/reacting-auth/pkg/api/model"
	"github.com/jinzhu/gorm"
)

type UserSecretDao struct {
}

func (u UserSecretDao) Get(uid int) model.UserSecret {
	var userSecret model.UserSecret
	db.Where("user_id = ?", uid).First(&userSecret)
	return userSecret
}

func (u UserSecretDao) Create(UserSecret *model.UserSecret) *gorm.DB {
	db := GetDb()
	return db.Save(UserSecret)
}

// Update - update UserSecret
func (u UserSecretDao) Update(UserSecret *model.UserSecret, ups map[string]interface{}) *gorm.DB {
	db := GetDb()
	return db.Model(UserSecret).Update(ups)
}
