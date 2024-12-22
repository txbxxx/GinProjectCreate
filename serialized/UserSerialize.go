/**
 * @Author tanchang
 * @Description 鉴权模块
 * @Date 2024/7/11 20:09
 * @File:  Token
 * @Software: GoLand
 **/

package serializes

import "Go-WebCreate/model"

type UserSerializeBase struct {
	Identity  string `json:"identity"`
	Name      string `json:"name"`
	Phone     string `json:"phone,omitempty"`
	IsAdmin   int    `json:"is_admin,omitempty"`
	Mail      string `json:"mail,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}

type UserLoginSerialize struct {
	IsAdmin int    `json:"is_admin"`
	Id      string `json:"id"`
	Token   string `json:"token"`
}

// UserSerializeList 用户序列化列表
func NewUserSerializeList(users []model.User) []UserSerializeBase {
	var userSerializeList []UserSerializeBase
	for _, user := range users {
		userSerializeList = append(userSerializeList, UserSerializeBase{
			Identity:  user.Identity,
			Name:      user.Name,
			Phone:     user.Phone,
			Mail:      user.Mail,
		})
	}
	return userSerializeList
}

// UserSerializeSingle 单个用户序列化
func NewUserSerializeSingle(user model.User) UserSerializeBase {
	return UserSerializeBase{
		Identity:  user.Identity,
		Name:      user.Name,
		Phone:     user.Phone,
		Mail:      user.Mail,
	}
}


func NewUserLoginSerialize (user model.User,token string) UserLoginSerialize {
	return UserLoginSerialize{
		IsAdmin: user.IsAdmin,
		Id:      user.Identity,
		Token:   token,
	}
}