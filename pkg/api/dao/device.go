package dao

// import (
// 	"fmt"
// 	"github.com/jinzhu/gorm"
// 	// "strings"
// 	"./pkg/api/dto"
// 	"./pkg/api/model"
// )

// type Device struct {
// }

// // List - users list
// func (d Device) List(listDto dto.GeneralListDto) ([]model.Device, int64) {
// 	var devices []model.Device
// 	var total int64
// 	db := GetDb()
// 	for sk, sv := range dto.TransformSearch(listDto.Q, dto.DeviceListSearchMapping) {
// 		db = db.Where(fmt.Sprintf("%s = ?", sk), sv)
// 	}
// 	db.Preload("Domain").Offset(listDto.Skip).Limit(listDto.Limit).Find(&devices)
// 	db.Model(&model.Device{}).Count(&total)
// 	return devices, total
// }

// // Create - new Device
// func (d Device) Create(device *model.Device) *gorm.DB {
// 	db := GetDb()
// 	return db.Create(device)
// }

// // Update - update role
// func (d Device) Update(device *model.Device, ups map[string]interface{}) *gorm.DB {
// 	db := GetDb()
// 	return db.Model(device).Update(ups)
// }

// // Delete - delete role
// func (d Device) Delete(device *model.Device) *gorm.DB {
// 	db := GetDb()
// 	return db.Delete(device)
// }
