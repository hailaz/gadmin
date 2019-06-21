package model

import (
	"errors"
)

// Menu 表名：menu
// 由数据库自动生成的结构体
type Menu struct {
	Id         int64    `json:"id"`          //
	MenuPath   string   `json:"path"`        //菜单路径
	Component  string   `json:"component"`   //页面模块
	Redirect   string   `json:"redirect"`    //重定向地址
	Name       string   `json:"name"`        //唯一关键名
	Hidden     bool     `json:"hidden"`      //是否隐藏
	Alwaysshow bool     `json:"alwaysshow"`  //是否常显示
	Sort       int      `json:"sort"`        //排序
	ParentName string   `json:"parent_name"` //父菜级关键名
	AutoCreate bool     `json:"auto_create"` //是否自动生成
	MenuMeta   MenuMeta `json:"meta"`
}

// TableName 获取表名
func (t *Menu) TableName() string {
	return "menu"
}

// Insert 插入一条记录
func (t *Menu) Insert() (int64, error) {
	r, err := defDB.Insert("menu", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = id
	return id, err
}

// Update 更新对象
func (t *Menu) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	r, err := defDB.Update("menu", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *Menu) DeleteById(id int64) (int64, error) {
	r, err := defDB.Delete("menu", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *Menu) GetById(id int64) (Menu, error) {
	obj := Menu{}
	err := defDB.Table("menu").Where("id", id).Struct(&obj)
	return obj, err
}
