package logger

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/os/glog"
)

func InitLogger() {
	path := g.Config().GetString("logpath")
	if path == "" {
		path = "log"
	}
	glog.SetPath(path)
}
