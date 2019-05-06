package model

import (
	"errors"
	"fmt"

	"github.com/gogf/gf/g/os/glog"

	"github.com/gogf/gf/g/database/gdb"
)

type Policy struct {
	Policy     string `json:"policy"`     //
	Name       string `json:"name"`       //
	Descrption string `json:"descrption"` //
}

// GetPolicyList 获取权限列表
//
// createTime:2019年05月06日 17:24:12
// author:hailaz
func GetPolicyList(page, limit int, defaultname string) ([]Policy, int) {
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

// GetAllPolicy description
//
// createTime:2019年04月30日 10:20:50
// author:hailaz
func GetAllPolicy() (gdb.Result, error) {
	return defDB.Table("policy_name").All()
}

// GetPolicyByFullPath description
//
// createTime:2019年05月06日 15:53:08
// author:hailaz
func GetPolicyByFullPath(path string) (PolicyName, error) {
	obj := PolicyName{}
	err := defDB.Table("policy_name").Where("full_path", path).Struct(&obj)
	return obj, err
}

// UpdatePolicyByFullPath description
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
