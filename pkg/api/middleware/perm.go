package middleware

import (
	"fmt"
	"net/http"
	"reacting-auth/pkg/api/utils/log"
	"strings"

	"github.com/gin-gonic/gin"
)

// Ignored permissions
var ignoredPerms = map[string]bool{}

// PermCheck - check permission automatically
func PermCheck(c *gin.Context) {
	route := strings.Split(c.Request.URL.RequestURI(), "?")[0]
	for _, p := range c.Params {
		route = strings.Replace(route, "/"+p.Value, "/:"+p.Key, 1)
	}
	route = strings.ToLower(c.Request.Method) + "@" + route
	uid := fmt.Sprintf("%#v", c.Value("userId"))
	if _, ok := ignoredPerms[route]; ok {
		c.Next()
		return
	}
	check, _ := accountService.CheckPermission(uid, "root", route)
	if !check {
		log.Warn(fmt.Sprintf("No permission for %s", route))
		c.JSON(http.StatusOK, gin.H{
			"code": 403,
			"msg":  "err.Err403",
		})
		c.Abort()
		return
	} else {
		log.Info(fmt.Sprintf("Pass permission check for %s", route))
	}
	c.Next()
}
