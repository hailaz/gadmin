package timer

import (
	"time"

	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/gogf/gf/g/os/gtimer"
	"github.com/hailaz/gadmin/app/api"
)

// InitTimer 初始化定时任务
//
// creatTime:2019年04月24日 14:50:34
// author:hailaz
func InitTimer() {

	gtimer.Add(time.Minute, RemoveTimeoutCryptoKey)

}

// RemoveTimeoutCryptoKey 移除超时加密key
//
// creatTime:2019年04月24日 15:42:43
// author:hailaz
func RemoveTimeoutCryptoKey() {
	kList := make([]interface{}, 0)
	nowSec := gtime.Second()
	//遍历加密key
	api.CryptoKeyList.Iterator(func(k interface{}, v interface{}) bool {
		ck := v.(api.CryptoKey)
		if nowSec-ck.TimeStamp >= 10 {
			kList = append(kList, k)
		}
		return true
	})
	//移除超时的加密key
	for _, v := range kList {
		api.CryptoKeyList.Remove(v)
		glog.Debugfln("remove key:%v", v)
	}
}
