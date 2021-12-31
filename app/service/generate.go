package service

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jimsmart/schema"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/inflection"
	_ "github.com/lib/pq"
	"github.com/serenize/snaker"
	"os"
	"os/exec"
	"path/filepath"
	"sloth/app/dbmeta"
	"sloth/app/entity"
	gtmpl "sloth/app/template"
	"strings"
	"text/template"
)

func Genrate(requestData entity.GenerateForm) entity.Result {
	var result entity.Result
	defer func() {
		if err := recover(); err != nil {
			result.SetMessage("get template error: " + err.(error).Error())
		}
	}()
	// Username is required
	sqlTable := requestData.Tabels
	sqlType := "mysql"
	var sqlConnStr string
	if requestData.DBPwd == "" || requestData.DBUserName == "" || requestData.DBHost == "" || requestData.DBPort == "" {
		return result.SetMessage("数据库，端口，用户名密码必填")
	} else {
		sqlConnStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?&parseTime=True", requestData.DBUserName, requestData.DBPwd, requestData.DBHost, requestData.DBPort, requestData.DB)
	}
	if sqlConnStr == "" {
		fmt.Println("sql connection string is required! Add it with --connstr=s")
		return result.SetMessage("sql connection string is required! Add it with --connstr=s")
	}

	var db, err = sql.Open(sqlType, sqlConnStr)
	if err != nil {
		fmt.Println("Error in open database: " + err.Error())
		return result.SetMessage("Error in open database: " + err.Error())

	}
	defer db.Close()

	// parse or read tables
	var tables []string
	if sqlTable != "" {
		tables = strings.Split(sqlTable, ",")
	} else {
		tables, err = schema.TableNames(db)
		if err != nil {
			fmt.Println("Error in fetching tables information from mysql information schema")
			return result.SetMessage("Error in fetching tables information from mysql information schema")
		}
	}

	projectDir := "./gen_files/" + requestData.GitUrl
	rootDir := projectDir + "/" + requestData.Tag
	protoDir := "proto"
	apiDir := "client"
	serviceDir := "server"
	entityDir := "entity"
	errDir := os.MkdirAll(rootDir, 0777)
	if errDir != nil {
		return result.SetMessage("创建根目录失败：" + errDir.Error())
	}
	pullCmd := fmt.Sprintf("./git_pull.sh %s %s %s %s %s", projectDir, requestData.GitUrl, requestData.Tag, requestData.GitUser, requestData.GitPwd)

	er, outStr, erStr := Shellout(pullCmd)
	println(er, outStr, erStr)

	if er != nil {
		return result.SetMessage(erStr + er.Error())
	}
	result.SetMessage(outStr)

	os.MkdirAll(filepath.Join(rootDir, protoDir), 0777)
	os.MkdirAll(filepath.Join(rootDir, apiDir, "grpc"), 0777)
	os.MkdirAll(filepath.Join(rootDir, serviceDir, "controller"), 0777)
	//os.MkdirAll(filepath.Join(rootDir, serviceDir, "entity"), 0777)
	os.MkdirAll(filepath.Join(rootDir, serviceDir, "dao"), 0777)
	os.MkdirAll(filepath.Join(rootDir, serviceDir, "service"), 0777)
	os.MkdirAll(filepath.Join(rootDir, entityDir), 0777)

	modelT := getTemplate(gtmpl.ModelTmpl)

	serverControllerT := getTemplate(gtmpl.ServerControllerTmpl)

	serverHttpControllerT := getTemplate(gtmpl.ServerHttpControllerTmpl)

	serverServiceT := getTemplate(gtmpl.ServerServiceTmpl)

	clientServiceT := getTemplate(gtmpl.ClientServiceTmpl)

	dalT := getTemplate(gtmpl.DalTmpl)
	protoT := getTemplate(gtmpl.ProtoTmpl)

	resultT := getTemplate(gtmpl.ResultTmpl)

	sqldbT := getTemplate(gtmpl.SqlDBTmpl)

	defer func() {
		errGen := recover()
		if errGen != nil {
			result.SetCode(entity.CODE_ERROR)
			result.SetMessage(fmt.Sprintf("genfile errr: %s", errGen))
		}
	}()
	genFile(filepath.Join(rootDir, "entity", "result.go"), dbmeta.ModelInfo{TransferChar: "`"}, resultT, "", true)
	genFile(filepath.Join(rootDir, serviceDir, "dao", "sqldb.go"), nil, sqldbT, "", true)
	// generate go files for each table
	for _, tableName := range tables {
		structName := dbmeta.FmtFieldName(tableName)
		structName = inflection.Singular(structName)

		model := dbmeta.CreateModel(db, requestData.DB, tableName, structName, requestData.GitUrl)
		genFile(filepath.Join(rootDir, "entity", tableName+".go"), model, modelT, "", true)
		genFile(filepath.Join(rootDir, serviceDir, "controller", tableName+"_controller", tableName+"_controller_gen.go"), model, serverControllerT, filepath.Join(rootDir, serviceDir, "controller", tableName)+"_controller", true) //server_controller

		//server_http_controller begin
		if requestData.IsHttpService == 1 {
			genFile(filepath.Join(rootDir, serviceDir, "controller", tableName+"_controller", tableName+"_http_controller_gen.go"), model, serverHttpControllerT, filepath.Join(rootDir, serviceDir, "controller", tableName)+"_controller", true) //server_http_controller
		}
		//server_http_controller end

		genFile(filepath.Join(rootDir, serviceDir, "service", tableName+"_service", tableName+"_service_gen.go"), model, serverServiceT, filepath.Join(rootDir, serviceDir, "service", tableName)+"_service", true) //serverService

		genFile(filepath.Join(rootDir, apiDir, "grpc", tableName+"_grpc", tableName+"_service_gen.go"), model, clientServiceT, filepath.Join(rootDir, apiDir, "grpc", tableName)+"_grpc", true) //client Service
		genFile(filepath.Join(rootDir, serviceDir, "dao", tableName+"_dao_gen.go"), model, dalT, "", true)                                                                                      //dal
		genFile(filepath.Join(rootDir, protoDir, tableName+"_proto", tableName+"_gen.proto"), model, protoT, filepath.Join(rootDir, protoDir, tableName)+"_proto", false)                       //proto
		exeStr := "protoc -I ./" + rootDir + "/" + protoDir + "/" + tableName + "_proto --go_out=plugins=grpc:./" + rootDir + "/" + protoDir + "/" + tableName + "_proto ./" + rootDir + "/" + protoDir + "/" + tableName + "_proto/" + tableName + "_gen.proto"
		shellerr, errStr1, errStr2 := Shellout(exeStr)
		if shellerr != nil {
			fmt.Printf("generate pb error: %v  %s %s\n", shellerr.Error(), errStr1, errStr2)
			return result.SetMessage(errStr1 + errStr2)

		}
		println(exeStr)

	}
	println("代码自动生成结束")
	pushCmd := fmt.Sprintf("./git_push.sh %s \"%s\" %s %s %s %s", rootDir, requestData.GitMsg, requestData.Tag, requestData.GitUrl, requestData.GitUser, requestData.GitPwd)

	er, outStr, erStr = Shellout(pushCmd)
	println(er, outStr, erStr)
	if er != nil {
		return result.SetMessage(erStr)
	}
	result.SetMessage(outStr)
	result.SetCode(entity.CODE_SUCCESS)
	return result
}

//获取模板 成功返回 templage 失败抛出err（请补获）
func getTemplate(t string) *template.Template {
	var funcMap = template.FuncMap{
		"pluralize":        inflection.Plural,
		"title":            strings.Title,
		"toLower":          strings.ToLower,
		"toLowerCamelCase": camelToLowerCamel,
		"toSnakeCase":      snaker.CamelToSnake,
	}

	tmpl, err := template.New("model").Funcs(funcMap).Parse(t)
	if err != nil {
		panic(errors.New(fmt.Sprintf("%s template load err:%s", t, err.Error())))
	}
	return tmpl
}

func camelToLowerCamel(s string) string {
	ss := strings.Split(s, "")
	ss[0] = strings.ToLower(ss[0])

	return strings.Join(ss, "")
}

func Shellout(command string) (error, string, string) {
	return LinuxShell(command)
}

const ShellToUse = "bash"

func LinuxShell(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func WindowsShell(command string) (error, string, string) {
	c := exec.Command("cmd", "/C", command)
	err := c.Run()
	return err, "", ""
}
