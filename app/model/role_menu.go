package model

import (
	"github.com/gogf/gf/database/gdb"
)

// DeleteRoleMenus description
//
// createTime:2019年05月21日 17:52:06
// author:hailaz
func DeleteRoleMenus(role string) {
	defDB.Delete("role_menu", "role_key=?", role)
}

// SetRoleMenus description
//
// createTime:2019年05月21日 17:54:38
// author:hailaz
func SetRoleMenus(role string, menus []string) {
	DeleteRoleMenus(role)
	ms := make(gdb.List, 0)
	for _, item := range menus {
		ms = append(ms, gdb.Map{"role_key": role, "menu_name": item})
	}

	defDB.Table("role_menu").Data(ms).Insert()
}
