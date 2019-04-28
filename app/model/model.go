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
	ACTION_ALL             = "(GET)|(POST)|(PUT)|(DELETE)|(PATCH)|(OPTIONS)|(HEAD)"
	ADMIN_NAME             = "admin" //超级管理员用户名
	ADMIN_NICK_NAME        = "超级管理员" //超级管理员显示名称
	ADMIN_DEFAULT_PASSWORD = "123"   //超级管理员默认密码
)

func InitModel() {
	defDB = g.DB()
	initUser()
	initCasbin()
}

// initUser description
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

// initCasbin description
//
// createTime:2019年04月23日 14:45:20
// author:hailaz
func initCasbin() {
	a := NewAdapter(defDB)
	Enforcer = casbin.NewEnforcer("./config/rbac.conf", a)
	Enforcer.LoadPolicy()
	//Enforcer.DeletePermissionsForUser("group_admin")
	Enforcer.AddPolicy("group_admin", "*", ACTION_ALL)
	Enforcer.AddPolicy("*", "/v1/user/loginkey", ACTION_GET)
	Enforcer.AddPolicy("*", "/v1/user/login", ACTION_POST)
	Enforcer.AddPolicy("*", "/v1/user/logout", ACTION_POST)
	Enforcer.AddGroupingPolicy(ADMIN_NAME, "group_admin")

}

// GetDB description
//
// createTime:2019年04月23日 11:53:21
// author:hailaz
func GetDB() gdb.DB {
	return defDB
}
