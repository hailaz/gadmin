package logger

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

// InitLogger 初始化日志设置
//
// createTime:2019年05月13日 09:46:02
// author:hailaz
func InitLogger() {
	path := g.Config().GetString("logpath", "log")
	glog.SetPath(path)
	glog.SetFlags(glog.F_TIME_STD | glog.F_FILE_SHORT)
}
