package dao

import (
	"fmt"
	"strings"

	"./pkg/api/dto"
	"./pkg/api/model"
	"github.com/jinzhu/gorm"
)

type Role struct {
}

//Get - get single roel info
func (u Role) Get(id int, preload bool) model.Role {
	var role model.Role
	db := GetDb()
	if preload {
		db = db.Preload("Domain")
	}
	db.Where("id = ?", id).First(&role)
	return role
}

// GetRolesByIds
func (Role) GetRolesByIds(ids string) []model.Role {
	var roles []model.Role
	db := GetDb()
	db.Where("id in (?)", strings.Split(ids, ",")).Find(&roles)
	return roles
}

// GetRolesByNames
func (Role) GetRolesByNames(names []string) []model.Role {
	var roles []model.Role
	db := GetDb()
	db = db.Preload("Domain")
	db.Where("role_name in (?)", names).Find(&roles)
	return roles
}

//Get - get single roel infoD
func (u Role) GetByName(name string) model.Role {
	var role model.Role
	db.Where("role_name = ?", name).Preload("Domain").First(&role)
	return role
}

// List - users list
func (u Role) List(listDto dto.GeneralListDto) ([]model.Role, int64) {
	var roles []model.Role
	var total int64
	db := GetDb()
	for sk, sv := range dto.TransformSearch(listDto.Q, dto.RoleListSearchMapping) {
		db = db.Where(fmt.Sprintf("%s = ?", sk), sv)
	}
	db.Preload("Domain").Find(&roles)
	db.Model(&model.Role{}).Count(&total)
	return roles, total
}

// Create - new role
func (r Role) Create(role *model.Role) *gorm.DB {
	var row model.Role
	db := GetDb()
	db.Where("name = ? or role_name = ?", role.Name, role.RoleName).First(&row)
	if row.Id > 0 {
		return nil
	}
	return db.Create(role)
}

// Update - update role
func (r Role) Update(role *model.Role, ups map[string]interface{}) *gorm.DB {
	db := GetDb()
	return db.Model(role).Update(ups)
}

// Delete - delete role
func (r Role) Delete(role *model.Role) *gorm.DB {
	db := GetDb()
	return db.Delete(role)
}
