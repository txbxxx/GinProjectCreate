/**
 * @Author tanchang
 * @Description 返回状态码
 * @Date 2024/12/20 17:09
 * @File:  respStatu
 * @Software: vscode
**/


package status


// RespStatu 返回状态码

const (
	// 成功
	Success = iota + 1
)

const (
    // 用户相关错误码
    UserSuccess = iota + 2000
    UserNotFound 
    UserPasswordError 
    UserAlreadyExists 
    UserRegisterError
    UserLoginError 
    TokenGenerateError 
)


const (
    // 数据库相关错误码
    DBSuccess =  iota + 3000
    DBConnectionError 
    DBQueryError
    DBInsertError 
    DBUpdateError 
    DBDeleteError 
)