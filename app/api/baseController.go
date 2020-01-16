package api

import (
	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/frame/gmvc"
	"github.com/gogf/gf/net/ghttp"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/library/code"
)

type BaseController struct {
	gmvc.Controller
}

// "析构函数"控制器方法
func (c *BaseController) Shut() {

}

type BaseResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Response API返回
//
// createTime:2019年04月25日 11:32:47
// author:hailaz
func Response(r *ghttp.Request, rs BaseResult) {
	r.Response.WriteJson(rs)
	r.ExitAll()
}

// Success 返回成功
//
// createTime:2019年04月25日 11:41:44
// author:hailaz
func Success(r *ghttp.Request, data interface{}) {
	Response(r, BaseResult{Code: code.RESPONSE_SUCCESS, Message: "success", Data: data})
}

// Fail 返回失败
//
// createTime:2019年04月25日 11:43:34
// author:hailaz
func Fail(r *ghttp.Request, errCode int, msg ...string) {
	if len(msg) > 0 {
		Response(r, BaseResult{Code: errCode, Message: msg[0]})
	} else {
		Response(r, BaseResult{Code: errCode, Message: "fail"})
	}

}

// funcName 获取当前用户
//
// createTime:2019年05月13日 10:01:17
// author:hailaz
func (c *BaseController) GetUser() *model.User {
	claims := jwt.ExtractClaims(c.Request)
	user, _ := model.GetUserByName(claims["username"].(string))
	return user
}
