package template

var ModelTmpl = `package entity

import (
    "database/sql"
    "time"

)

var (
    _ = time.Second
    _ = sql.LevelDefault
)

//go:generate gormgen -structs {{.StructName}} -output {{.TableName}}_gen.go
type {{.StructName}} struct {
    {{range .Fields}}{{.}}
    {{end}}
}

// TableName sets the insert table name for this struct type
func ({{.ShortStructName}} *{{.StructName}}) TableName() string {
	return "{{.TableName}}"
}
`
