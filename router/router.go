package router

import (
	"fmt"

	"github.com/gogf/gf/encoding/gparser"
	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"github.com/hailaz/gadmin/app/api"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/library/common"
)

var routerMap = make(map[string]model.RolePolicy)

func showURL(r *ghttp.Request) {
	glog.Debug("请求路径：", r.Method, r.Request.RequestURI)
	//r.Response.CORSDefault()
}

// InitRouter 初始化路由
//
// createTime:2019年05月13日 09:32:58
// author:hailaz
func InitRouter(s *ghttp.Server) {
	initApiDocRouter(s)

	s.BindHookHandler("/*", ghttp.HOOK_BEFORE_SERVE, showURL)
	InitV1(s)

	model.ReSetPolicy("system", routerMap)
}

// authHook 鉴权钩子
//
// createTime:2019年05月13日 09:33:58
// author:hailaz
// authHook is the HOOK function implements JWT logistics.
func authHook(r *ghttp.Request) {
	switch r.Request.RequestURI { //登录相关免鉴权
	case "/v1/loginkey":
		return
	case "/v1/login":
		return
	default:
		r.Response.CORSDefault()                //开启跨域
		api.GfJWTMiddleware.MiddlewareFunc()(r) //鉴权中间件
	}
}

// InitV1 初始化V1
//
// createTime:2019年04月25日 09:24:06
// author:hailaz
func InitV1(s *ghttp.Server) {
	//v1 := s.Group("/v1")
	//权限验证
	s.Group("/v1").ALL("/*any", authHook, ghttp.HOOK_BEFORE_SERVE)
	userCtrl := new(api.UserController)
	roleCtrl := new(api.RoleController)
	policyCtrl := new(api.PolicyController)
	menuCtrl := new(api.MenuController)

	// user
	BindGroup(s, "/v1", []ghttp.GroupItem{
		//登录
		{"false", "GET", "/loginkey", api.GetLoginCryptoKey},                   //获取登录加密公钥
		{"false", "POST", "/login", api.GfJWTMiddleware.LoginHandler},          //登录
		{"false", "GET", "/refresh_token", api.GfJWTMiddleware.RefreshHandler}, //获取登录加密公钥
		{"false", "POST", "/logout", api.Logout},                               //登出
		//menu
		{"", "REST", "/menu", menuCtrl},
		// 用户
		{"false", "GET", "/user/info", userCtrl, "Info"},
		{"false", "GET", "/user/menu", userCtrl, "Menu"},
		{"", "REST", "/user", userCtrl},
		// 角色
		{"", "REST", "/role", roleCtrl},
		{"", "PUT", "/role/byuser", roleCtrl, "SetRoleByUserName"},
		{"", "PUT", "/role/menu", roleCtrl, "SetRoleMenus"},
		// 权限
		{"", "REST", "/policy", policyCtrl},
		{"", "GET", "/policy/byrole", policyCtrl, "GetPolicyByRole"},
		{"", "PUT", "/policy/byrole", policyCtrl, "SetPolicyByRole"},
	})
}

// BindGroup 绑定分组路由
//
// createTime:2019年04月29日 16:45:55
// author:hailaz
func BindGroup(s *ghttp.Server, path string, items []ghttp.GroupItem) {
	for index, item := range items {
		if gconv.String(item[0]) == "false" { //不走权限的api
			addPolicy("*", path+gconv.String(item[2]), common.GetAction(gconv.String(item[1])))
		} else { //走权限的api
			if gconv.String(item[1]) == "REST" { //rest api
				addPolicy("system", path+gconv.String(item[2]), model.ACTION_GET)
				addPolicy("system", path+gconv.String(item[2]), model.ACTION_POST)
				addPolicy("system", path+gconv.String(item[2]), model.ACTION_PUT)
				addPolicy("system", path+gconv.String(item[2]), model.ACTION_DELETE)
			} else {
				addPolicy("system", path+gconv.String(item[2]), common.GetAction(gconv.String(item[1])))
			}
		}
		// 裁切掉权限控制
		items[index] = items[index][1:]
	}

	g := s.Group(path)
	g.Bind(items)

}

// addPolicy 记录需要系统路由
//
// createTime:2019年04月29日 17:18:25
// author:hailaz
func addPolicy(role, path, atc string) {
	routerMap[fmt.Sprintf("%v %v %v", role, path, atc)] = model.RolePolicy{Role: role, Path: path, Atc: atc}
}

// initApiDocRouter 初始化api文档路由
//
// createTime:2019年05月14日 09:05:26
// author:hailaz
func initApiDocRouter(s *ghttp.Server) {
	s.BindHandler("/swagger.{mime}", func(r *ghttp.Request) {
		r.Response.CORSDefault()
		r.Response.Header().Set("Access-Control-Allow-Origin", "*")
		//r.Response.Writeln(r.Get("mime"))
		p, err := gparser.Load("docfile/swagger.yaml")
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}

		switch r.Get("mime") {
		case "json":
			j, _ := p.ToJson()
			r.Response.WriteJson(j)
		default:
			y, _ := p.ToYaml()
			r.Response.Write(y)
		}

	})
	s.BindHandler("/swagger", func(r *ghttp.Request) {
		r.Response.RedirectTo("https://petstore.swagger.io/?url=http://localhost:" + g.Config().GetString("port", "8080") + "/swagger.yaml")
	})
}
