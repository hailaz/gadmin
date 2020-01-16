package model

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/database/gdb"
)

type MenuMetaOut struct {
	Id       int64  `json:"-"`       //
	MenuName string `json:"-"`       //菜单关键名
	Title    string `json:"title"`   //名称
	Icon     string `json:"icon"`    //图标
	Nocache  bool   `json:"noCache"` //是否缓存
}

type MenuOut struct {
	Id          int64       `json:"-"`           //
	MenuPath    string      `json:"path"`        //菜单路径
	Component   string      `json:"component"`   //页面模块
	Redirect    string      `json:"redirect"`    //重定向地址
	Name        string      `json:"name"`        //唯一关键名
	Hidden      bool        `json:"hidden"`      //是否隐藏
	Alwaysshow  bool        `json:"alwaysShow"`  //是否常显示
	Sort        int         `json:"sort"`        //排序
	ParentName  string      `json:"parent_name"` //父菜级关键名
	AutoCreate  bool        `json:"auto_create"` //是否自动生成
	MenuMetaOut MenuMetaOut `json:"meta"`
	Children    []MenuOut   `json:"children"`
}

// IsStringInSlice description
//
// createTime:2019年05月21日 15:50:15
// author:hailaz
func IsStringInSlice(str string, strs []string) bool {
	for _, item := range strs {
		if item == str {
			return true
		}
	}
	return false
}

// GetMenuByRoleName description
//
// createTime:2019年05月16日 17:19:53
// author:hailaz
func GetMenuByRoleName(roles []string) []MenuOut {
	menus := make([]MenuOut, 0)
	if IsStringInSlice(ADMIN_NAME, roles) {
		r, _ := defDB.Table("menu").All()
		r.ToStructs(&menus)
	} else {
		roleSlice := make(g.Slice, 0)
		for _, item := range roles {
			roleSlice = append(roleSlice, item)
		}
		r, _ := defDB.Table("role_menu rm,menu m").Where("rm.menu_name=m.name AND rm.role_key IN (?)", roleSlice).Fields("m.*").All()
		r.ToStructs(&menus)
	}

	for index, item := range menus {
		meta := MenuMetaOut{}
		r, _ := defDB.Table("menu_meta").Where("menu_name=?", item.Name).One()
		r.ToStruct(&meta)
		menus[index].MenuMetaOut = meta
	}
	menuRoot := make([]MenuOut, 0)
	childs := make([]*MenuOut, 0)
	for index, item := range menus { //分类菜单，一级菜单与非一级菜单
		if item.ParentName == "" {
			menuRoot = append(menuRoot, item)
		} else {
			childs = append(childs, &menus[index])
		}
	}
	for index, _ := range menuRoot {
		FindChildren(&menuRoot[index], childs)
	}

	return menuRoot
}

// FindChildren 找子菜单
//
// createTime:2019年05月17日 09:15:52
// author:hailaz
func FindChildren(mo *MenuOut, list []*MenuOut) {
	for _, item := range list {
		if item.ParentName == mo.Name {
			mo.Children = append(mo.Children, *item)
		}
	}
	for index := 0; index < len(mo.Children); index++ {
		FindChildren(&mo.Children[index], list)
	}
}

// GetMenuByName 根据名称获取菜单
//
// createTime:2019年04月23日 17:14:22
// author:hailaz
func GetMenuByName(name string) (*Menu, error) {
	m := Menu{}
	err := defDB.Table("menu").Where("name", name).Struct(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// IsMenuExist 菜单是否存在
//
// createTime:2019年05月17日 11:04:45
// author:hailaz
func IsMenuExist(name string) bool {
	m := Menu{}
	defDB.Table("menu").Where("name", name).Struct(&m)
	if m.Id > 0 {
		return true
	}
	return false
}

// InsertMenuWithMeta 插入菜单
//
// createTime:2019年05月17日 11:12:13
// author:hailaz
func InsertMenuWithMeta(list gdb.List) {
	for _, item := range list {
		if !IsMenuExist(item["name"].(string)) {
			mate := item["meta"].(gdb.Map)
			mate["menu_name"] = item["name"].(string)
			delete(item, "meta")
			defDB.Insert("menu", item)
			defDB.Insert("menu_meta", mate)
		}
	}
}

// GetMenuList 获取菜单列表
//
// createTime:2019年05月17日 16:17:33
// author:hailaz
func GetMenuList(page, limit int) ([]MenuOut, int) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	total, _ := defDB.Table("menu").Count()
	if total == 0 {
		return nil, 0
	}

	menuList := make([]MenuOut, 0)
	if total < page*limit {
		if total < limit {
			page = 1
		}
	}
	r, err := defDB.Table("menu").Limit((page-1)*limit, (page-1)*limit+limit).Select()
	if err != nil {
		return nil, 0
	}
	r.ToStructs(&menuList)
	for index, item := range menuList {
		meta := MenuMetaOut{}
		r, _ := defDB.Table("menu_meta").Where("menu_name=?", item.Name).One()
		r.ToStruct(&meta)
		menuList[index].MenuMetaOut = meta
	}
	return menuList, total
}

// UpdateMenuByName description
//
// createTime:2019年05月17日 17:54:40
// author:hailaz
func UpdateMenuByName(name string, dataMap gdb.Map) error {
	mate := dataMap["meta"].(gdb.Map)
	delete(dataMap, "meta")
	_, err := defDB.Update("menu", dataMap, "name=?", name)
	if err != nil {
		return err
	}
	_, err = defDB.Update("menu_meta", mate, "menu_name=?", name)
	if err != nil {
		return err
	}
	return nil
}
