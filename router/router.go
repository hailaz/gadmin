package router

import (
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/hailaz/gadmin/app/api"
	"github.com/hailaz/gadmin/app/model"
)

func InitRouter(s *ghttp.Server) {
	userCtrl := new(api.UserController)
	s.Group("/v1").Bind([]ghttp.GroupItem{
		{"ALL", "*", api.NewAuthorizer(model.Enforcer), ghttp.HOOK_BEFORE_SERVE}, //权限验证
		{"GET", "/show", userCtrl, "Show"},
	})

}
