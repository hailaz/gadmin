package api

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/library/code"
)

type MenuController struct {
	BaseController
}

func (c *MenuController) Get() {
	page := c.Request.GetInt("page", 1)
	limit := c.Request.GetInt("limit", 10)
	var list struct {
		List  []model.MenuOut `json:"items"`
		Total int             `json:"total"`
	}
	list.List, list.Total = model.GetMenuList(page, limit)

	Success(c.Request, list)
}

func (c *MenuController) Post() {
	data, _ := c.Request.GetJson()
	m := model.MenuOut{}
	data.ToStruct(&m)
	model.InsertMenuWithMeta(gdb.List{
		{
			"name":        m.Name,
			"menu_path":   m.MenuPath,
			"component":   m.Component,
			"sort":        m.Sort,
			"parent_name": m.ParentName,
			"hidden":      m.Hidden,
			"redirect":    m.Redirect,
			"alwaysshow":  m.Alwaysshow,
			"meta": gdb.Map{
				"title":   m.MenuMetaOut.Title,
				"icon":    m.MenuMetaOut.Icon,
				"noCache": m.MenuMetaOut.Nocache,
			},
		},
	})
	Success(c.Request, "添加成功")
}

func (c *MenuController) Put() {
	data, _ := c.Request.GetJson()
	m := model.MenuOut{}
	data.ToStruct(&m)
	err := model.UpdateMenuByName(
		m.Name,
		gdb.Map{
			"menu_path":   m.MenuPath,
			"component":   m.Component,
			"sort":        m.Sort,
			"parent_name": m.ParentName,
			"hidden":      m.Hidden,
			"redirect":    m.Redirect,
			"alwaysshow":  m.Alwaysshow,
			"meta": gdb.Map{
				"title":   m.MenuMetaOut.Title,
				"icon":    m.MenuMetaOut.Icon,
				"noCache": m.MenuMetaOut.Nocache,
			},
		},
	)
	if err != nil {
		Fail(c.Request, code.RESPONSE_ERROR, err.Error())
	}
	Success(c.Request, "修改成功")
}

func (c *MenuController) Delete() {
	data, _ := c.Request.GetJson()
	name := data.GetString("name")
	m, err := model.GetMenuByName(name)
	if err != nil {
		Fail(c.Request, code.RESPONSE_ERROR, err.Error())
	}
	if m.AutoCreate {
		Fail(c.Request, code.RESPONSE_ERROR)
	}
	res, _ := m.DeleteById(m.Id)
	if res <= 0 {
		Fail(c.Request, code.RESPONSE_ERROR)
	}
	Success(c.Request, "Delete")
}
