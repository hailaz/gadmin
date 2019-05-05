package api

import (
	"github.com/hailaz/gadmin/app/model"
)

type RoleController struct {
	BaseController
}

func (c *RoleController) Get() {
	type Role struct {
		Id       int64  `json:"id"`        //
		RoleName string `json:"role_name"` //
	}
	var list struct {
		List  []Role `json:"items"`
		Total int    `json:"total"`
	}
	roleList := make([]Role, 0)
	roles := model.Enforcer.GetAllRoles()

	for i, item := range roles {
		r := Role{Id: int64(i), RoleName: item}
		roleList = append(roleList, r)
	}
	list.List = roleList
	list.Total = len(roleList)
	Success(c.Controller, list)
}

func (c *RoleController) Post() {
	Success(c.Controller, "Post")
}

func (c *RoleController) Put() {
	Success(c.Controller, "Put")
}

func (c *RoleController) Delete() {
	Success(c.Controller, "Delete")
}
