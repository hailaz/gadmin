package model

import (
	"errors"

	"github.com/gogf/gf/g/os/gtime"
)

// User 表名：user
// 由数据库自动生成的结构体
type User struct {
	Id           int64  `json:"id"`             //
	UserName     string `json:"user_name"`      //
	NickName     string `json:"nick_name"`      //
	Password     string `json:"password"`       //
	Email        string `json:"email"`          //
	Phone        string `json:"phone"`          //
	Qq           string `json:"qq"`             //
	Sex          int    `json:"sex"`            //
	Age          int    `json:"age"`            //
	AddTime      string `json:"add_time"`       //
	UpdateTime   string `json:"update_time"`    //
	AddUserId    int64  `json:"add_user_id"`    //
	ThirdPartyId int64  `json:"third_party_id"` //
}

// TableName 获取表名
func (t *User) TableName() string {
	return "user"
}

// Insert 插入一条记录
func (t *User) Insert() (int64, error) {
	t.AddTime = gtime.Now().String()
	t.UpdateTime = gtime.Now().String()
	r, err := defDB.Insert("user", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = id
	return id, err
}

// Update 更新对象
func (t *User) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	t.UpdateTime = gtime.Now().String()
	r, err := defDB.Update("user", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *User) DeleteById(id int64) (int64, error) {
	r, err := defDB.Delete("user", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *User) GetById(id int64) (User, error) {
	obj := User{}
	err := defDB.Table("user").Where("id", id).Struct(&obj)
	return obj, err
}
