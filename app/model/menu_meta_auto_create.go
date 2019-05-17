package model

import (
	"errors"
)

// MenuMeta 表名：menu_meta
// 由数据库自动生成的结构体
type MenuMeta struct {
	Id       int64  `json:"id"`        //
	MenuName string `json:"menu_name"` //菜单关键名
	Title    string `json:"title"`     //名称
	Icon     string `json:"icon"`      //图标
	Nocache  bool   `json:"noCache"`   //是否缓存
}

// TableName 获取表名
func (t *MenuMeta) TableName() string {
	return "menu_meta"
}

// Insert 插入一条记录
func (t *MenuMeta) Insert() (int64, error) {
	r, err := defDB.Insert("menu_meta", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = id
	return id, err
}

// Update 更新对象
func (t *MenuMeta) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	r, err := defDB.Update("menu_meta", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *MenuMeta) DeleteById(id int64) (int64, error) {
	r, err := defDB.Delete("menu_meta", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *MenuMeta) GetById(id int64) (MenuMeta, error) {
	obj := MenuMeta{}
	err := defDB.Table("menu_meta").Where("id", id).Struct(&obj)
	return obj, err
}
