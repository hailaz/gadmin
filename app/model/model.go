package model

import (
	"github.com/casbin/casbin"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	// _ "github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
)

var defDB gdb.DB
var Enforcer *casbin.Enforcer

const (
	ACTION_GET             = "(GET)"
	ACTION_POST            = "(POST)"
	ACTION_PUT             = "(PUT)"
	ACTION_DELETE          = "(DELETE)"
	ACTION_ALL             = "(GET)|(POST)|(PUT)|(DELETE)|(PATCH)|(OPTIONS)|(HEAD)"
	ADMIN_NAME             = "admin"  //超级管理员用户名
	ADMIN_NICK_NAME        = "超级管理员"  //超级管理员显示名称
	ADMIN_DEFAULT_PASSWORD = "123456" //超级管理员默认密码
)

// InitModel 初始化数据模型
//
// createTime:2019年05月13日 09:47:08
// author:hailaz
func InitModel() {
	defDB = g.DB()
	//defDB.SetDebug(true)
	initUser()
	initCasbin()
	initMenu()
}

// initUser 初始化用户
//
// createTime:2019年04月23日 14:57:23
// author:hailaz
func initUser() {
	u, err := GetUserByName(ADMIN_NAME)
	if err == nil && u != nil && u.Id != 0 {
		return
	}
	admin := User{
		UserName: ADMIN_NAME,
		NickName: ADMIN_NICK_NAME,
		Password: ADMIN_DEFAULT_PASSWORD,
	}
	admin.Insert()
}

// initMenu 初始化菜单数据
//
// createTime:2019年05月16日 15:39:54
// author:hailaz
func initMenu() {
	InsertMenuWithMeta(gdb.List{
		{
			"name":        "user",
			"menu_path":   "/user",
			"component":   "layout",
			"redirect":    "/user/list",
			"sort":        "0",
			"parent_name": "",
			"auto_create": true,
			"meta": gdb.Map{
				"title":   "user",
				"icon":    "user",
				"noCache": 0},
		},
		{
			"name":        "userList",
			"menu_path":   "list",
			"component":   "user/user",
			"sort":        "0",
			"parent_name": "user",
			"auto_create": true,
			"meta": gdb.Map{
				"title":   "userList",
				"icon":    "",
				"noCache": 0},
		},
		{
			"name":        "roleList",
			"menu_path":   "/role/list",
			"component":   "user/role",
			"sort":        "1",
			"parent_name": "user",
			"auto_create": true,
			"meta": gdb.Map{
				"title":   "roleList",
				"icon":    "",
				"noCache": 0},
		},
		{
			"name":        "policyList",
			"menu_path":   "/policy/list",
			"component":   "user/policy",
			"sort":        "2",
			"parent_name": "user",
			"auto_create": true,
			"meta": gdb.Map{
				"title":   "policyList",
				"icon":    "",
				"noCache": 0},
		},
		{
			"name":        "menuList",
			"menu_path":   "/menu/list",
			"component":   "user/menu",
			"sort":        "3",
			"parent_name": "user",
			"auto_create": true,
			"meta": gdb.Map{
				"title":   "menuList",
				"icon":    "",
				"noCache": 0},
		},
	})

}

// initCasbin 初始化Casbin
//
// createTime:2019年04月23日 14:45:20
// author:hailaz
func initCasbin() {
	a := NewAdapter(defDB)
	Enforcer = casbin.NewEnforcer("./config/rbac.conf", a)
	Enforcer.LoadPolicy()
	//Enforcer.DeletePermissionsForUser("group_admin")
	Enforcer.AddPolicy(ADMIN_NAME, "*", ACTION_ALL)
	//Enforcer.AddGroupingPolicy("system", "user")

}

// GetDB 获取默认DB
//
// createTime:2019年04月23日 11:53:21
// author:hailaz
func GetDB() gdb.DB {
	return defDB
}
