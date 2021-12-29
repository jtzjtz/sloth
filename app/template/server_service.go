package template

var ServerServiceTmpl = `package {{.TableName}}_service

import (
	"{{.ProjectName}}/server/dao"
	"{{.ProjectName}}/entity"
	"{{.ProjectName}}/proto/{{.TableName}}_proto"
	"github.com/jtzjtz/kit/convert"
	"github.com/jtzjtz/kit/database"
)

type {{ .StructName}}Service struct {
}

//创建实体
func (instance *{{ .StructName}}Service) Create{{ .StructName}}({{.StructName|toLowerCamelCase}}Entity entity.{{ .StructName}}) (entity.{{ .StructName}}, error) {
	dao := dao.{{ .StructName}}DAO{}
	{{.StructName|toLowerCamelCase}}New,err := dao.Create({{.StructName|toLowerCamelCase}}Entity)
	return {{.StructName|toLowerCamelCase}}New, err
}

//更新实体
func (instance *{{ .StructName}}Service) Update{{ .StructName}}(query {{.TableName}}_proto.Query, {{.StructName|toLowerCamelCase}}Update map[string]interface{}) (err error) {
	dao := dao.{{ .StructName}}DAO{}
	if query.SqlQuery != "" {
		err = dao.UpdateBySql({{.StructName|toLowerCamelCase}}Update, query.SqlQuery)
	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &{{.TableName}}_proto.{{ .StructName}}{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		err = dao.Update({{.StructName|toLowerCamelCase}}Update, sqlCondition)

	}
	return err

}

//删除实体
func (instance *{{ .StructName}}Service) Delete{{ .StructName}}(query {{.TableName}}_proto.Query) (err error) {
	dao := dao.{{ .StructName}}DAO{}
	if query.SqlQuery != "" {
		err = dao.DeleteBySql(query.SqlQuery)

	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &{{.TableName}}_proto.{{ .StructName}}{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		err = dao.Delete(sqlCondition)
	}
	return err
}

//查询实体
func (instance *{{ .StructName}}Service) Get{{ .StructName}}(query {{.TableName}}_proto.Query) *entity.{{ .StructName}} {
	var isWriteDB bool
	isWriteDB = {{.TableName}}_proto.DB_WRITE == query.Db
	options := database.SqlOptions{}
	dao := dao.{{ .StructName}}DAO{}
	if query.OrderBy != nil {
		options.OrderBy = query.OrderBy
	}
	options.SelectField = query.SelectField
	if query.SqlQuery != "" {
		entity := dao.FirstBySql(query.SqlQuery, isWriteDB, options)
		return entity

	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &{{.TableName}}_proto.{{ .StructName}}{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		entity := dao.First(sqlCondition, isWriteDB, options)
		return entity

	}
}

//查询实体分页列表
func (instance *{{ .StructName}}Service) Get{{ .StructName}}PageList(query {{.TableName}}_proto.Query, page int, pageNum int) (list []entity.{{ .StructName}}) {
	var isWriteDB bool
	dao := dao.{{ .StructName}}DAO{}
	isWriteDB = {{.TableName}}_proto.DB_WRITE == query.Db
	orderBy := query.OrderBy
	start := 0
	limit := pageNum
	if page == 0 {
		page = 1
	}
	start = (page - 1) * pageNum
	options := database.SqlOptions{}
	options.SelectField = query.SelectField
	if query.SqlQuery != "" {
		list = dao.GetPageListBySql(query.SqlQuery, start, limit, orderBy, isWriteDB, options)
	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &{{.TableName}}_proto.{{ .StructName}}{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		list = dao.GetPageList(sqlCondition, start, limit, orderBy, isWriteDB, options)
	}
	return list

}

//查询实体列表
func (instance *{{ .StructName}}Service) Get{{ .StructName}}List(query {{.TableName}}_proto.Query) (list []entity.{{ .StructName}}) {
	var isWriteDB bool
	dao := dao.{{ .StructName}}DAO{}
	isWriteDB = {{.TableName}}_proto.DB_WRITE == query.Db
	options := database.SqlOptions{}
	if query.OrderBy != nil {
		options.OrderBy = query.OrderBy
	}
	options.SelectField = query.SelectField
	if query.SqlQuery != "" {
		list = dao.GetListBySql(query.SqlQuery, isWriteDB, options)
	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &{{.TableName}}_proto.{{ .StructName}}{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		list = dao.GetList(sqlCondition, isWriteDB, options)
	}
	return list
}

//查询实体个数
func (instance *{{ .StructName}}Service) Get{{ .StructName}}Count(query {{.TableName}}_proto.Query) (count int) {
	var isWriteDB bool
	dao := dao.{{ .StructName}}DAO{}
	isWriteDB = {{.TableName}}_proto.DB_WRITE == query.Db
	if query.SqlQuery != "" {
		count = dao.CountBySql(query.SqlQuery, isWriteDB)
	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &{{.TableName}}_proto.{{ .StructName}}{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		count = dao.Count(sqlCondition, isWriteDB)
	}
	return count
}


`
