/**
 * @Author tanchang
 * @Description user contorller
 * @Date 2024/7/11 21:12
 * @File:  User
 * @Software: GoLand
 **/

package control

import (
	"Go-WebCreate/serialized/resp"
	"Go-WebCreate/serialized/status"
	"Go-WebCreate/service/userSvc"

	"github.com/gin-gonic/gin"
)


var svc = userSvc.NewUserSvc()

type UserRegisterReq struct {
	Name     string `form:"name" json:"name" binding:"required,min=5,max=10" `
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
	Phone    string `form:"phone" json:"phone" binding:"required,len=11"`
	Mail     string `form:"mail" json:"mail" binding:"required,email"`
}


type UserLoginrReq struct {
	Name     string `form:"name" json:"name" binding:"required,min=5,max=10" `
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}





// Login 登录
func Login(c *gin.Context) { 
	userReq := &UserLoginrReq{}
	err := c.ShouldBindJSON(userReq)
	if err == nil { 
		login := svc.Login(userReq.Name, userReq.Password)
		c.JSON(200, login)
	} else {
		c.AbortWithStatusJSON(400,resp.NewErrorResp(status.UserLoginError, err.Error()))
	}
}

// Register 注册
func Register(c *gin.Context) {
	userReq := &UserRegisterReq{}
	err := c.ShouldBindJSON(userReq)
	if err == nil {
		register := svc.Register(userReq.Name, userReq.Password, userReq.Phone, userReq.Mail)
		c.JSON(200, register)
	} else {
		c.AbortWithStatusJSON(400, resp.NewErrorResp(status.UserRegisterError, err.Error()))
	}
}

