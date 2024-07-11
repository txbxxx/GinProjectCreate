/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 21:59
 * @File:  UserLoginSvc
 * @Software: GoLand
 **/

package service

import (
	"Go-WebCreate/model"
	"Go-WebCreate/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserLoginService struct {
	Name     string `form:"name" json:"name" binding:"required,min=5,max=10" `
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

func (This *UserLoginService) Login(name, password string) gin.H {
	//生成password的md5值
	password = utils.GetMd5(password)
	utils.DB.Model(&model.User{})
	//获取用户信息
	var data model.User
	var cunt int64
	errData := utils.DB.Model(&model.User{}).Where("name = ? and password = ?", name, password).Count(&cunt).Find(&data).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			}

		}
		return gin.H{
			"code": -1,
			"msg":  "服务器内部错误，请联系管理员或者提交issue" + errData.Error(),
		}
	}
	if cunt == 0 {
		return gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		}
	}

	//生成token
	token, errToken := utils.GenerateToken(data.Identity, data.Name, data.IsAdmin)
	if errToken != nil {
		return gin.H{
			"code": "-1",
			"msg":  "生成Token失败" + errToken.Error(),
		}
	}

	return gin.H{
		"code": 200,
		"msg":  "登录成功！！",
		"data": gin.H{
			"token":    token,
			"is_admin": data.IsAdmin,
		},
	}
}
