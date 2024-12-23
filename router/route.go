/**
 * @Author tanchang
 * @Description 路由
 * @Date 2024/7/11 15:28
 * @File:  route
 * @Software: GoLand
 **/

package router

import (
	"Go-WebCreate/control"
	"Go-WebCreate/middleware/cors"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	httpServer := gin.Default()
	//跨域
	httpServer.Use(cors.Cors(),gin.Recovery(),gin.Logger())

	user := httpServer.Group("/user")
	{
		user.POST("/login", control.Login)
		user.POST("/register", control.Register)
	}

	return httpServer
}
