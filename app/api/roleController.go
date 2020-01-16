package api

import (
	"github.com/gogf/gf/os/glog"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/library/code"
)

type RoleController struct {
	BaseController
}

func (c *RoleController) Get() {
	page := c.Request.GetInt("page", 1)
	limit := c.Request.GetInt("limit", 10)
	username := c.Request.GetString("username")
	var list struct {
		List         []model.Role `json:"items"`
		UserRoleList []model.Role `json:"role_items"`
		Total        int          `json:"total"`
	}
	list.List, list.Total = model.GetRoleList(page, limit, UNDEFIND_POLICY_NAME)
	if username != "" {
		list.UserRoleList = model.GetRoleByUserName(username)
	}

	Success(c.Request, list)
}

func (c *RoleController) Post() {
	data, _ := c.Request.GetJson()
	name := data.GetString("name")
	role := data.GetString("role")

	err := model.AddRole(role, name)
	if err != nil {
		Fail(c.Request, code.RESPONSE_ERROR, err.Error())
	}

	Success(c.Request, "Post")
}

func (c *RoleController) Put() {
	data, _ := c.Request.GetJson()
	name := data.GetString("name")
	role := data.GetString("role")
	glog.Debug(name, role)
	if name == UNDEFIND_POLICY_NAME {
		Fail(c.Request, code.RESPONSE_ERROR)
	} else {
		err := model.UpdateRoleByRoleKey(role, name)
		if err != nil {
			Fail(c.Request, code.RESPONSE_ERROR, err.Error())
		}
	}
	Success(c.Request, "修改成功")
}

func (c *RoleController) Delete() {
	data, _ := c.Request.GetJson()
	role := data.GetString("role")

	err := model.DeleteRole(role)
	if err != nil {
		Fail(c.Request, code.RESPONSE_ERROR, err.Error())
	}
	Success(c.Request, "Delete")
}

func (c *RoleController) SetRoleByUserName() {
	data, _ := c.Request.GetJson()
	roles := data.GetStrings("roles")
	username := data.GetString("username")
	model.SetRoleByUserName(username, roles)

	Success(c.Request, "success")
}

func (c *RoleController) SetRoleMenus() {
	data, _ := c.Request.GetJson()
	role := data.GetString("role")
	menus := data.GetStrings("menus")
	model.SetRoleMenus(role, menus)
	Success(c.Request, "success")
}
