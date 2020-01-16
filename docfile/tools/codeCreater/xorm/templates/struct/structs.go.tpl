package {{.Models}}
import (
{{$ilen := len .Imports}}
{{if gt $ilen 0}}

	{{range .Imports}}"{{.}}"{{end}}

{{end}}
{{range .Tables}}
{{if (PKIsId .)}}"errors"{{end}}
"github.com/gogf/gf/os/gtime"
)
// {{Mapper .Name}} 表名：{{.Name}}
// 由数据库自动生成的结构体
type {{Mapper .Name}} struct { {{$table := .}}
{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}	{{Mapper $col.Name}}	{{Type $col}} {{Tag $table $col}}
{{end}}}



// TableName 获取表名
func (t *{{Mapper .Name}}) TableName() string {
	return "{{.Name}}"
}

// Insert 插入一条记录
func (t *{{Mapper .Name}}) Insert() (int64, error) {
	t.AddTime = gtime.Now().String()
	t.UpdateTime = gtime.Now().String()
	r, err := defDB.Insert("{{.Name}}", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = id
	return id, err
}
{{if not (PKIsId .)}}/*{{end}}
// Update 更新对象
func (t *{{Mapper .Name}}) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	t.UpdateTime = gtime.Now().String()
	r, err := defDB.Update("{{.Name}}", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *{{Mapper .Name}}) DeleteById(id int64) (int64, error) {
	r, err := defDB.Delete("{{.Name}}", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}


// GetById 通过id查询记录
func (t *{{Mapper .Name}}) GetById(id int64) ({{Mapper .Name}}, error) {
	obj := {{Mapper .Name}}{}
	err := defDB.Table("{{.Name}}").Where("id", id).Struct(&obj)
	return obj, err
}
{{if not (PKIsId .)}}*/{{end}}
{{end}}