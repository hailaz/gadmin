package api

import (
	"encoding/base64"

	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/library/code"
	"github.com/hailaz/gadmin/library/common"
)

type UserController struct {
	BaseController
}

func (c *UserController) Show() {
	c.Response.Writeln("Controller Show")
}

func (c *UserController) List() {
	type User struct {
		Id           int64  `json:"id"`             //
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
	var userList struct {
		List  []User `json:"items"`
		Total int    `json:"total"`
	}
	r, err := model.GetUserList()
	if err != nil {
		glog.Error(err)
		Fail(c.Controller, code.RESPONSE_ERROR)
	}
	r.ToStructs(&userList.List)
	Success(c.Controller, userList)
}

func (c *UserController) AddUser() {
	user := c.Request.GetString("user")
	pwd := c.Request.GetString("pwd")
	u := model.User{UserName: user, Password: pwd}
	u.Insert()

	Success(c.Controller, "success")
}

// GetLoginCryptoKey description
//
// createTime:2019年04月24日 13:57:34
// author:hailaz
func (c *UserController) GetLoginCryptoKey() {
	kid := c.Request.Session.Id()
	ck := common.GenCryptoKey(kid)

	glog.Debug("kid:" + kid)

	Success(c.Controller, ck)

}

// Login 登录
//
// createTime:2019年04月24日 15:49:56
// author:hailaz
func (c *UserController) Login() {
	data := c.Request.GetJson()
	name := data.GetString("username")
	pwd := data.GetString("password")
	kid := data.GetString("kid")

	if ck := common.GetCryptoKey(kid); ck != nil {
		if gtime.Second()-ck.TimeStamp >= 5 {
			Fail(c.Controller, code.RESPONSE_ERROR, "密钥超时")
			return
		}
		glog.Debugfln("%v", ck.Id)
		//glog.Debugfln("%v", ck.Key)
		glog.Debugfln("%v %v", name, pwd)
		decodePwd, err := base64.StdEncoding.DecodeString(pwd)
		if err != nil {
			Fail(c.Controller, code.RESPONSE_ERROR)
			return
		}
		decryptPwd, _ := common.RsaDecrypt(decodePwd, []byte(ck.PrivateKey))
		glog.Debug(string(decryptPwd))
		password := string(decryptPwd)
		//glog.Debugfln("%v %v", name, password)
		if password != "" {
			u, err := model.GetUserByName(name)
			if err != nil {
				glog.Error(err)
				Fail(c.Controller, 1)
				return
			}
			if u.Password == model.EncryptPassword(password) {
				token, _ := common.CreateJWT(u.UserName)
				var tk struct {
					Token string `json:"token"`
				}
				tk.Token = token
				Success(c.Controller, tk)
				return
			}

		}
	}
	Fail(c.Controller, code.RESPONSE_ERROR)
}

// {
//     roles: ['admin'],
//     introduction: 'I am a super administrator',
//     avatar: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif',
//     name: 'Super Admin'
//   }
func (c *UserController) Info() {
	var info struct {
		Roles        []string `json:"roles"`
		Introduction string   `json:"introduction"`
		Avatar       string   `json:"avatar"`
		Name         string   `json:"name"`
	}
	info.Roles = []string{"admin"}
	info.Introduction = "I am a super administrator"
	info.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	info.Name = "Super Admin"
	Success(c.Controller, info)
}

func (c *UserController) Logout() {
	Success(c.Controller, "success")
}

func (c *UserController) Get() {
	page := c.Request.GetInt("page", 1)
	limit := c.Request.GetInt("limit", 10)
	var userList struct {
		List  []model.UserOut `json:"items"`
		Total int             `json:"total"`
	}
	userList.List, userList.Total = model.GetUserByPageLimt(page, limit)
	Success(c.Controller, userList)
}
func (c *UserController) Post() {
	Success(c.Controller, "success")
}
func (c *UserController) Put() {
	Success(c.Controller, "success")
}
func (c *UserController) Delete() {
	data := c.Request.GetJson()
	id := data.GetInt64("id")
	if id < 1 {
		Fail(c.Controller, code.RESPONSE_ERROR)
	}
	u := new(model.User)
	user, err := u.GetById(id)
	if err != nil {
		Fail(c.Controller, code.RESPONSE_ERROR, err.Error())
	}
	if user.UserName == "admin" {
		Fail(c.Controller, code.RESPONSE_ERROR, "无权限")
	}
	res, _ := u.DeleteById(id)
	if res < 0 {
		Fail(c.Controller, code.RESPONSE_ERROR)
	}
	Success(c.Controller, "success")
}
