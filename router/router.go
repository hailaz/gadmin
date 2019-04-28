package router

import (
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/hailaz/gadmin/app/api"
	"github.com/hailaz/gadmin/app/model"
)

func cors(r *ghttp.Request) {
	glog.Debug(r.Request.RequestURI)
	r.Response.CORSDefault()
}

func InitRouter(s *ghttp.Server) {
	/*s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.RedirectTo("/admin/index.html")
	})*/

	s.BindHookHandler("/*any", ghttp.HOOK_BEFORE_SERVE, cors)

	InitV1(s)
}

// InitV1 初始化V1
//
// createTime:2019年04月25日 09:24:06
// author:hailaz
func InitV1(s *ghttp.Server) {
	v1 := s.Group("/v1")
	userCtrl := new(api.UserController)
	//权限验证
	v1.ALL("/*any", api.NewAuthorizer(model.Enforcer), ghttp.HOOK_BEFORE_SERVE)
	v1.Bind([]ghttp.GroupItem{
		{"GET", "/user/loginkey", userCtrl, "GetLoginCryptoKey"},
		{"GET", "/user/info", userCtrl, "Info"},
		{"POST", "/user/login", userCtrl, "Login"},
		{"POST", "/user/logout", userCtrl, "Logout"},
	})
	v1.Bind([]ghttp.GroupItem{
		{"GET", "/show", userCtrl, "Show"},
		{"GET", "/add", userCtrl, "AddUser"},
	})
}
