package router

import (
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/hailaz/gadmin/app/api"
	"github.com/hailaz/gadmin/app/model"
)

func InitRouter(s *ghttp.Server) {

	loginCtrl := new(api.LoginController)
	s.BindControllerMethod("GET:/loginkey", loginCtrl, "GetLoginCryptoKey")
	s.BindControllerMethod("POST:/login", loginCtrl, "Login")
	v1 := s.Group("/v1")
	userCtrl := new(api.UserController)
	//权限验证
	v1.ALL("*", api.NewAuthorizer(model.Enforcer), ghttp.HOOK_BEFORE_SERVE)
	v1.Bind([]ghttp.GroupItem{
		{"GET", "/show", userCtrl, "Show"},
	})

	v1.Bind([]ghttp.GroupItem{
		{"GET", "/login", userCtrl, "Login"},
	})

}
