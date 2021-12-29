package template

var ServerHttpControllerTmpl = `package {{.TableName}}_controller

import (
	"{{.ProjectName}}/proto/{{.TableName}}_proto"
	"github.com/jtzjtz/kit/convert"
	"github.com/gin-gonic/gin"
	"{{.ProjectName}}/entity"
	"{{.ProjectName}}/server/service/{{.TableName}}_service"
	"strconv"
)

// 新增
func Post{{ .StructName}}(ctx *gin.Context) {
	result := entity.Result{}
	// 参数绑定
	{{.StructName|toLowerCamelCase}}Data := entity.{{ .StructName}}{}
	err := ctx.ShouldBind(&{{.StructName|toLowerCamelCase}}Data)
	if err != nil {
		result.SetCode(entity.CODE_ERROR).SetMessage(err.Error())
		ctx.JSON(200, result)
		return
	}
	// 创建
	{{.StructName|toLowerCamelCase}}Service := {{.TableName}}_service.{{ .StructName}}Service{}
	{{.StructName|toLowerCamelCase}}CreateEntity, err := {{.StructName|toLowerCamelCase}}Service.Create{{ .StructName}}({{.StructName|toLowerCamelCase}}Data)
	if err != nil {
		result.SetCode(entity.CODE_ERROR).SetMessage(err.Error())
		ctx.JSON(200, result)
		return
	}
	result.SetCode(entity.CODE_SUCCESS).SetMessage("success").SetData({{.StructName|toLowerCamelCase}}CreateEntity)
	ctx.JSON(200, result)
}

// 获取单条
func Get{{ .StructName}}(ctx *gin.Context) {
	result := entity.Result{}
	// 获取参数
	primaryKey := ctx.Param("id")
	// 查询条件
	query := {{.TableName}}_proto.Query{}
	query.SqlQuery = "{{.SqlPrimaryKey}} =" + primaryKey
	// 查询数据
	{{.StructName|toLowerCamelCase}}Service := {{.TableName}}_service.{{ .StructName}}Service{}
	{{.StructName|toLowerCamelCase}}Entity := {{.StructName|toLowerCamelCase}}Service.Get{{ .StructName}}(query)
	result.SetCode(entity.CODE_SUCCESS).SetMessage("success").SetData({{.StructName|toLowerCamelCase}}Entity)
	ctx.JSON(200, result)
}

//查询多条 如果不传page查询list
func Get{{ .StructName}}s(ctx *gin.Context) {
	result := entity.Result{}
	_ = ctx.Request.ParseForm()
	// 获取 sort 和 where
	sqlSort, sqlWhere := convert.GetSortAndWhereFormUrl(entity.{{ .StructName}}{}, ctx.Request.Form)
	query := {{.TableName}}_proto.Query{}
	query.SqlQuery = sqlWhere
	query.OrderBy = sqlSort
	// 获取 主从
	if ctx.Query("db_write") == "1" {
		query.Db = {{.TableName}}_proto.DB_WRITE
	}
	// 获取 page
	page, _ := strconv.Atoi(ctx.Query("page"))
	var isPageQuery bool //是否查询分页数据 如果不传page 认为查询list
	if page == 0 {
		page = 1
	} else {
		isPageQuery = true
	}
	// 获取 page_num
	pageNum, _ := strconv.Atoi(ctx.Query("page_num"))
	if pageNum == 0 {
		pageNum = 10
	}
	// 查询数据
	{{.StructName|toLowerCamelCase}}Service := {{.TableName}}_service.{{ .StructName}}Service{}
	pageData := entity.PageData{}
	if isPageQuery {
	pageData.List = {{.StructName|toLowerCamelCase}}Service.Get{{ .StructName}}PageList(query, page, pageNum)
	pageData.Count = {{.StructName|toLowerCamelCase}}Service.Get{{ .StructName}}Count(query)
	pageData.CurrentPage = page
	} else {
		pageData.List = {{.StructName|toLowerCamelCase}}Service.Get{{ .StructName}}List(query)
	}
	result.SetCode(entity.CODE_SUCCESS).SetMessage("success").SetData(pageData)
	ctx.JSON(200, result)

}

// 编辑单条
func Put{{ .StructName}}(ctx *gin.Context) {
	result := entity.Result{}
	_ = ctx.Request.ParseForm()

	// 获取参数
	primaryKey := ctx.Param("id")
	// 参数绑定
	{{.StructName|toLowerCamelCase}}Data := &entity.{{ .StructName}}{}
	err := ctx.ShouldBind({{.StructName|toLowerCamelCase}}Data)
	if err != nil {
		result.SetCode(entity.CODE_ERROR).SetMessage(err.Error())
		ctx.JSON(200, result)
		return
	}
	// 将 entity.{{ .StructName}} 转换成 {{.TableName}}_proto.{{ .StructName}}
	editData := {{.TableName}}_proto.{{ .StructName}}{}
	_ = convert.EntityToEntity({{.StructName|toLowerCamelCase}}Data, &editData)
	updateAndCondition := &{{.TableName}}_proto.UpdateAndCondition{}
	updateAndCondition.Query = &{{.TableName}}_proto.Query{}
	updateAndCondition.Query.SqlQuery = "{{.SqlPrimaryKey}} =" + primaryKey
	updateAndCondition.Entity = &editData
	updateAndCondition.UpdateEmptyFields = convert.GetEmptyEntityFieldFromPost(ctx.Request.PostForm,editData)
	entityUpdate := convert.EntityToMapWithEmpty(updateAndCondition.Entity, updateAndCondition.UpdateEmptyFields)
	// 编辑
	{{.StructName|toLowerCamelCase}}Service := {{.TableName}}_service.{{ .StructName}}Service{}
	err = {{.StructName|toLowerCamelCase}}Service.Update{{ .StructName}}(*updateAndCondition.Query, entityUpdate)
	if err != nil {
		result.SetCode(entity.CODE_ERROR).SetMessage(err.Error())
		ctx.JSON(200, result)
		return
	}
	result.SetCode(entity.CODE_SUCCESS).SetMessage("success")
	ctx.JSON(200, result)
}

// 删除单条
func Del{{ .StructName}}(ctx *gin.Context) {
	result := entity.Result{}
	// 获取参数
	primaryKey := ctx.Param("id")
	// 查询条件
	query := &{{.TableName}}_proto.Query{}
	query.SqlQuery = "{{.SqlPrimaryKey}} =" + primaryKey

	// 删除
	{{.StructName|toLowerCamelCase}}Service := {{.TableName}}_service.{{ .StructName}}Service{}
	err := {{.StructName|toLowerCamelCase}}Service.Delete{{ .StructName}}(*query)
	if err != nil {
		result.SetCode(entity.CODE_ERROR).SetMessage(err.Error())
		ctx.JSON(200, result)
		return
	}
	result.SetCode(entity.CODE_SUCCESS).SetMessage("success")
	ctx.JSON(200, result)
}


`
