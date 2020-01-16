package model

import (
	"errors"

	"github.com/gogf/gf/os/gtime"
)

// User 表名：user
// 由数据库自动生成的结构体
type User struct {
	Id           int64  `json:"id"`           //
	Status       int    `json:"status"`       //状态
	UserName     string `json:"user_name"`    //用户名
	NickName     string `json:"nick_name"`    //昵称
	Password     string `json:"password"`     //密码
	Email        string `json:"email"`        //邮箱
	Phone        string `json:"phone"`        //手机号
	Sex          int    `json:"sex"`          //性别
	Age          int    `json:"age"`          //年龄
	AddTime      string `json:"add_time"`     //添加时间
	UpdateTime   string `json:"update_time"`  //修改时间
	AddUserId    int64  `json:"add_user_id"`  //操作用户
	Introduction string `json:"Introduction"` //介绍
	Avatar       string `json:"avatar"`       //头像
}

// TableName 获取表名
func (t *User) TableName() string {
	return "user"
}

// Insert 插入一条记录
func (t *User) Insert() (int64, error) {
	t.AddTime = gtime.Now().String()
	t.UpdateTime = gtime.Now().String()
	t.Password = EncryptPassword(t.Password)
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
