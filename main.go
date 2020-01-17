package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/library/logger"
	"github.com/hailaz/gadmin/library/timer"
	"github.com/hailaz/gadmin/router"
)

func init() {
	cfg := gcmd.GetOpt("c", "config.default.toml")

	if !gfile.Exists(cfg) && !gfile.Exists("config/"+cfg) {
		glog.Fatalf("config file(%s) no exist", cfg)
	}

	// 设置默认配置文件，默认的 config.toml 将会被覆盖
	g.Config().SetFileName(cfg)

	// 初始化日志
	logger.InitLogger()

	// 测试数据库
	err := g.DB().PingMaster()
	if err != nil {
		glog.Fatalf("DB err:%v", err)
	}

	// 初始化数据库
	model.InitModel()

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
