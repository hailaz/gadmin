package api

type UserController struct {
	BaseController
}

func (c *UserController) Show() {
	c.Response.Writeln("Controller Show")
}

func (c *UserController) Login() {
	c.Response.Writeln("Login Success")
}
