package model

import (
	"errors"
)

// RoleMenu 表名：role_menu
// 由数据库自动生成的结构体
type RoleMenu struct {
	Id       int64  `json:"id"`        //
	RoleKey  string `json:"role_key"`  //角色key
	MenuName string `json:"menu_name"` //菜单关键名
}

// TableName 获取表名
func (t *RoleMenu) TableName() string {
	return "role_menu"
}

// Insert 插入一条记录
func (t *RoleMenu) Insert() (int64, error) {
	r, err := defDB.Insert("role_menu", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = id
	return id, err
}

// Update 更新对象
func (t *RoleMenu) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	r, err := defDB.Update("role_menu", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *RoleMenu) DeleteById(id int64) (int64, error) {
	r, err := defDB.Delete("role_menu", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *RoleMenu) GetById(id int64) (RoleMenu, error) {
	obj := RoleMenu{}
	err := defDB.Table("role_menu").Where("id", id).Struct(&obj)
	return obj, err
}
