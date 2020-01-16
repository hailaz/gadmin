package model

import (
	"fmt"
	"runtime"

	"github.com/casbin/casbin/model"
	"github.com/casbin/casbin/persist"
	"github.com/gogf/gf/database/gdb"
)

type CasbinRule struct {
	Id    int64  `json:"id"`     //
	PType string `json:"p_type"` //
	V0    string `json:"v0"`     //
	V1    string `json:"v1"`     //
	V2    string `json:"v2"`     //
	V3    string `json:"v3"`     //
	V4    string `json:"v4"`     //
	V5    string `json:"v5"`     //
}

const (
	CASBINRULE_TABLE_NAME = "casbin_rule"
)

// Adapter represents the Xorm adapter for policy storage.
type Adapter struct {
	o gdb.DB
}

func NewAdapter(db gdb.DB) *Adapter {
	a := &Adapter{o: db}

	// Open the DB, create it if not existed.
	a.open()

	// Call the destructor when the object is released.
	runtime.SetFinalizer(a, finalizer)

	return a
}

// finalizer is the destructor for Adapter.
func finalizer(a *Adapter) {
}

func (a *Adapter) open() {

}

func (a *Adapter) close() {
	a.o = nil
}

func (a *Adapter) createTable() {
}

func (a *Adapter) dropTable() {
}

func loadPolicyLine(line CasbinRule, model model.Model) {
	lineText := line.PType
	if line.V0 != "" {
		lineText += ", " + line.V0
	}
	if line.V1 != "" {
		lineText += ", " + line.V1
	}
	if line.V2 != "" {
		lineText += ", " + line.V2
	}
	if line.V3 != "" {
		lineText += ", " + line.V3
	}
	if line.V4 != "" {
		lineText += ", " + line.V4
	}
	if line.V5 != "" {
		lineText += ", " + line.V5
	}

	persist.LoadPolicyLine(lineText, model)
}

// LoadPolicy loads policy from database.
func (a *Adapter) LoadPolicy(model model.Model) error {
	var lines []CasbinRule
	err := a.o.Table(CASBINRULE_TABLE_NAME).Structs(&lines)

	if err != nil {
		return err
	}

	for _, line := range lines {
		loadPolicyLine(line, model)
	}

	return nil
}

func savePolicyLine(ptype string, rule []string) CasbinRule {
	line := CasbinRule{}

	line.PType = ptype
	if len(rule) > 0 {
		line.V0 = rule[0]
	}
	if len(rule) > 1 {
		line.V1 = rule[1]
	}
	if len(rule) > 2 {
		line.V2 = rule[2]
	}
	if len(rule) > 3 {
		line.V3 = rule[3]
	}
	if len(rule) > 4 {
		line.V4 = rule[4]
	}
	if len(rule) > 5 {
		line.V5 = rule[5]
	}

	return line
}

// SavePolicy saves policy to database.
func (a *Adapter) SavePolicy(model model.Model) error {

	var lines []CasbinRule

	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			lines = append(lines, line)
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			lines = append(lines, line)
		}
	}

	_, err := a.o.BatchInsert(CASBINRULE_TABLE_NAME, lines)
	return err
}

// AddPolicy adds a policy rule to the storage.
func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	_, err := a.o.Insert(CASBINRULE_TABLE_NAME, &line)
	return err
}

// RemovePolicy removes a policy rule from the storage.
func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	qs := a.o.Table(CASBINRULE_TABLE_NAME).Safe()
	qs = qs.Where("p_type", ptype)
	for index := 0; index < len(rule); index++ {
		qs = qs.And(fmt.Sprintf("v%d", index), rule[index])
	}
	_, err := qs.Delete()
	return err

}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	qs := a.o.Table(CASBINRULE_TABLE_NAME).Safe()
	qs = qs.Where("p_type", ptype)
	for index := 0; index <= 5; index++ {
		if fieldIndex <= index && index < fieldIndex+len(fieldValues) {
			qs = qs.And(fmt.Sprintf("v%d", index), fieldValues[index-fieldIndex])
		}
	}
	_, err := qs.Delete()
	return err
}
