/**
 * @Author tanchang
 * @Description UserService
 * @Date  2024/12/20 14:21
 * @File:  UserService
 * @Software: Vscode
 **/

package userSvc

import (
	"Go-WebCreate/model"
	serializes "Go-WebCreate/serialized"
	DB "Go-WebCreate/utils/DB/mariadb"
	token "Go-WebCreate/utils/token"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type  UserService struct {

}


func (t *UserService) Login(u,pwd string) interface{} {
	//获取用户信息
	var (
		data model.User
		sql  = DB.GetSqlConn()
	)
	errData := sql.Model(&model.User{}).
				Where("name = ? and password = ?", u, token.GetMd5(pwd)).
				Find(&data).Error

	if errData != nil {
		logrus.Error("获取用户信息失败: " + errData.Error())
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


	//生成token
	token, errToken := token.GenerateToken(data.Identity, data.Name, data.IsAdmin)
	if errToken != nil {
		logrus.Error("生成Token失败: " + errToken.Error())
		return gin.H{
			"code": "-1",
			"msg":  "生成Token失败",
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

func (u *UserService) Logout() interface{} {
	panic("TODO: Implement")
}


func (t *UserService) Register(u,pwd,phone,mail string) interface{} {
	//查找是否存在用户
	var (
		userData model.User
		sql      = DB.GetSqlConn()
	)
	errSearchUser := sql.Model(&model.User{}).Where("name = ? ", u).Find(&userData).Error
	if errSearchUser != nil {
		logrus.Error("查找用户失败" + errSearchUser.Error())
		return gin.H{
			"code": -1,
			"msg":  "服务器内部错误，请联系管理员或者提交issue",
		}

	}
	if userData != (model.User{}) {
		return gin.H{
			"code": -1,
			"msg":  "用户已存在",
		}
	}

	//创建用户对象
	user := &model.User{
		Name:     u,
		Password: token.GetMd5(pwd),
		Identity: token.GenerateUUID(),
		Phone:    phone,
		Mail:     mail,
	}
	//插入数据库
	errInsert := sql.Model(&model.User{}).Create(user).Error
	if errInsert != nil {
		return gin.H{
			"code": -1,
			"msg":  "服务器内部错误，请联系管理员或者提交issue" + errInsert.Error(),
		}
	}

	//生成token
	token, errToken := token.GenerateToken(user.Identity, user.Name, user.IsAdmin)
	if errToken != nil {
		return gin.H{
			"code": -1,
			"msg":  "生成Token失败" + errToken.Error(),
		}
	}

	return gin.H{
		"code": 200,
		"msg":  "注册成功！！",
		"data": gin.H{
			"data":  serializes.UserSerializeSingle(*user),
			"token": token,
		},
	}
}



func NewUserSvc() *UserService {
	return new(UserService)
}