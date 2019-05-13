package model

import (
	"github.com/casbin/casbin"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var defDB gdb.DB
var Enforcer *casbin.Enforcer

const (
	ACTION_GET             = "(GET)"
	ACTION_POST            = "(POST)"
	ACTION_PUT             = "(PUT)"
	ACTION_DELETE          = "(DELETE)"
	ACTION_ALL             = "(GET)|(POST)|(PUT)|(DELETE)|(PATCH)|(OPTIONS)|(HEAD)"
	ADMIN_NAME             = "admin" //超级管理员用户名
	ADMIN_NICK_NAME        = "超级管理员" //超级管理员显示名称
	ADMIN_DEFAULT_PASSWORD = "123"   //超级管理员默认密码
)

// InitModel 初始化数据模型
//
// createTime:2019年05月13日 09:47:08
// author:hailaz
func InitModel() {
	defDB = g.DB()
	initUser()
	initCasbin()
}

// initUser 初始化用户
//
// createTime:2019年04月23日 14:57:23
// author:hailaz
func initUser() {
	u, err := GetUserByName(ADMIN_NAME)
	if err != nil || u.Id != 0 {
		return
	}
	admin := User{
		UserName: ADMIN_NAME,
		NickName: ADMIN_NICK_NAME,
		Password: ADMIN_DEFAULT_PASSWORD,
	}
	admin.Insert()
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
