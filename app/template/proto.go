package template

var ProtoTmpl = `syntax = "proto3";
option go_package ="./;{{.TableName}}_proto";
package {{.TableName}}_proto;
// protoc -I ./{{.TableName}}_proto --go_out=plugins=grpc:./{{.TableName}}_proto ./{{.TableName}}_proto/{{.TableName}}_gen.proto
service {{ .StructName}}Service {

    //创建
    rpc Create{{ .StructName}}({{ .StructName}}) returns (EntityResult) {
    }
    //1. 按照查询条件更新，查询可以按照实体查询和sql查询，二选一，sql查询优先使用，
    // 2.如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明，如果要更新的是默认值 请使用update_empty_fields 大写字段数组
    rpc Update{{ .StructName}} (UpdateAndCondition) returns (Result) {
    }
    //1. 按照条件查询当个实体，查询可以按照实体查询和sql查询，二选一，sql查询优先使用，
    // 2.如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明
    // 3.增加读写库查询判断
    rpc Get{{ .StructName}} (Query) returns (EntityResult) {
    }
    //1. 按照条件查询列表，查询可以按照实体查询和sql查询，二选一，sql查询优先使用，
    // 2.如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明
    // 3.增加读写库查询判断
    rpc Get{{ .StructName}}List (Query) returns (ListResult) {
    }
    //1. 按照条件删除，查询可以按照实体查询和sql查询，二选一，sql查询优先使用，
    // 2.如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明
    rpc Delete{{ .StructName}} (Query) returns (Result) {
    }
    //1. 按照条件查询，查询可以按照实体查询和sql查询，二选一，sql查询优先使用，
    // 2.如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明
    // 3.增加读写库查询判断
    rpc Get{{ .StructName}}PageList (PageQuery) returns (PageResult) {
    }
    //1. 按照条件查询数量，数量值为result.data，查询可以按照实体查询和sql查询，二选一，sql查询优先使用，
    // 2.如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明
    // 3.增加读写库查询判断
    rpc Get{{ .StructName}}Count (Query) returns (Result) {
    }
}

//主从库
enum DB {
    //读库
    READ = 0;
    //写库
    WRITE = 1;
}

message {{ .StructName}} {
    {{range .ProtoFields}}{{.}}
    {{end}}
}

message Result {
    int32 code = 1;
    string msg = 2;
    string data = 3;
}

// 按条件更新数据
message UpdateAndCondition {
    // 更新数据，如果filed对应为空值，则需要在 update_empty_fields 中声明
    {{ .StructName}} entity = 1;
    //查询条件
    Query query = 2;
    // 需要赋空值的字段 大写字段数组
    repeated string update_empty_fields = 3;

}

//按条件查询
message Query {
    // 按照实体查询 （和sql查询二选一，sql查询优先使用），如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明
    {{ .StructName}} entity_query = 1;
    // 按照sql查询（和实体查询二选一，sql查询优先使用），如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明
    // "id=1 and status in(3,4) and createtime >'2018' "
    string sql_query = 2;

    // 用空值当检索条件的字段 大写字段数组 
    repeated string query_empty_fields = 3;
    //排序条件 key值为数据库字段 ["id":"desc","create_time":asc"]
    map<string, string> order_by = 4;
    //是否查询主库，默认不读取 DB.READ
    DB db = 5;
    //制定select查询的字段，如 "id,username,age"
    string select_field=6;
}

//分页查询
message PageQuery {
    //查询条件
    Query query = 1;
    int32 page = 2;
    int32 page_num = 3;
}

message EntityResult {
    int32 code = 1;
    string msg = 2;
    {{ .StructName}} data = 3;
}

message ListResult {
    int32 code = 1;
    string msg = 2;
    repeated {{ .StructName}} data = 3;
}

message PageResult {
    int32 code = 1;
    string msg = 2;
    {{ .StructName}}PageData data = 3;
}

message {{ .StructName}}PageData {
    int32 current_page = 1;
    int32 count = 2;
    repeated {{ .StructName}} list = 3;
}

`
