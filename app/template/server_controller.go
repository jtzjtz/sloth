package template

var ServerControllerTmpl = `package {{.TableName}}_controller

import (

	"context"
	"{{.ProjectName}}/entity"
	"{{.ProjectName}}/proto/{{.TableName}}_proto"
	"{{.ProjectName}}/server/service/{{.TableName}}_service"
	"github.com/jtzjtz/kit/convert"
	"strconv"
)

type {{ .StructName}}Controller struct {
}


func (instance *{{ .StructName}}Controller) Create{{ .StructName}}(ctx context.Context, req *{{.TableName}}_proto.{{ .StructName}}) (result *{{.TableName}}_proto.EntityResult, err error) {
	result=new({{.TableName}}_proto.EntityResult)
	defer func() {
		if err:= recover(); err != nil{
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
   {{.StructName|toLowerCamelCase}}Entity := entity.{{ .StructName}}{}
	err = convert.EntityToEntity(req, &{{.StructName|toLowerCamelCase}}Entity)
	if err != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = err.Error()
		return result, err
	}
	service := {{.TableName}}_service.{{ .StructName}}Service{}
	{{.StructName|toLowerCamelCase}}New, errCreate := service.Create{{ .StructName}}({{.StructName|toLowerCamelCase}}Entity)
	if errCreate != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = errCreate.Error()
		return result, errCreate
	}
	{{.StructName|toLowerCamelCase}}Out := {{.TableName}}_proto.{{ .StructName}}{}
	convert.EntityToEntity({{.StructName|toLowerCamelCase}}New, &{{.StructName|toLowerCamelCase}}Out)
	result.Data = &{{.StructName|toLowerCamelCase}}Out
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = "create success"
	return result, err
}
func (instance *{{ .StructName}}Controller) Update{{ .StructName}}(ctx context.Context, req *{{.TableName}}_proto.UpdateAndCondition) (result *{{.TableName}}_proto.Result, err error) {
	result=new({{.TableName}}_proto.Result)
	defer func() {
		if err:= recover(); err != nil{
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
    service := {{.TableName}}_service.{{ .StructName}}Service{}
	if req.Entity == nil {
		req.Entity = &{{.TableName}}_proto.{{ .StructName}}{}
	}
	entityUpdate := convert.EntityToMapWithEmpty(req.Entity, req.UpdateEmptyFields)
	if err = service.Update{{ .StructName}}(*req.Query, entityUpdate); err != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = err.Error()
		return result, err
	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = "create success"
	return result, err
}
func (instance *{{ .StructName}}Controller) Get{{ .StructName}}(ctx context.Context, req *{{.TableName}}_proto.Query) (result *{{.TableName}}_proto.EntityResult, err error) {
	result=new({{.TableName}}_proto.EntityResult)
    defer func() {
		if err:= recover(); err != nil{
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
    service := {{.TableName}}_service.{{ .StructName}}Service{}
	{{.StructName|toLowerCamelCase}}Entity := service.Get{{ .StructName}}(*req)
	{{.StructName|toLowerCamelCase}}Out := {{.TableName}}_proto.{{ .StructName}}{}
	if {{.StructName|toLowerCamelCase}}Entity ==nil {
		result.Data=nil
	}else {
		convert.EntityToEntity({{.StructName|toLowerCamelCase}}Entity, &{{.StructName|toLowerCamelCase}}Out)
		result.Data = &{{.StructName|toLowerCamelCase}}Out

	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	return result, nil

}
func (instance *{{ .StructName}}Controller) Get{{ .StructName}}List(ctx context.Context, req *{{.TableName}}_proto.Query) (result *{{.TableName}}_proto.ListResult, err error) {
	result=new({{.TableName}}_proto.ListResult)
    defer func() {
		if err:= recover(); err != nil{
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
    service := {{.TableName}}_service.{{ .StructName}}Service{}
	{{.StructName|toLowerCamelCase}}List := service.Get{{ .StructName}}List(*req)
	{{.StructName|toLowerCamelCase}}Out := []*{{.TableName}}_proto.{{ .StructName}}{}
	for _, {{.StructName|toLowerCamelCase}}Entity := range {{.StructName|toLowerCamelCase}}List {
		{{.StructName|toLowerCamelCase}}Temp := {{.TableName}}_proto.{{ .StructName}}{}
		convert.EntityToEntity({{.StructName|toLowerCamelCase}}Entity, &{{.StructName|toLowerCamelCase}}Temp)
		{{.StructName|toLowerCamelCase}}Out = append({{.StructName|toLowerCamelCase}}Out, &{{.StructName|toLowerCamelCase}}Temp)

	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	result.Data = {{.StructName|toLowerCamelCase}}Out
	return result, nil
}
func (instance *{{ .StructName}}Controller) Delete{{ .StructName}}(ctx context.Context, req *{{.TableName}}_proto.Query) (result *{{.TableName}}_proto.Result, err error) {
	result=new({{.TableName}}_proto.Result)
    defer func() {
		if err:= recover(); err != nil{
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
    service := {{.TableName}}_service.{{ .StructName}}Service{}
	if err = service.Delete{{ .StructName}}(*req); err != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = err.Error()
		return result, err
	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = "delete success"
	return result, err
}
func (instance *{{ .StructName}}Controller) Get{{ .StructName}}PageList(ctx context.Context, req *{{.TableName}}_proto.PageQuery) (result *{{.TableName}}_proto.PageResult, err error) {
	result=new({{.TableName}}_proto.PageResult)
    defer func() {
		if err:= recover(); err != nil{
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
    service := {{.TableName}}_service.{{ .StructName}}Service{}
	{{.StructName|toLowerCamelCase}}List := service.Get{{ .StructName}}PageList(*req.Query, int(req.Page), int(req.PageNum))
	{{.StructName|toLowerCamelCase}}Out := []*{{.TableName}}_proto.{{ .StructName}}{}
	{{.StructName|toLowerCamelCase}}PageData := {{.TableName}}_proto.{{ .StructName}}PageData{}
	for _, {{.StructName|toLowerCamelCase}}Entity := range {{.StructName|toLowerCamelCase}}List {
		{{.StructName|toLowerCamelCase}}Temp := {{.TableName}}_proto.{{ .StructName}}{}
		convert.EntityToEntity({{.StructName|toLowerCamelCase}}Entity, &{{.StructName|toLowerCamelCase}}Temp)
		{{.StructName|toLowerCamelCase}}Out = append({{.StructName|toLowerCamelCase}}Out, &{{.StructName|toLowerCamelCase}}Temp)

	}
	{{.StructName|toLowerCamelCase}}PageData.List = {{.StructName|toLowerCamelCase}}Out
	{{.StructName|toLowerCamelCase}}PageData.CurrentPage = req.Page
	{{.StructName|toLowerCamelCase}}PageData.Count = int32(service.Get{{ .StructName}}Count(*req.Query))
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	result.Data = &{{.StructName|toLowerCamelCase}}PageData
	return result, nil
}
func (instance *{{ .StructName}}Controller) Get{{ .StructName}}Count(ctx context.Context, req *{{.TableName}}_proto.Query) (result *{{.TableName}}_proto.Result, err error) {
	result=new({{.TableName}}_proto.Result)
    defer func() {
		if err:= recover(); err != nil{
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
    service := {{.TableName}}_service.{{ .StructName}}Service{}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	result.Data = strconv.Itoa(service.Get{{ .StructName}}Count(*req))
	return result, nil
}

`
