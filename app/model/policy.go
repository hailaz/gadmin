package model

import (
	"errors"
	"fmt"

	"github.com/gogf/gf/os/glog"

	"github.com/gogf/gf/database/gdb"
)

type Policy struct {
	Policy     string `json:"policy"`     //
	Name       string `json:"name"`       //
	Descrption string `json:"descrption"` //
}

type RolePolicy struct {
	Role string
	Path string
	Atc  string
}

// GetPolicyList 获取权限列表
//
// createTime:2019年05月06日 17:24:12
// author:hailaz
func GetPolicyList(page, limit int, defaultname string) ([]Policy, int) {
	if page < 1 {
		page = 1
	}
	policyList := make([]Policy, 0)
	policys := Enforcer.GetPermissionsForUser("system")
	total := len(policys)
	r, _ := GetAllPolicy()
	pn := make([]PolicyName, 0)
	r.ToStructs(&pn)

	for _, item := range policys {
		full := fmt.Sprintf("%v:%v", item[1], item[2])
		p := Policy{Policy: full, Name: defaultname}
		for _, itempn := range pn {
			if itempn.FullPath == full {
				p.Name = itempn.Name
				p.Descrption = itempn.Descrption
				break
			}
		}
		policyList = append(policyList, p)
	}
	if limit == -1 {
		return policyList, total
	}
	if len(policyList) < page*limit {
		if len(policyList) < limit {
			policyList = policyList
		} else {
			policyList = policyList[(page-1)*limit:]
		}
	} else {
		policyList = policyList[(page-1)*limit : (page-1)*limit+limit]
	}
	return policyList, total
}

// GetPolicyByRole 根据角色获取权限
//
// createTime:2019年05月07日 11:35:33
// author:hailaz
func GetPolicyByRole(role string) []Policy {
	policyList := make([]Policy, 0)
	policys := Enforcer.GetPermissionsForUser(role)
	glog.Debug(policys)
	for _, item := range policys {
		full := fmt.Sprintf("%v:%v", item[1], item[2])
		p := Policy{Policy: full}
		policyList = append(policyList, p)
	}
	return policyList
}

// GetAllPolicy 获取所有权限名称
//
// createTime:2019年04月30日 10:20:50
// author:hailaz
func GetAllPolicy() (gdb.Result, error) {
	return defDB.Table("policy_name").All()
}

// GetPolicyByFullPath 根据权限全路径获取权限
//
// createTime:2019年05月06日 15:53:08
// author:hailaz
func GetPolicyByFullPath(path string) (PolicyName, error) {
	obj := PolicyName{}
	err := defDB.Table("policy_name").Where("full_path", path).Struct(&obj)
	return obj, err
}

// UpdatePolicyByFullPath 更新权限信息
//
// createTime:2019年05月06日 15:47:35
// author:hailaz
func UpdatePolicyByFullPath(path, name string) error {
	p, err := GetPolicyByFullPath(path)
	// 不存在插入新数据
	if err != nil || p.Id == 0 {
		p.FullPath = path
		p.Name = name
		id, _ := p.Insert()
		if id > 0 {
			return nil
		} else {
			return errors.New("update fail")
		}
	}
	// 存在则更新
	p.Name = name
	i, err := p.Update()
	if err != nil {
		glog.Error(err)
		return err
	}
	if i < 0 {
		return errors.New("update fail")
	}
	return nil
}

// ReSetPolicy 更新路由权限
//
// createTime:2019年04月29日 17:30:26
// author:hailaz
func ReSetPolicy(role string, rmap map[string]RolePolicy) {
	old := Enforcer.GetPermissionsForUser(role)
	for _, item := range old {
		glog.Debug(item)
		full := fmt.Sprintf("%v %v %v", item[0], item[1], item[2])
		if _, ok := rmap[full]; ok { //从待插入列表中删除已存在的路由
			delete(rmap, full)
		} else { //删除不存在的旧路由
			Enforcer.DeletePermissionForUser(item[0], item[1], item[2])
			if role == "system" {
				p, _ := GetPolicyByFullPath(fmt.Sprintf("%v:%v", item[1], item[2]))
				if p.Id > 0 {
					p.DeleteById(p.Id)
				}
			}
		}
	}
	for _, item := range rmap { //插入新路由
		Enforcer.AddPolicy(item.Role, item.Path, item.Atc)
	}
}
