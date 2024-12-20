/**
 * @Author tanchang
 * @Description 鉴权模块
 * @Date 2024/7/11 20:09
 * @File:  Token
 * @Software: GoLand
 **/

package serializes

import "Go-WebCreate/model"

type UserSerialize struct {
	Identity  string `json:"identity"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	IsAdmin   int    `json:"is_admin"`
	Mail      string `json:"mail"`
	CreatedAt int64  `json:"created_at"`
}

// UserSerializeList 用户序列化列表
func UserSerializeList(users []model.User) []UserSerialize {
	var userSerializeList []UserSerialize
	for _, user := range users {
		userSerializeList = append(userSerializeList, UserSerialize{
			Identity:  user.Identity,
			Name:      user.Name,
			Phone:     user.Phone,
			Mail:      user.Mail,
			CreatedAt: user.CreatedAt.Unix(),
		})
	}
	return userSerializeList
}

// UserSerializeSingle 单个用户序列化
func UserSerializeSingle(user model.User) UserSerialize {
	return UserSerialize{
		Identity:  user.Identity,
		Name:      user.Name,
		Phone:     user.Phone,
		Mail:      user.Mail,
		CreatedAt: user.CreatedAt.Unix(),
	}
}
