/**
 * @Author tanchang
 * @Description 错误返回
 * @Date 2024/12/22 16:02
 * @File:  Resp
 * @Software: vscode
**/

package resp

import "time"


type Resp struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
	Time	int64  `json:"time"`
	Data	interface{} `json:"data,omitempty"`
    Token   string `json:"token,omitempty"`
}

func NewErrorResp(code int, message string) *Resp {
    return &Resp{
        Code:    code,
        Message: message,
		Time:   time.Now().Unix(),
    }
}

func NewSuccessResp(code int, message string,data interface{},token string) *Resp {
    return &Resp{
        Code:    code,
        Token :  token,
        Message: message,
		Time:   time.Now().Unix(),
		Data:   data,
    }
}

func NewLoginSuccessResp(code int, message string,data interface{},token string) *Resp {
    return &Resp{
        Code:    code,
        Message: message,
        Time:   time.Now().Unix(),
        Data:   data,
        Token:  token,
    }
}