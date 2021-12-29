package template

var DalTmpl = `package dao

import (
	"fmt"
	"github.com/jtzjtz/kit/database"
	"{{.ProjectName}}/entity"
	"gorm.io/gorm"
)

type {{ .StructName}}DAO struct {
}


func (instance *{{ .StructName}}DAO) Create({{ .StructName | toLower}} entity.{{ .StructName}}) (entity.{{ .StructName}}, error) {

	db := SqlDBWrite
	err := db.Create(&{{ .StructName | toLower}}).Error
	return {{ .StructName | toLower}}, err
}

//Update(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
func (instance *{{ .StructName}}DAO) Update(data map[string]interface{}, conditions []database.SqlCondition) error {

	db := SqlDBWrite
	return instance.buildQuery(db, conditions).Model(&entity.{{ .StructName}}{}).Updates(data).Error

}

//Update(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
func (instance *{{ .StructName}}DAO) UpdateBySql(data map[string]interface{}, sqlQuery string) error {

	db := SqlDBWrite
	return db.Where(sqlQuery).Model(&entity.{{ .StructName}}{}).Updates(data).Error

}
func (instance *{{ .StructName}}DAO) Save({{ .StructName | toLower}} entity.{{ .StructName}}) error {
	db := SqlDBWrite
	return db.Save(&{{ .StructName | toLower}}).Error
}

func (instance *{{ .StructName}}DAO) Delete(conditions []database.SqlCondition) error {
	db := SqlDBWrite
	return instance.buildQuery(db, conditions).Model(&entity.{{ .StructName}}{}).Delete(entity.{{ .StructName}}{}).Error
}

func (instance *{{ .StructName}}DAO) DeleteBySql(sqlQuery string) error {
	db := SqlDBWrite
	return db.Where(sqlQuery).Model(&entity.{{ .StructName}}{}).Delete(entity.{{ .StructName}}{}).Error
}

func (instance *{{ .StructName}}DAO) GetBy{{.PrimaryKey}}({{.PrimaryKey | toLower}} {{.PrimaryType}}, isWriteDB bool, options ...database.SqlOptions) *entity.{{ .StructName}} {
	db := SqlDBRead
	if (isWriteDB) {
		db = SqlDBWrite
	}
	{{ .StructName | toLower}} := &entity.{{ .StructName}}{}
	condition := []database.SqlCondition{}
	condition = append(condition, struct {
		QueryName string
		Predicate database.SqlPredicate
		Value     interface{}
	}{QueryName: "{{.SqlPrimaryKey}}", Value: {{.PrimaryKey | toLower}}, Predicate: database.SqlEqualPredicate})
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
		if option.OrderBy != nil {
			for orderKey, ascDes := range option.OrderBy {
				db = db.Order(orderKey + " " + ascDes)
			}
		}
	}
	if instance.buildQuery(db, condition).First({{ .StructName | toLower}}).RowsAffected == 0{
		return nil
	}	
	return {{ .StructName | toLower}}
}


func (instance *{{ .StructName}}DAO) First(condition [] database.SqlCondition, isWriteDB bool, options ...database.SqlOptions) *entity.{{ .StructName}} {
	db := SqlDBRead
	if (isWriteDB) {
		db = SqlDBWrite
	}
	{{ .StructName | toLower}} := &entity.{{ .StructName}}{}
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
		if option.OrderBy != nil {
			for orderKey, ascDes := range option.OrderBy {
				db = db.Order(orderKey + " " + ascDes)
			}
		}
	}
	if instance.buildQuery(db, condition).First({{ .StructName | toLower}}).RowsAffected == 0{
		return nil
	}	
	
	return {{ .StructName | toLower}}
}

func (instance *{{ .StructName}}DAO) FirstBySql(sqlQuery string, isWriteDB bool, options ...database.SqlOptions) *entity.{{ .StructName}} {
	db := SqlDBRead
	if (isWriteDB) {
		db = SqlDBWrite
	}
	{{ .StructName | toLower}} := &entity.{{ .StructName}}{}
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
		if option.OrderBy != nil {
			for orderKey, ascDes := range option.OrderBy {
				db = db.Order(orderKey + " " + ascDes)
			}
		}
	}
	if db.Where(sqlQuery).First({{ .StructName | toLower}}).RowsAffected == 0{
		return nil
	}	
	
	return {{ .StructName | toLower}}
}

func (instance *{{ .StructName}}DAO) Count(condition [] database.SqlCondition, isWriteDB bool) int {
	var c int64
	db := SqlDBRead

	if (isWriteDB) {
		db = SqlDBWrite
	}
	instance.buildQuery(db, condition).Model(&entity.{{ .StructName}}{}).Count(&c)
	
	return int(c)
}

func (instance *{{ .StructName}}DAO) CountBySql(sqlQuery string, isWriteDB bool) int {
	var c int64
	db := SqlDBRead

	if (isWriteDB) {
		db = SqlDBWrite
	}
	db.Where(sqlQuery).Model(&entity.{{ .StructName}}{}).Count(&c)
	
	return int(c)
}

func (instance *{{ .StructName}}DAO) GetList(conditions []database.SqlCondition, isWriteDB bool, options ...database.SqlOptions) []entity.{{ .StructName}} {

	db := SqlDBRead
	{{ .StructName | toLower}}s := []entity.{{ .StructName}}{}
	if (isWriteDB) {
		db = SqlDBWrite
	}
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
		if option.OrderBy != nil {
			for orderKey, ascDes := range option.OrderBy {
				db = db.Order(orderKey + " " + ascDes)
			}
		}
	}
	instance.buildQuery(db, conditions).Find(&{{ .StructName | toLower}}s)
	return {{ .StructName | toLower}}s
}

func (instance *{{ .StructName}}DAO) GetListBySql(sqlQuery string, isWriteDB bool, options ...database.SqlOptions) []entity.{{ .StructName}} {

	db := SqlDBRead
	{{ .StructName | toLower}}s := []entity.{{ .StructName}}{}
	if (isWriteDB) {
		db = SqlDBWrite
	}
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
		if option.OrderBy != nil {
			for orderKey, ascDes := range option.OrderBy {
				db = db.Order(orderKey + " " + ascDes)
			}
		}
	}
	db.Where(sqlQuery).Find(&{{ .StructName | toLower}}s)
	return {{ .StructName | toLower}}s
}

func (instance *{{ .StructName}}DAO) GetPageList(conditions []database.SqlCondition, start int, limit int, orderBy map[string]string, isWriteDB bool, options ...database.SqlOptions) []entity.{{ .StructName}} {
	db := SqlDBRead
	if (isWriteDB) {
		db = SqlDBWrite
	}
	{{ .StructName | toLower}}s := []entity.{{ .StructName}}{}
	for orderKey, ascDes := range orderBy {
		db = db.Order(orderKey + " " + ascDes)
	}
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
	}
	instance.buildQuery(db, conditions).Offset(start).Limit(limit).Find(&{{ .StructName | toLower}}s)
	return {{ .StructName | toLower}}s
}

func (instance *{{ .StructName}}DAO) GetPageListBySql(sqlQuery string, start int, limit int, orderBy map[string]string, isWriteDB bool, options ...database.SqlOptions) []entity.{{ .StructName}} {
	db := SqlDBRead
	if (isWriteDB) {
		db = SqlDBWrite
	}
	{{ .StructName | toLower}}s := []entity.{{ .StructName}}{}
	for orderKey, ascDes := range orderBy {
		db = db.Order(orderKey + " " + ascDes)
	}
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
	}
	db.Where(sqlQuery).Offset(start).Limit(limit).Find(&{{ .StructName | toLower}}s)
	return {{ .StructName | toLower}}s
}

func (instance *{{ .StructName}}DAO) buildQuery(db *gorm.DB, condition [] database.SqlCondition) *gorm.DB {

	for _, where := range condition {
		if(where.Predicate==""){
			where.Predicate=database.SqlEqualPredicate
		}
		db = db.Where(fmt.Sprintf("%v %v ", where.QueryName, where.Predicate), where.Value)
	}

	return db
}

`
