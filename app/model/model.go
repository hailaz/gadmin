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
	ACTION_ALL = "(GET)|(POST)|(PUT)|(DELETE)|(PATCH)|(OPTIONS)|(HEAD)"
)

func init() {
	// 设置默认配置文件，默认的 config.toml 将会被覆盖
	g.Config().SetFileName("database.conf")
	defDB = g.DB()
	initUser()
	initCasbin()
}

// initUser description
//
// creatTime:2019年04月23日 14:57:23
// author:hailaz
func initUser() {
	u := User{}
	err := defDB.Table("user").Where("user_name", "admin").Struct(&u)
	if err != nil {
		return
	}
	if u.Id != 0 {
		return
	}
	admin := User{
		UserName: "admin",
		NickName: "超级管理员",
		Password: "123",
	}
	admin.Insert()
}

// initCasbin description
//
// creatTime:2019年04月23日 14:45:20
// author:hailaz
func initCasbin() {
	a := NewAdapter(defDB)
	Enforcer = casbin.NewEnforcer("./config/rbac.conf", a)
	Enforcer.LoadPolicy()
	//Enforcer.DeletePermissionsForUser("group_admin")
	Enforcer.AddPolicy("group_admin", "*", ACTION_ALL)
	Enforcer.AddGroupingPolicy("admin", "group_admin")

}

// GetDB description
//
// creatTime:2019年04月23日 11:53:21
// author:hailaz
func GetDB() gdb.DB {
	return defDB
}
