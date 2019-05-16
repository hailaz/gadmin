package model

import (
	"errors"
)

// MenuMate 表名：menu_mate
// 由数据库自动生成的结构体
type MenuMate struct {
	Id       int64  `json:"id"`        //
	MenuName string `json:"menu_name"` //菜单关键名
	Title    string `json:"title"`     //名称
	Icon     string `json:"icon"`      //图标
	Nocache  bool   `json:"noCache"`   //是否缓存
}

// TableName 获取表名
func (t *MenuMate) TableName() string {
	return "menu_mate"
}

// Insert 插入一条记录
func (t *MenuMate) Insert() (int64, error) {
	r, err := defDB.Insert("menu_mate", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = id
	return id, err
}

// Update 更新对象
func (t *MenuMate) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	r, err := defDB.Update("menu_mate", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *MenuMate) DeleteById(id int64) (int64, error) {
	r, err := defDB.Delete("menu_mate", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *MenuMate) GetById(id int64) (MenuMate, error) {
	obj := MenuMate{}
	err := defDB.Table("menu_mate").Where("id", id).Struct(&obj)
	return obj, err
}
