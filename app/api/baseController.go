package api

import (
	"github.com/gogf/gf/g/frame/gmvc"
)

type BaseController struct {
	gmvc.Controller
}

// "析构函数"控制器方法
func (c *BaseController) Shut() {

}
