package controllers

import (
	"github.com/ahadiwasti/reacting-auth/pkg/api/dto"
	"github.com/ahadiwasti/reacting-auth/pkg/api/service"
	"github.com/gin-gonic/gin"
)

var domainService = service.DomainService{}

type DomainController struct {
	BaseController
}

// @Tags Domain
// @Summary 项目信息
// @Security ApiKeyAuth
// @Param id path int true "项目id"
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"id":1,"name":"test"}}"
// @Router /v1/domains/{id} [get]
func (d *DomainController) Get(c *gin.Context) {
	var gDto dto.GeneralGetDto
	if d.BindAndValidate(c, &gDto) {
		data := domainService.InfoOfId(gDto)
		//role not found
		if data.Id < 1 {
			fail(c, ErrNoUser)
			return
		}
		resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

// @Tags Domain
// @Summary 项目列表[分页+搜索]
// @Security ApiKeyAuth
// @Param limit query int false "条数"
// @Param skip query int false "偏移量"
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"result":[...],"total":1}}"
// @Router /v1/domains [get]
func (d *DomainController) List(c *gin.Context) {
	var listDto dto.GeneralListDto
	if d.BindAndValidate(c, &listDto) {
		data, total := domainService.List(listDto)
		resp(c, map[string]interface{}{
			"result": data,
			"total":  total,
		})
	}
}

// @Tags Domain
// @Summary 新增项目
// @Security ApiKeyAuth
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"id":1}}"
// @Router /v1/domains [post]
func (d *DomainController) Create(c *gin.Context) {
	var domainDto dto.DomainCreateDto
	if d.BindAndValidate(c, &domainDto) {
		created := domainService.Create(domainDto)
		if created.Id <= 0 {
			fail(c, ErrAddFail)
			return
		}
		resp(c, map[string]interface{}{
			"id": created.Id,
		})
	}
}

// @Tags Domain
// @Summary 删除项目
// @Security ApiKeyAuth
// @Param id path int true "项目id"
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"id":1}}"
// @Router /v1/domains/{id} [delete]
func (d *DomainController) Delete(c *gin.Context) {
	var domainDto dto.GeneralDelDto
	if d.BindAndValidate(c, &domainDto) {
		affected := domainService.Delete(domainDto)
		if affected <= 0 {
			fail(c, ErrDelFail)
			return
		}
		ok(c, "ok.DeletedDone")
	}
}

// @Tags Domain
// @Summary 编辑项目
// @Security ApiKeyAuth
// @Param id path int true "项目id"
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"id":1}}"
// @Router /v1/domains/{id} [put]
func (d *DomainController) Edit(c *gin.Context) {
	var domainDto dto.DomainEditDto
	if d.BindAndValidate(c, &domainDto) {
		affected := domainService.Update(domainDto)
		if affected > 0 {
		}
		ok(c, "ok.UpdateDone")
	}
}
