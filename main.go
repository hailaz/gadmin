package main

import (
	"github.com/gogf/gf/g"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/library/logger"
	"github.com/hailaz/gadmin/router"
)

func init() {
	// 设置默认配置文件，默认的 config.toml 将会被覆盖
	g.Config().SetFileName("config.toml")
	// 初始化数据库
	model.InitModel()
	// 初始化日志
	logger.InitLogger()

}

func main() {

	// u := new(model.User)

	// user, _ := u.GetById(3)

	// glog.Println(user)
	// u.UserName = gtime.Now().String()
	// glog.Println(u.Insert())
	// glog.Println(u)
	// u.UserName = "t" + gtime.Now().String()
	// glog.Println(u.Update())

	// glog.Println(u.DeleteById(5))

	// var users []model.User
	// // 或者 users := ([]User)(nil)
	// err := model.GetDB().Table("user").Structs(&users)
	// glog.Println(err, users)
	s := g.Server()
	s.SetIndexFolder(true)
	s.SetServerRoot("static")
	// 初始化路由
	router.InitRouter(s)

	port := g.Config().GetInt("port")
	if port == 0 {
		port = 8080
	}
	s.SetPort(port)
	s.Run()

}
