package logger

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/os/glog"
)

// InitLogger 初始化日志设置
//
// createTime:2019年05月13日 09:46:02
// author:hailaz
func InitLogger() {
	path := g.Config().GetString("logpath")
	if path == "" {
		path = "log"
	}
	glog.SetPath(path)
}
