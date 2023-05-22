package controllers

// import (
// 	"github.com/gin-gonic/gin"
// 	"./pkg/api/dto"
// 	"./pkg/api/service"
// )

// var deviceService = service.DeviceService{}

// type DeviceController struct {
// 	BaseController
// }

// // @Tags Role
// // @Summary 角色列表[分页+搜索]
// // @Security ApiKeyAuth
// // @Param limit query int false "条数"
// // @Param skip query int false "偏移量"
// // @Produce  json
// // @Success 200 {string} json "{"code":200,"data":{"result":[...],"total":1}}"
// // @Router /v1/roles [get]
// func (d *DeviceController) List(c *gin.Context) {
// 	var listDto dto.GeneralListDto
// 	if d.BindAndValidate(c, &listDto) {
// 		data, total := deviceService.List(listDto)
// 		resp(c, map[string]interface{}{
// 			"result": data,
// 			"total":  total,
// 		})
// 	}
// }

// // @Tags Role
// // @Summary 新增角色
// // @Security ApiKeyAuth
// // @Produce  json
// // @Success 200 {string} json "{"code":200,"data":{"id":1}}"
// // @Router /v1/roles [post]
// func (d *DeviceController) Create(c *gin.Context) {
// 	var roleDto dto.RoleCreateDto
// 	if d.BindAndValidate(c, &roleDto) {
// 		newRole, err := deviceService.Create(roleDto)
// 		if err != nil {
// 			ErrAddFail.Moreinfo = err.Error()
// 			fail(c, ErrAddFail)
// 			ErrAddFail.Moreinfo = ""
// 			return
// 		}

// 		resp(c, map[string]interface{}{
// 			"result": newRole,
// 		})
// 	}
// }

// // @Summary 更新角色信息
// // @Tags Role
// // @Security ApiKeyAuth
// // @Produce  json
// // @Success 200 {string} json "{"code":200,"data":{"result":[...],"total":1}}"
// // @Router /v1/roles/:id [put]
// // Edit - u of crud
// func (d *DeviceController) Edit(c *gin.Context) {
// 	var roleDto dto.RoleEditDto
// 	if d.BindAndValidate(c, &roleDto) {
// 		affected := deviceService.Update(roleDto)
// 		if affected < 0 {
// 			fail(c, ErrNoRecord)
// 			return
// 		}
// 		ok(c, "ok.UpdateDone")
// 	}
// }

// // @Summary 删除角色信息
// // @Tags Role
// // @Security ApiKeyAuth
// // @Produce  json
// // @Success 200 {string} json "{"code":200,"data":{"result":[...],"total":1}}"
// // @Router /v1/roles/:id [delete]
// // Delete - d of crud
// func (d *DeviceController) Delete(c *gin.Context) {
// 	var roleDto dto.GeneralDelDto
// 	if d.BindAndValidate(c, &roleDto) {
// 		if deviceService.Delete(roleDto) < 1 {
// 			fail(c, ErrNoRecord)
// 			return
// 		}
// 		ok(c, "ok.DeleteDone")
// 	}
// }
