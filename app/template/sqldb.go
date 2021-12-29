package template

var SqlDBTmpl = `package dao


import "gorm.io/gorm"

//以下两个变量务必在项目里初始化
var (
	SqlDBWrite *gorm.DB
	SqlDBRead  *gorm.DB
)
`
