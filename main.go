package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/library/logger"
	"github.com/hailaz/gadmin/library/timer"
	"github.com/hailaz/gadmin/router"
)

func init() {
	// 设置默认配置文件，默认的 config.toml 将会被覆盖
	g.Config().SetFileName("config.toml")
	// 初始化数据库
	model.InitModel()
	// 初始化日志
	logger.InitLogger()

	timer.InitTimer()

}

func main() {
	s := g.Server()
	s.SetIndexFolder(false)
	s.SetIndexFiles([]string{"index.html"})
	s.SetServerRoot(".")
	// 初始化路由
	router.InitRouter(s)

	s.SetPort(g.Config().GetInt("port", 8080))
	s.Run()

}
