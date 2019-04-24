package api

import (
	"github.com/gogf/gf/g/crypto/gaes"
	"github.com/gogf/gf/g/encoding/gbase64"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/os/gtime"

	"github.com/gogf/gf/g/container/gmap"
	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/util/grand"
)

type LoginController struct {
	gmvc.Controller
}

type CryptoKey struct {
	CryptoType string
	Key        string
	TimeStamp  int64 //sec
}

var CryptoKeyList = gmap.New()

// GetLoginCryptoKey description
//
// creatTime:2019年04月24日 13:57:34
// author:hailaz
func (c *LoginController) GetLoginCryptoKey() {
	sid := c.Request.Session.Id()
	ck := CryptoKey{CryptoType: "AES", Key: grand.RandLetters(16), TimeStamp: gtime.Second()}
	CryptoKeyList.Set(sid, ck)
	//c.Response.Writeln("sid:" + fmt.Sprintf("%v",sid))
	c.Response.WriteJson(ck)
}

// Login 登录
//
// creatTime:2019年04月24日 15:49:56
// author:hailaz
func (c *LoginController) Login() {
	sid := c.Request.Session.Id()
	var ck CryptoKey
	if k := CryptoKeyList.Get(sid); k != nil {
		ck = k.(CryptoKey)
		glog.Debugfln("%v", ck)
		data := c.Request.GetJson()
		name := data.GetString("name")
		pwd := data.GetString("pwd")
		glog.Debugfln("%v %v", name, pwd)
		password := DecryptPassword(pwd, ck.Key)
		glog.Debugfln("%v %v", name, password)
		if password != "" {
			c.Response.Writeln("success")
			return
		}
	}
	c.Response.Writeln("false")
}

// DecryptPassword 解密密码
// 输入base64编码的加密数据
// 算法模式：“CBC”
// 补码方式：“PKCS5Padding”
// 密钥偏移量：“I Love Go Frame!”
// creatTime:2019年04月24日 16:50:59
// author:hailaz
func DecryptPassword(data, key string) string {
	temp, err := gbase64.Decode(data)
	if err != nil {
		return ""
	}
	b, err := gaes.Decrypt([]byte(temp), []byte(key))
	if err != nil {
		return ""
	}
	return string(b)
}
