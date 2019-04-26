package model

import "github.com/gogf/gf/g/crypto/gmd5"

const (
	ENCRYPTMD5 = "gadmin"
)

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
