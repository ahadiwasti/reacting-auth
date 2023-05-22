package service

// import (
// 	"github.com/pkg/errors"
// 	// "strconv"
// 	// "strings"
// 	"github.com/ahadiwasti/reacting-auth/pkg/api/dao"
// 	"github.com/ahadiwasti/reacting-auth/pkg/api/dto"
// 	"github.com/ahadiwasti/reacting-auth/pkg/api/log"
// 	"github.com/ahadiwasti/reacting-auth/pkg/api/model"
// )

// var deviceDao = dao.Device{}

// // DeviceService
// type DeviceService struct {
// }

// // List - users list with pagination
// func (DeviceService) List(dto dto.GeneralListDto) ([]model.Device, int64) {
// 	return deviceDao.List(dto)
// }
// // Create - create a new device
// func (rs DeviceService) Create(dto dto.Device) (model.Device, error) {
// 	deviceModel := model.Device{
// 		iddevice:     	dto.iddevice,
// 		devicename:   	dto.devicename,
// 		deviceid:     	dto.deviceid,
// 		deviceserial:   dto.deviceserial,
// 		assettag:    	dto.assettag,
// 	}
// 	c := deviceDao.Create(&deviceModel)
// 	if c == nil {
// 		return model.Device{}, errors.New("Duplicated role")
// 	} else {
// 		if c.Error != nil {
// 			log.Error(c.Error.Error())
// 			return model.Device{}, c.Error
// 		}
// 	}
// 	return deviceModel, nil
// }

// // Update - update role's information
// func (rs DeviceService) Update(deviceDto dto.Device) int64 {
// 	c := deviceDao.Update(&model.Device{id: deviceDto.iddevice}, map[string]interface{}{
// 		"devicename":         deviceDto.devicename,
// 		"deviceid":       deviceDto.deviceid,
// 		"deviceserial":    deviceDto.deviceserial,
// 		"assettag":     deviceDto.assettag,
// 	})
// 	// rs.AssignPermission(deviceDto.id, deviceDto.Menuids)
// 	// if deviceDto.DataPermids != "" {
// 	// 	_ = rs.AssignDataPerm(deviceDto.id, deviceDto.DataPermids)
// 	// }

// 	return c.RowsAffected
// }

// // Delete - delete role
// func (rl DeviceService) Delete(dto dto.Device) int64 {
// 	deviceModel := deviceDao.Get(dto.iddevice, false)
// 	if deviceModel.iddevice < 1 {
// 		return -1
// 	}
// 	//1. delete role
// 	c := deviceDao.Delete(&deviceModel)

// 	return c.RowsAffected
// }
