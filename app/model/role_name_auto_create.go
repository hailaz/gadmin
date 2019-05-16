package model

import (
	"errors"
)

// RoleName 表名：role_name
// 由数据库自动生成的结构体
type RoleName struct {
	Id         int64  `json:"id"`         //
	RoleKey    string `json:"role_key"`   //角色key
	Name       string `json:"name"`       //角色名
	Descrption string `json:"descrption"` //描述
}

// TableName 获取表名
func (t *RoleName) TableName() string {
	return "role_name"
}

// Insert 插入一条记录
func (t *RoleName) Insert() (int64, error) {
	r, err := defDB.Insert("role_name", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = id
	return id, err
}

// Update 更新对象
func (t *RoleName) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	r, err := defDB.Update("role_name", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *RoleName) DeleteById(id int64) (int64, error) {
	r, err := defDB.Delete("role_name", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *RoleName) GetById(id int64) (RoleName, error) {
	obj := RoleName{}
	err := defDB.Table("role_name").Where("id", id).Struct(&obj)
	return obj, err
}
