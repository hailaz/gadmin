package api

import (
	"encoding/base64"

	"github.com/hailaz/gadmin/library/code"

	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/library/common"
)

type LoginController struct {
	gmvc.Controller
}

// GetLoginCryptoKey description
//
// createTime:2019年04月24日 13:57:34
// author:hailaz
func (c *LoginController) GetLoginCryptoKey() {
	kid := c.Request.Session.Id()
	ck := common.GenCryptoKey(kid)

	glog.Debug("kid:" + kid)

	Success(c.Response, ck)

}

// Login 登录
//
// createTime:2019年04月24日 15:49:56
// author:hailaz
func (c *LoginController) Login() {
	data := c.Request.GetJson()
	name := data.GetString("username")
	pwd := data.GetString("password")
	kid := data.GetString("kid")

	if ck := common.GetCryptoKey(kid); ck != nil {
		if gtime.Second()-ck.TimeStamp >= 5 {
			Fail(c.Response, code.RESPONSE_ERROR, "密钥超时")
			return
		}
		//glog.Debugfln("%v", ck)
		//glog.Debugfln("%v", ck.Key)
		//glog.Debugfln("%v %v", name, pwd)
		decodePwd, err := base64.StdEncoding.DecodeString(pwd)
		if err != nil {
			Fail(c.Response, code.RESPONSE_ERROR)
			return
		}
		decryptPwd, _ := common.RsaDecrypt(decodePwd, []byte(ck.PrivateKey))
		glog.Debug(string(decryptPwd))
		password := string(decryptPwd)
		//glog.Debugfln("%v %v", name, password)
		if password != "" {
			u, err := model.GetUserByName(name)
			if err != nil {
				Fail(c.Response, 1)
				return
			}
			if u.Password == model.EncryptPassword(password) {
				token, _ := common.CreateJWT(u.UserName)
				var tk struct {
					Token string `json:"token"`
				}
				tk.Token = token
				Success(c.Response, tk)
				return
			}

		}
	}
	Fail(c.Response, code.RESPONSE_ERROR)
}
