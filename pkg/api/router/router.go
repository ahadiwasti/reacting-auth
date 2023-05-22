package router

import (
	_ "./docs"
	"./pkg/api/controllers"
	"./pkg/api/middleware"
	"github.com/appleboy/gin-jwt/v2"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/swaggo/gin-swagger"
)

var jwtAuth *jwt.GinJWTMiddleware
var jwtAuths *jwt.GinJWTMiddleware

func SetUp(e *gin.Engine, cors bool) {
	e.Use(
		gin.Recovery(),
	)
	if cors {
		e.Use(middleware.Cors())
	}
	e.Use(middleware.SetLangVer())
	e.GET("/healthcheck", controllers.Healthy)
	e.Static("/file", "./data/images/") // 添加图片资源路径
	//version fragment
	v1 := e.Group("/admin/v1")
	//install
	InstallController := controllers.InstallController{}
	v1.POST("/install", InstallController.Install)

	if viper.GetBool("project.merge") {
		e.LoadHTMLGlob("./pkg/webui/dist/*.html") // 添加入口index.html
		//e.LoadHTMLFiles("./pkg/webui/dist/static/*/*")   // 添加资源路径
		e.Static("/admin/static", "./pkg/webui/dist/static")   // 添加资源路径
		e.StaticFile("/admin/", "./pkg/webui/dist/index.html") //前端接口
	}
}
