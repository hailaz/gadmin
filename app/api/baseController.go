package api

import (
	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/net/ghttp"
)

type BaseController struct {
	gmvc.Controller
}

// "析构函数"控制器方法
func (c *BaseController) Shut() {

}

type BaseResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response description
//
// creatTime:2019年04月25日 11:32:47
// author:hailaz
func Response(rsp *ghttp.Response, rs BaseResult) {
	rsp.WriteJson(rs)
}

// Success description
//
// creatTime:2019年04月25日 11:41:44
// author:hailaz
func Success(rsp *ghttp.Response, data interface{}) {
	Response(rsp, BaseResult{Code: 0, Msg: "success", Data: data})
}

// Fail description
//
// creatTime:2019年04月25日 11:43:34
// author:hailaz
func Fail(rsp *ghttp.Response, errCode int, msg ...string) {
	if len(msg) > 0 {
		Response(rsp, BaseResult{Code: errCode, Msg: msg[0]})
	} else {
		Response(rsp, BaseResult{Code: errCode, Msg: "fail"})
	}

}
