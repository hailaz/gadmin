package api

import (
	"github.com/gogf/gf/g/os/glog"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/library/code"
)

type RoleController struct {
	BaseController
}

func (c *RoleController) Get() {
	page := c.Request.GetInt("page", 1)
	limit := c.Request.GetInt("limit", 10)
	var list struct {
		List  []model.Role `json:"items"`
		Total int          `json:"total"`
	}
	list.List, list.Total = model.GetRoleList(page, limit, UNDEFIND_POLICY_NAME)

	Success(c.Controller, list)
}

func (c *RoleController) Post() {
	data := c.Request.GetJson()
	name := data.GetString("name")
	role := data.GetString("role")

	err := model.AddRole(role, name)
	if err != nil {
		Fail(c.Controller, code.RESPONSE_ERROR, err.Error())
	}

	Success(c.Controller, "Post")
}

func (c *RoleController) Put() {
	data := c.Request.GetJson()
	name := data.GetString("name")
	role := data.GetString("role")
	glog.Debug(name, role)
	if name == UNDEFIND_POLICY_NAME {
		Fail(c.Controller, code.RESPONSE_ERROR)
	} else {
		err := model.UpdateRoleByRoleKey(role, name)
		if err != nil {
			Fail(c.Controller, code.RESPONSE_ERROR, err.Error())
		}
	}
	Success(c.Controller, "修改成功")
}

func (c *RoleController) Delete() {
	data := c.Request.GetJson()
	role := data.GetString("role")

	err := model.DeleteRole(role)
	if err != nil {
		Fail(c.Controller, code.RESPONSE_ERROR, err.Error())
	}
	Success(c.Controller, "Delete")
}
