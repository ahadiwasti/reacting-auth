package middleware

import (
	"./pkg/api/model"
	"github.com/appleboy/gin-jwt/v2"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// JwtPrepare : parse jwt and set login session info
func JwtPrepare(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get("id")
	c.Set("userId", claims["id"])
	c.Set("userName", user.(model.UserClaims).Name)
	c.Next()
}
