package model

import (
	"errors"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

const (
	ENCRYPTMD5 = "gadmin"
)

type UserOut struct {
	Id           int64  `json:"id"`             //
	Status       int    `json:"status"`         //
	UserName     string `json:"user_name"`      //
	NickName     string `json:"nick_name"`      //
	Email        string `json:"email"`          //
	Phone        string `json:"phone"`          //
	Sex          int    `json:"sex"`            //
	Age          int    `json:"age"`            //
	AddTime      string `json:"add_time"`       //
	UpdateTime   string `json:"update_time"`    //
	AddUserId    int64  `json:"add_user_id"`    //
	ThirdPartyId int64  `json:"third_party_id"` //
	Introduction string `json:"Introduction"`   //
	Avatar       string `json:"avatar"`         //
}

type UserInfo struct {
	Roles        []string `json:"roles"`
	Introduction string   `json:"introduction"`
	Avatar       string   `json:"avatar"`
	Name         string   `json:"name"`
}

// GetUserByPageLimt 获取用户
//
// createTime:2019年05月07日 16:11:41
// author:hailaz
func GetUserByPageLimt(page, limit int) ([]UserOut, int) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	total, _ := defDB.Table("user").Count()
	if total == 0 {
		return nil, 0
	}

	userList := make([]UserOut, 0)
	if total < page*limit {
		if total < limit {
			page = 1
		}
	}
	r, err := defDB.Table("user").Limit((page-1)*limit, (page-1)*limit+limit).Select()
	if err != nil {
		return nil, 0
	}
	r.ToStructs(&userList)
	return userList, total

}

// GetAllUser 获取所有用户
//
// createTime:2019年04月30日 10:20:50
// author:hailaz
func GetAllUser() (gdb.Result, error) {
	return defDB.Table("user").All()
}

// GetUserByName 根据用户名获取用户
//
// createTime:2019年04月23日 17:14:22
// author:hailaz
func GetUserByName(name string) (*User, error) {
	u := User{}
	err := defDB.Table("user").Where("user_name", name).Struct(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// EncryptPassword 加密密码
//
// createTime:2019年04月25日 10:19:13
// author:hailaz
func EncryptPassword(data string) string {
	md5Str, _ := gmd5.EncryptString(data + ENCRYPTMD5)
	return md5Str
}

// UpdateUserById 更新用户
//
// createTime:2019年05月08日 14:28:18
// author:hailaz
func UpdateUserById(id int64, udmap gdb.Map) error {
	udmap["update_time"] = gtime.Now().String()
	r, err := defDB.Table("user").Data(udmap).Where("id=?", id).Update()
	if err != nil {
		return err
	}
	i, _ := r.RowsAffected()
	if i < 0 {
		return errors.New("update fail")
	}
	return nil
}

// {
//     roles: ['admin'],
//     introduction: 'I am a super administrator',
//     avatar: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif',
//     name: 'Super Admin'
//   }
// GetUserInfo 获取用户信息
//
// createTime:2019年05月08日 16:53:24
// author:hailaz
func (u *User) GetUserInfo() UserInfo {
	info := UserInfo{}
	if u.UserName == ADMIN_NAME {
		info.Roles = []string{ADMIN_NAME}
	} else {
		info.Roles, _ = Enforcer.GetRolesForUser(u.UserName)
	}

	info.Avatar = u.Avatar
	info.Introduction = u.Introduction
	info.Name = u.NickName

	return info
}
