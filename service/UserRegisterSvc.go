/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 21:40
 * @File:  UserSvc
 * @Software: GoLand
 **/

package service

type UserRegisterService struct {
	Identity string `form:"identity" json:"identity" binding:"required" `
	Name     string `form:"name" json:"name" binding:"required,min=5,max=10" `
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
	Phone    string `form:"phone" json:"phone" binding:"required,len=11"`
	Mail     string `form:"mail" json:"mail" binding:"required,email"`
	IsAdmin  int    `form:"is_admin" json:"is_admin" binding:"required,oneof=1 2"`
}
