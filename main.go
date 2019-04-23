package main

import (
	"github.com/gogf/gf/g"
)

func init() {

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
	s.SetServerRoot("public")
	s.SetPort(8080)
	s.Run()
}
