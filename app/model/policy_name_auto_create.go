package model

import (
	"errors"
)

// PolicyName 表名：policy_name
// 由数据库自动生成的结构体
type PolicyName struct {
	Id         int64  `json:"id"`         //
	FullPath   string `json:"full_path"`  //权限完整路径
	Name       string `json:"name"`       //权限名称
	Descrption string `json:"descrption"` //描述
}

// TableName 获取表名
func (t *PolicyName) TableName() string {
	return "policy_name"
}

// Insert 插入一条记录
func (t *PolicyName) Insert() (int64, error) {
	r, err := defDB.Insert("policy_name", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = id
	return id, err
}

// Update 更新对象
func (t *PolicyName) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	r, err := defDB.Update("policy_name", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *PolicyName) DeleteById(id int64) (int64, error) {
	r, err := defDB.Delete("policy_name", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *PolicyName) GetById(id int64) (PolicyName, error) {
	obj := PolicyName{}
	err := defDB.Table("policy_name").Where("id", id).Struct(&obj)
	return obj, err
}
