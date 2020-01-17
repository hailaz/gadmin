package api

import (
	"github.com/gogf/gf/database/gdb"

	"github.com/gogf/gf/os/glog"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/library/code"
)

type UserController struct {
	BaseController
}

func (c *UserController) Info() {
	u := c.GetUser()
	if u != nil {
		Success(c.Request, u.GetUserInfo())
	}
	Fail(c.Request, code.RESPONSE_ERROR, "获取用户信息失败")
}

func (c *UserController) Menu() {
	rolename := c.Request.GetString("rolename")
	if rolename != "" {
		var list struct {
			Menus     []model.MenuOut `json:"menus"`
			RoleMenus []model.MenuOut `json:"role_menus"`
		}
		list.Menus = model.GetMenuByRoleName([]string{model.ADMIN_NAME})
		list.RoleMenus = model.GetMenuByRoleName([]string{rolename})
		Success(c.Request, list)
	}
	u := c.GetUser()
	if u != nil {
		if u.UserName == model.ADMIN_NAME {
			Success(c.Request, model.GetMenuByRoleName([]string{model.ADMIN_NAME}))
		} else {
			roles, _ := model.Enforcer.GetRolesForUser(u.UserName)
			Success(c.Request, model.GetMenuByRoleName(roles))
		}

	}
	Fail(c.Request, code.RESPONSE_ERROR, "获取用户菜单失败")
}

func (c *UserController) Get() {
	page := c.Request.GetInt("page", 1)
	limit := c.Request.GetInt("limit", 10)
	var userList struct {
		List  []model.UserOut `json:"items"`
		Total int             `json:"total"`
	}
	userList.List, userList.Total = model.GetUserByPageLimt(page, limit)
	Success(c.Request, userList)
}
func (c *UserController) Post() {
	data, _ := c.Request.GetJson()
	username := data.GetString("user_name")

	nickname := data.GetString("nick_name")
	email := data.GetString("email")
	password := data.GetString("password")
	passwordconfirm := data.GetString("passwordconfirm")
	phone := data.GetString("phone")

	u, err := model.GetUserByName(username)
	if u != nil || err == nil {
		Fail(c.Request, code.RESPONSE_ERROR, "用户已存在")
	}
	if password == "" {
		Fail(c.Request, code.RESPONSE_ERROR, "密码为空")
	}
	if password != passwordconfirm {
		Fail(c.Request, code.RESPONSE_ERROR, "输入密码不一致")
	}
	addu := c.GetUser()
	var addUserId int64 = 0
	if addu != nil {
		addUserId = addu.Id
	}
	user := model.User{UserName: username, Password: password, NickName: nickname, Email: email, Phone: phone, AddUserId: addUserId}
	uid, _ := user.Insert()
	if uid > 0 {
		Success(c.Request, "success")
	}

	glog.Debug(uid)
	glog.Debug(data.ToJsonString())
	Fail(c.Request, code.RESPONSE_ERROR)
}
func (c *UserController) Put() {
	data, _ := c.Request.GetJson()
	username := data.GetString("user_name")
	nickname := data.GetString("nick_name")
	email := data.GetString("email")
	password := data.GetString("password")
	passwordconfirm := data.GetString("passwordconfirm")
	phone := data.GetString("phone")

	u, err := model.GetUserByName(username)
	if err != nil || u.Id == 0 {
		Fail(c.Request, code.RESPONSE_ERROR, "用户不存在")
	}
	umap := gdb.Map{}
	if nickname != u.NickName && nickname != "" {
		umap["nick_name"] = nickname
	}
	if email != u.Email && email != "" {
		umap["email"] = email
	}
	if phone != u.Phone && phone != "" {
		umap["phone"] = phone
	}
	if password == "" {
		err := model.UpdateUserById(u.Id, umap)
		if err != nil {
			Fail(c.Request, code.RESPONSE_ERROR, err.Error())
		}
	} else {
		if password != passwordconfirm {
			Fail(c.Request, code.RESPONSE_ERROR, "输入密码不一致")
		}
		umap["password"] = model.EncryptPassword(password)
		err := model.UpdateUserById(u.Id, umap)
		if err != nil {
			Fail(c.Request, code.RESPONSE_ERROR, err.Error())
		}
	}

	Success(c.Request, "success")
}
func (c *UserController) Delete() {
	data, _ := c.Request.GetJson()
	id := data.GetInt64("id")
	if id < 1 {
		Fail(c.Request, code.RESPONSE_ERROR)
	}
	u := new(model.User)
	user, err := u.GetById(id)
	if err != nil {
		Fail(c.Request, code.RESPONSE_ERROR, err.Error())
	}
	if user.UserName == model.ADMIN_NAME {
		Fail(c.Request, code.RESPONSE_ERROR, "无权限")
	}
	res, _ := u.DeleteById(id)
	if res <= 0 {
		Fail(c.Request, code.RESPONSE_ERROR)
	}
	model.Enforcer.DeleteRolesForUser(user.UserName)
	Success(c.Request, "success")
}
