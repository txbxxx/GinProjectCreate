/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 21:12
 * @File:  User
 * @Software: GoLand
 **/

package control

import (
	"Go-WebCreate/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var userLoginService service.UserLoginService
	err := c.ShouldBind(&userLoginService)
	if err != nil {
		login := userLoginService.Login(userLoginService.Name, userLoginService.Password)
		c.JSON(200, login)
	} else {
		c.JSON(200, gin.H{"err": err})
	}

}
