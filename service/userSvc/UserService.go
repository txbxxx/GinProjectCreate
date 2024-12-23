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
	"Go-WebCreate/serialized/resp"
	"Go-WebCreate/serialized/status"
	DB "Go-WebCreate/utils/DB/mariadb"
	token "Go-WebCreate/utils/token"
	"Go-WebCreate/utils/log"
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type  UserService struct {

}


func (t *UserService) Login(u,pwd string) *resp.Resp {
	//获取用户信息
	var (
		data model.User
		sql  = DB.GetSqlConn()
	)
	errData := sql.Model(&model.User{}).
				Where("name = ? and password = ?", u, token.GetMd5(pwd)).
				Find(&data).Error

	if errData != nil && errors.Is(errData, gorm.ErrRecordNotFound) {
		logrus.Error("获取用户信息失败: " + errData.Error())
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return  resp.NewErrorResp(status.UserLoginError, "用户不存在")
		}
		return  resp.NewErrorResp(status.DBQueryError, "服务器内部错误，请联系管理员或者提交issue")
	}


	//生成token
	token, errToken := token.GenerateToken(data.Identity, data.Name, data.IsAdmin)
	if errToken != nil {
		log.Errorln("生成Token失败: " + errToken.Error())
		return resp.NewErrorResp(status.TokenGenerateError, "生成Token失败")
	}

	return  resp.NewLoginSuccessResp(status.Success, "登录成功", serializes.NewUserSerializeSingle(data), token)
}


func (u *UserService) Logout() interface{} {
	panic("TODO: Implement")
}



func (t *UserService) Register(u,pwd,phone,mail string) *resp.Resp {
	//查找是否存在用户
	var (
		userData model.User
		sql      = DB.GetSqlConn()
	)
	errSearchUser := sql.Model(&model.User{}).Where("name = ? ", u).Find(&userData).Error
	if errSearchUser != nil && !errors.Is(errSearchUser, gorm.ErrRecordNotFound) {
        logrus.Error("查找用户失败: " + errSearchUser.Error())
        return resp.NewErrorResp(status.DBQueryError, "服务器内部错误，请联系管理员或者提交issue")
    }
    if errSearchUser == nil {
        return resp.NewErrorResp(status.UserAlreadyExists, "用户已存在")
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
		logrus.Errorf("插入用户失败: %s", errInsert.Error())
		return resp.NewErrorResp(status.DBInsertError, "服务器内部错误，请联系管理员或者提交issue")
	}

	//生成token
	tokenAuth, errToken := token.GenerateToken(user.Identity, user.Name, user.IsAdmin)
	if errToken != nil {
		return resp.NewErrorResp(status.TokenGenerateError, "生成Token失败")
	}
	return resp.NewSuccessResp(status.UserSuccess, "注册成功", serializes.NewUserSerializeSingle(*user),tokenAuth)
}



func NewUserSvc() *UserService {
	return new(UserService)
}