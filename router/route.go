/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 15:28
 * @File:  route
 * @Software: GoLand
 **/

package router

import (
	"Go-WebCreate/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	httpServer := gin.Default()
	//跨域
	httpServer.Use(middleware.Cors())

	return httpServer
}
