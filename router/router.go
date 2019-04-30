package router

import (
	"fmt"

	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/util/gconv"
	"github.com/hailaz/gadmin/app/api"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/library/common"
)

type Policy struct {
	Role string
	Path string
	Atc  string
}

var routerMap = make(map[string]Policy)

func cors(r *ghttp.Request) {
	glog.Debug(r.Request.RequestURI)
	r.Response.CORSDefault()
}

func InitRouter(s *ghttp.Server) {

	s.BindHookHandler("/v1/*any", ghttp.HOOK_BEFORE_SERVE, cors)

	InitV1(s)

	//glog.Debug(model.Enforcer.GetRolesForUser("admin"))
	ReSetPolicy()
}

// InitV1 初始化V1
//
// createTime:2019年04月25日 09:24:06
// author:hailaz
func InitV1(s *ghttp.Server) {
	//v1 := s.Group("/v1")
	//权限验证
	s.Group("/v1").ALL("/*any", api.NewAuthorizer(model.Enforcer), ghttp.HOOK_BEFORE_SERVE)
	userCtrl := new(api.UserController)
	// user
	BindGroup(s, "/v1", []ghttp.GroupItem{
		{"GET", "/user/loginkey", userCtrl, "GetLoginCryptoKey"},
		{"GET", "/user/info", userCtrl, "Info"},
		{"POST", "/user/login", userCtrl, "Login"},
		{"POST", "/user/logout", userCtrl, "Logout"},

		{"GET", "/user/list", userCtrl, "List"},
		{"POST", "/user/add", userCtrl, "AddUser"},
	})
	// role
	BindGroup(s, "/v1", []ghttp.GroupItem{
		{"GET", "/show", userCtrl, "Show"},
		{"GET", "/add", userCtrl, "AddUser"},
	})
}

// ReSetPolicy description
//
// createTime:2019年04月29日 17:30:26
// author:hailaz
func ReSetPolicy() {
	old := model.Enforcer.GetPermissionsForUser("system")
	for _, item := range old {
		glog.Debug(item)
		full := fmt.Sprintf("%v %v %v", item[0], item[1], item[2])
		if _, ok := routerMap[full]; ok { //从待插入列表中删除已存在的路由
			delete(routerMap, full)
		} else { //删除不存在的旧路由
			model.Enforcer.DeletePermissionForUser(item[0], item[1], item[2])
		}
	}
	for _, item := range routerMap { //插入新路由
		model.Enforcer.AddPolicy(item.Role, item.Path, item.Atc)
	}
}

// BindGroup path string,description
//
// createTime:2019年04月29日 16:45:55
// author:hailaz
func BindGroup(s *ghttp.Server, path string, items []ghttp.GroupItem) {
	g := s.Group(path)
	g.Bind(items)
	for _, item := range items {
		addPolicy("system", path+gconv.String(item[1]), common.GetAction(gconv.String(item[0])))
	}

}

// addPolicy 记录需要系统路由
//
// createTime:2019年04月29日 17:18:25
// author:hailaz
func addPolicy(role, path, atc string) {
	routerMap[fmt.Sprintf("%v %v %v", role, path, atc)] = Policy{Role: role, Path: path, Atc: atc}
}
