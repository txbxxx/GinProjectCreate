/**
 * @Author tanchang
 * @Description 用户服务接口
 * @Date 2024/12/20 14:24
 * @File:  UserInterface
 * @Software: vscode
 **/

package userSvc

type UserInterface interface {
	Login() interface{}
	Logout() interface{}
	Register() interface{}
}
