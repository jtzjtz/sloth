package template

var ControllerTmpl = `package dal

import (
	"{{.ProjectName}}/app/entity"
	"{{.ProjectName}}/app/proto/oms_order_main"
	"context"
	"google.golang.org/grpc"
)

type {{ .StructName}}Service struct {
}


func (this *{{ .StructName}}DAL) Create({{ .StructName | toLower}} entity.{{ .StructName}}) entity.{{ .StructName}} {

	db := database.SqlDataBase_W
	db.Create(&{{ .StructName | toLower}})
	return {{ .StructName | toLower}}
}

//Update(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
func (this *{{ .StructName}}DAL) Update(data map[string]interface{}, conditions []database.SqlCondition) error {

	db := database.SqlDataBase_W
	return this.buildQuery(db, conditions).Model(&entity.{{ .StructName}}{}).Updates(data).Error

}

//Update(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
func (this *{{ .StructName}}DAL) UpdateBySql(data map[string]interface{}, sqlQuery string) error {

	db := database.SqlDataBase_W
	return db.Where(sqlQuery).Model(&entity.{{ .StructName}}{}).Updates(data).Error

}
func (this *{{ .StructName}}DAL) Save({{ .StructName | toLower}} entity.{{ .StructName}}) error {
	db := database.SqlDataBase_W
	return db.Save(&{{ .StructName | toLower}}).Error
}

func (this *{{ .StructName}}DAL) Delete(conditions []database.SqlCondition) error {
	db := database.SqlDataBase_W
	return this.buildQuery(db, conditions).Model(&entity.{{ .StructName}}{}).Delete(entity.{{ .StructName}}{}).Error
}

func (this *{{ .StructName}}DAL) DeleteBySql(sqlQuery string) error {
	db := database.SqlDataBase_W
	return db.Where(sqlQuery).Model(&entity.{{ .StructName}}{}).Delete(entity.{{ .StructName}}{}).Error
}

func (this *{{ .StructName}}DAL) GetBy{{.PrimaryKey}}({{.PrimaryKey | toLower}} {{.PrimaryType}}, isWriteDB bool) *entity.{{ .StructName}} {
	db := database.SqlDataBase_R
	if (isWriteDB) {
		db = database.SqlDataBase_W
	}
	{{ .StructName | toLower}} := &entity.{{ .StructName}}{}
	condition := []database.SqlCondition{}
	condition = append(condition, struct {
		QueryName string
		Predicate database.SqlPredicate
		Value     interface{}
	}{QueryName: "{{.SqlPrimaryKey}}", Value: {{.PrimaryKey | toLower}}, Predicate: database.SqlEqualPredicate})
	this.buildQuery(db, condition).First({{ .StructName | toLower}})
	return {{ .StructName | toLower}}
}


func (this *{{ .StructName}}DAL) First(condition [] database.SqlCondition, isWriteDB bool) *entity.{{ .StructName}} {
	db := database.SqlDataBase_R
	if (isWriteDB) {
		db = database.SqlDataBase_W
	}
	{{ .StructName | toLower}} := &entity.{{ .StructName}}{}

	this.buildQuery(db, condition).First({{ .StructName | toLower}})
	
	return {{ .StructName | toLower}}
}

func (this *{{ .StructName}}DAL) FirstBySql(sqlQuery string, isWriteDB bool) *entity.{{ .StructName}} {
	db := database.SqlDataBase_R
	if (isWriteDB) {
		db = database.SqlDataBase_W
	}
	{{ .StructName | toLower}} := &entity.{{ .StructName}}{}

	db.Where(sqlQuery).First({{ .StructName | toLower}})
	
	return {{ .StructName | toLower}}
}

func (this *{{ .StructName}}DAL) Count(condition [] database.SqlCondition, isWriteDB bool) int {
	var c int
	db := database.SqlDataBase_R

	if (isWriteDB) {
		db = database.SqlDataBase_W
	}
	this.buildQuery(db, condition).Model(&entity.{{ .StructName}}{}).Count(&c)
	
	return c
}

func (this *{{ .StructName}}DAL) CountBySql(sqlQuery string, isWriteDB bool) int {
	var c int
	db := database.SqlDataBase_R

	if (isWriteDB) {
		db = database.SqlDataBase_W
	}
	db.Where(sqlQuery).Model(&entity.{{ .StructName}}{}).Count(&c)
	
	return c
}

func (this *{{ .StructName}}DAL) GetList(conditions []database.SqlCondition, isWriteDB bool) []entity.{{ .StructName}} {

	db := database.SqlDataBase_R
	{{ .StructName | toLower}}s := []entity.{{ .StructName}}{}
	if (isWriteDB) {
		db = database.SqlDataBase_W
	}
	this.buildQuery(db, conditions).Find(&{{ .StructName | toLower}}s)
	return {{ .StructName | toLower}}s
}

func (this *{{ .StructName}}DAL) GetListBySql(sqlQuery string, isWriteDB bool) []entity.{{ .StructName}} {

	db := database.SqlDataBase_R
	{{ .StructName | toLower}}s := []entity.{{ .StructName}}{}
	if (isWriteDB) {
		db = database.SqlDataBase_W
	}
	db.Where(sqlQuery).Find(&{{ .StructName | toLower}}s)
	return {{ .StructName | toLower}}s
}

func (this *{{ .StructName}}DAL) GetListWithOrder(conditions []database.SqlCondition, orderby map[string]string, isWriteDB bool) []entity.{{ .StructName}} {

	db := database.SqlDataBase_R
	{{ .StructName | toLower}}s := []entity.{{ .StructName}}{}
	if (isWriteDB) {
		db = database.SqlDataBase_W
	}
	for orderKey, ascDes := range orderby {
		db = db.Order(orderKey + " " + ascDes)
	}
	this.buildQuery(db, conditions).Find(&{{ .StructName | toLower}}s)
	return {{ .StructName | toLower}}s
}

func (this *{{ .StructName}}DAL) GetListWithOrderBySql(sqlQuery string, orderby map[string]string, isWriteDB bool) []entity.{{ .StructName}} {

	db := database.SqlDataBase_R
	{{ .StructName | toLower}}s := []entity.{{ .StructName}}{}
	if (isWriteDB) {
		db = database.SqlDataBase_W
	}
	for orderKey, ascDes := range orderby {
		db = db.Order(orderKey + " " + ascDes)
	}
	db.Where(sqlQuery).Find(&{{ .StructName | toLower}}s)
	return {{ .StructName | toLower}}s
}

func (this *{{ .StructName}}DAL) GetPageList(conditions []database.SqlCondition, start int, limit int, orderby map[string]string, isWriteDB bool) []entity.{{ .StructName}} {
	db := database.SqlDataBase_R
	if (isWriteDB) {
		db = database.SqlDataBase_W
	}
	{{ .StructName | toLower}}s := []entity.{{ .StructName}}{}
	for orderKey, ascDes := range orderby {
		db = db.Order(orderKey + " " + ascDes)
	}
	this.buildQuery(db, conditions).Offset(start).Limit(limit).Find(&{{ .StructName | toLower}}s)
	return {{ .StructName | toLower}}s
}

func (this *{{ .StructName}}DAL) GetPageListBySql(sqlQuery string, start int, limit int, orderby map[string]string, isWriteDB bool) []entity.{{ .StructName}} {
	db := database.SqlDataBase_R
	if (isWriteDB) {
		db = database.SqlDataBase_W
	}
	{{ .StructName | toLower}}s := []entity.{{ .StructName}}{}
	for orderKey, ascDes := range orderby {
		db = db.Order(orderKey + " " + ascDes)
	}
	db.Where(sqlQuery).Offset(start).Limit(limit).Find(&{{ .StructName | toLower}}s)
	return {{ .StructName | toLower}}s
}

func (this *{{ .StructName}}DAL) buildQuery(db *gorm.DB, condition [] database.SqlCondition) *gorm.DB {

	for _, where := range condition {
		if(where.Predicate==""){
			where.Predicate=database.SqlEqualPredicate
		}
		db = db.Where(fmt.Sprintf("%v %v ", where.QueryName, where.Predicate), where.Value)
	}

	return db
}

`
