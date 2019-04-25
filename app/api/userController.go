package api

import (
	"github.com/hailaz/gadmin/app/model"
)

type UserController struct {
	BaseController
}

func (c *UserController) Show() {
	c.Response.Writeln("Controller Show")
}

func (c *UserController) AddUser() {
	user := c.Request.GetString("user")
	pwd := c.Request.GetString("pwd")
	u := model.User{UserName: user, Password: pwd}
	u.Insert()
	c.Response.Writeln("Success")
}
