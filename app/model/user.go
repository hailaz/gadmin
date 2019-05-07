package model

import (
	"github.com/gogf/gf/g/crypto/gmd5"
	"github.com/gogf/gf/g/database/gdb"
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

// GetUserByPageLimt description
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
		} else {

		}

	}
	r, err := defDB.Table("user").Limit((page-1)*limit, (page-1)*limit+limit).Select()
	if err != nil {
		return nil, 0
	}
	r.ToStructs(&userList)
	return userList, total

}

// GetUserList description
//
// createTime:2019年04月30日 10:20:50
// author:hailaz
func GetUserList() (gdb.Result, error) {
	return defDB.Table("user").All()
}

// GetUserByName description
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
	return gmd5.EncryptString(data + ENCRYPTMD5)
}
