package template

var ClientServiceTmpl = `package {{.TableName}}_grpc

import (
	"{{.ProjectName}}/entity"
	"{{.ProjectName}}/proto/{{.TableName}}_proto"
	"context"
	"google.golang.org/grpc"
)

type {{ .StructName}}Service struct {
}

func (t *{{ .StructName}}Service) Create{{ .StructName}}(ctx context.Context, {{.StructName|toLowerCamelCase}}Entity *{{.TableName}}_proto.{{ .StructName}}, conn *grpc.ClientConn) (result *{{.TableName}}_proto.EntityResult) {
	result = &{{.TableName}}_proto.EntityResult{Code: entity.RESULTERRORINT32}
	if {{.StructName|toLowerCamelCase}}Entity == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件参数错误"
	} else {
		client := {{.TableName}}_proto.New{{ .StructName}}ServiceClient(conn)
		resultReponse, err := client.Create{{ .StructName}}(ctx, {{.StructName|toLowerCamelCase}}Entity)
		if err != nil || resultReponse == nil {
			result.Msg = err.Error()
			result.Code = entity.RESULTERRORINT32
		} else {
			result.Code = resultReponse.GetCode()
			result.Msg = resultReponse.GetMsg()
			result.Data = resultReponse.GetData()
		}
	}
	return result	
}

func (t *{{ .StructName}}Service) Update{{ .StructName}}(ctx context.Context, orderUpdateAndCondition *{{.TableName}}_proto.UpdateAndCondition, conn *grpc.ClientConn) (result *{{.TableName}}_proto.Result) {
	result = &{{.TableName}}_proto.Result{Code: entity.RESULTERRORINT32}
	if orderUpdateAndCondition.Entity == nil || orderUpdateAndCondition.Query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if orderUpdateAndCondition.Query.EntityQuery == nil && orderUpdateAndCondition.Query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := {{.TableName}}_proto.New{{ .StructName}}ServiceClient(conn)
		resultReponse, err := client.Update{{ .StructName}}(ctx, orderUpdateAndCondition)
		if err != nil || resultReponse == nil {
			result.Msg = err.Error()
			result.Code = entity.RESULTERRORINT32
		} else {
			result.Code = resultReponse.GetCode()
			result.Msg = resultReponse.GetMsg()
			result.Data = resultReponse.GetData()
		}
	}
	return result
}

func (t *{{ .StructName}}Service) Get{{ .StructName}}(ctx context.Context, query *{{.TableName}}_proto.Query, conn *grpc.ClientConn) (result *{{.TableName}}_proto.EntityResult) {
	result = &{{.TableName}}_proto.EntityResult{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := {{.TableName}}_proto.New{{ .StructName}}ServiceClient(conn)
		resultReponse, err := client.Get{{ .StructName}}(ctx, query)
		if err != nil || resultReponse == nil {
			result.Msg = err.Error()
			result.Code = entity.RESULTERRORINT32
		} else {
			result.Code = resultReponse.GetCode()
			result.Msg = resultReponse.GetMsg()
			result.Data = resultReponse.GetData()
		}
	}
	return result
}

func (t *{{ .StructName}}Service) Get{{ .StructName}}List(ctx context.Context, query *{{.TableName}}_proto.Query, conn *grpc.ClientConn) (result *{{.TableName}}_proto.ListResult) {
	result = &{{.TableName}}_proto.ListResult{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := {{.TableName}}_proto.New{{ .StructName}}ServiceClient(conn)
		resultReponse, err := client.Get{{ .StructName}}List(ctx, query)
		if err != nil || resultReponse == nil {
			result.Msg = err.Error()
			result.Code = entity.RESULTERRORINT32
		} else {
			result.Code = resultReponse.GetCode()
			result.Msg = resultReponse.GetMsg()
			result.Data = resultReponse.GetData()
		}
	}
	return result
}

func (this *{{ .StructName}}Service) Delete{{ .StructName}}(ctx context.Context, query *{{.TableName}}_proto.Query, conn *grpc.ClientConn) (result *{{.TableName}}_proto.Result) {
	result = &{{.TableName}}_proto.Result{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := {{.TableName}}_proto.New{{ .StructName}}ServiceClient(conn)
		resultReponse, err := client.Delete{{ .StructName}}(ctx, query)
		if err != nil || resultReponse == nil {
			result.Msg = err.Error()
			result.Code = entity.RESULTERRORINT32
		} else {
			result.Code = resultReponse.GetCode()
			result.Msg = resultReponse.GetMsg()
			result.Data = resultReponse.GetData()
		}
	}
	return result
}

func (t *{{ .StructName}}Service) Get{{ .StructName}}PageList(ctx context.Context, pageQuery *{{.TableName}}_proto.PageQuery, conn *grpc.ClientConn) (result *{{.TableName}}_proto.PageResult) {
	result = &{{.TableName}}_proto.PageResult{Code: entity.RESULTERRORINT32}
	if pageQuery == nil || pageQuery.Query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if pageQuery.Query.EntityQuery == nil && pageQuery.Query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"

	} else {
		client := {{.TableName}}_proto.New{{ .StructName}}ServiceClient(conn)
		resultReponse, err := client.Get{{ .StructName}}PageList(ctx, pageQuery)
		if err != nil || resultReponse == nil {
			result.Msg = err.Error()
			result.Code = entity.RESULTERRORINT32
		} else {
			result.Code = resultReponse.GetCode()
			result.Msg = resultReponse.GetMsg()
			result.Data = resultReponse.GetData()
		}
	}
	return result
}

func (t *{{ .StructName}}Service) Get{{ .StructName}}Count(ctx context.Context, query *{{.TableName}}_proto.Query, conn *grpc.ClientConn) (result *{{.TableName}}_proto.Result) {
	result = &{{.TableName}}_proto.Result{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"

	} else {
		client := {{.TableName}}_proto.New{{ .StructName}}ServiceClient(conn)
		resultReponse, err := client.Get{{ .StructName}}Count(ctx, query)
		if err != nil || resultReponse == nil {
			result.Msg = err.Error()
			result.Code = entity.RESULTERRORINT32
		} else {
			result.Code = resultReponse.GetCode()
			result.Msg = resultReponse.GetMsg()
			result.Data = resultReponse.GetData()
		}
	}
	return result
}


`
