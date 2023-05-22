package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/ahadiwasti/reacting-auth/pkg/api/dto"
	"github.com/ahadiwasti/reacting-auth/pkg/api/log"
	"github.com/ahadiwasti/reacting-auth/pkg/api/service"
	"github.com/gin-gonic/gin"
)

var logService = service.LogService{}
var ignoreRoutes = map[string]bool{
	"/v1/account/idle":               false,
	"/admin/v1/account/idle":         true,
	"/v1/account/password":           true,
	"/v1/account/require-change-pwd": true,
}

func AccessLog(c *gin.Context) {
	//too much useless or some important logs
	if _, ok := ignoreRoutes[c.Request.URL.Path]; ok {
		c.Next()
		return
	}
	b, _ := json.Marshal(c.Request.URL.Query())
	body, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	uid := 0
	if c.Value("userId") != nil {
		uid = int(c.Value("userId").(float64))
	}
	orLogDto := dto.OperationLogDto{
		UserId:           uid,
		RequestUrl:       c.Request.URL.Path,
		OperationMethod:  c.Request.Method,
		Params:           "[GET] -> " + string(b) + " | [POST] -> " + string(body),
		Ip:               c.ClientIP(),
		IpLocation:       "", //TODO...待接入获取ip位置服务
		OperationResult:  "success",
		OperationSuccess: 1,
		OperationContent: "-",
	}
	err := logService.InsertOperationLog(orLogDto)
	if err != nil {
		log.Error(err.Error())
	}
	c.Next()
}
