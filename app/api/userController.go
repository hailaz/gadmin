package api

type UserController struct {
	BaseController
}

func (c *UserController) Show() {
	c.Response.Writeln("Controller Show")
}
