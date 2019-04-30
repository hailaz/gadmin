package api

import (
	"github.com/gogf/gf/g/frame/gmvc"
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

// Response description
//
// createTime:2019年04月25日 11:32:47
// author:hailaz
func Response(c gmvc.Controller, rs BaseResult) {
	c.Response.WriteJson(rs)
	c.Exit()
}

// Success description
//
// createTime:2019年04月25日 11:41:44
// author:hailaz
func Success(c gmvc.Controller, data interface{}) {
	Response(c, BaseResult{Code: code.RESPONSE_SUCCESS, Message: "success", Data: data})
}

// Fail description
//
// createTime:2019年04月25日 11:43:34
// author:hailaz
func Fail(c gmvc.Controller, errCode int, msg ...string) {
	if len(msg) > 0 {
		Response(c, BaseResult{Code: errCode, Message: msg[0]})
	} else {
		Response(c, BaseResult{Code: errCode, Message: "fail"})
	}

}
