/**
 * @Author tanchang
 * @Description main
 * @Date 2024/7/11 14:52
 * @File:  main
 * @Software: GoLand
 **/

package main

import (
	"os"

	"Go-WebCreate/conf"
	"Go-WebCreate/router"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化配置
	conf.Init()
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := router.Router()
	//启动http服务
	err := r.Run(os.Getenv("GIN_PORT"))
	if err != nil {
		return
	}
}
