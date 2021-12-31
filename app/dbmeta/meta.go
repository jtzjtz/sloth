package dbmeta

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/jimsmart/schema"
)

type ModelInfo struct {
	PackageName     string
	StructName      string
	ShortStructName string
	TableName       string
	Fields          []string
	Columns         []*sql.ColumnType
	ProtoFields     []string
	SqlPrimaryKey   string
	PrimaryKey      string
	PrimaryType     interface{}
	ProjectName     string
	//ProtoName       string
	//GitUrl       string
	Tables       []string
	TransferChar string // 转义字符`
	//ServiceName string // service 项目名称

}

// commonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
var commonInitialisms = map[string]bool{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SSH":   true,
	"TLS":   true,
	"TTL":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
}

var intToWordMap = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

// Constants for return types of golang
const (
	golangByteArray  = "[]byte"
	gureguNullInt    = "int"
	sqlNullInt       = "int"
	golangInt        = "int"
	golangInt64      = "int"
	gureguNullFloat  = "float32"
	sqlNullFloat     = "float32"
	golangFloat      = "float32"
	golangFloat32    = "float32"
	golangFloat64    = "float64"
	gureguNullString = "string"
	sqlNullString    = "string"
	gureguNullTime   = "string"
	golangTime       = "string"

	protoInt    = "int32"
	protoFloat  = "float"
	protoString = "string"
)

type Field struct {
	FieldName string
	FieldDesc string
	DataType  string
	IsNull    string
	Length    int
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getFiledInfo(db *sql.DB, dbName string, tableName string) map[string]Field {
	sqlStr := `SELECT COLUMN_NAME fName,column_comment fDesc,DATA_TYPE dataType,
						IS_NULLABLE isNull,IFNULL(CHARACTER_MAXIMUM_LENGTH,0) sLength
			FROM information_schema.columns 
			WHERE table_schema = ? AND table_name = ?`

	data := map[string]Field{}

	rows, err := db.Query(sqlStr, dbName, tableName)
	checkErr(err)

	for rows.Next() {
		var f Field
		err = rows.Scan(&f.FieldName, &f.FieldDesc, &f.DataType, &f.IsNull, &f.Length)
		checkErr(err)

		data[f.FieldName] = f
	}
	return data
}

// GenerateStruct generates a struct for the given table.
func GenerateStruct(db *sql.DB, tableName string, structName string, pkgName string, jsonAnnotation bool, gormAnnotation bool, gureguTypes bool) *ModelInfo {
	cols, _ := schema.Table(db, tableName)
	fields := generateFieldsTypes(db, cols, 0, jsonAnnotation, gormAnnotation, gureguTypes)
	protoFields := generateProtoFieldsTypes(db, cols, gureguTypes)

	//fields := generateMysqlTypes(db, columnTypes, 0, jsonAnnotation, gormAnnotation, gureguTypes)
	nullable, _ := cols[0].Nullable()
	var modelInfo = &ModelInfo{
		PackageName:     pkgName,
		StructName:      structName,
		TableName:       tableName,
		ShortStructName: strings.ToLower(string(structName[0])),
		Fields:          fields,
		Columns:         cols,
		ProtoFields:     protoFields,
		SqlPrimaryKey:   cols[0].Name(),
		PrimaryKey:      FmtFieldName(stringifyFirstChar(cols[0].Name())),
		PrimaryType:     sqlTypeToGoType(strings.ToLower(cols[0].DatabaseTypeName()), nullable, gureguTypes),
	}

	return modelInfo
}

// Generate proto fields string
func generateProtoFieldsTypes(db *sql.DB, columns []*sql.ColumnType, gureguTypes bool) []string {
	var fields []string
	var field = ""
	for i, c := range columns {
		key := c.Name()
		valueType := sqlTypeToProtoType(strings.ToLower(c.DatabaseTypeName()))
		if valueType == "" { // unknown type
			continue
		}
		fieldName := key

		//@inject_tag: valid:"ip"
		//field = fmt.Sprintf("//@inject_tag: form:\"%s\"",
		//	fieldName,
		//)
		//fields = append(fields, field)

		//fieldName := FmtFieldName(stringifyFirstChar(key))
		field = fmt.Sprintf("%s %s = %s;",
			valueType,
			fieldName,
			strconv.Itoa(i+1))

		fields = append(fields, field)
	}
	return fields
}

// Generate fields string
func generateFieldsTypes(db *sql.DB, columns []*sql.ColumnType, depth int, jsonAnnotation bool, gormAnnotation bool, gureguTypes bool) []string {

	//sort.Strings(keys)

	var fields []string
	var field = ""
	for i, c := range columns {
		nullable, _ := c.Nullable()
		key := c.Name()
		valueType := sqlTypeToGoType(strings.ToLower(c.DatabaseTypeName()), nullable, gureguTypes)
		if valueType == "" { // unknown type
			continue
		}
		fieldName := FmtFieldName(stringifyFirstChar(key))

		var annotations []string
		if gormAnnotation == true {
			if i == 0 {
				annotations = append(annotations, fmt.Sprintf("gorm:\"column:%s;primary_key\"", key))
			} else {
				annotations = append(annotations, fmt.Sprintf("gorm:\"column:%s\"", key))
			}

		}
		if jsonAnnotation == true {
			annotations = append(annotations, fmt.Sprintf("json:\"%s\" form:\"%s\"", key, key))
		}
		if len(annotations) > 0 {
			field = fmt.Sprintf("%s %s `%s`",
				fieldName,
				valueType,
				strings.Join(annotations, " "))

		} else {
			field = fmt.Sprintf("%s %s",
				fieldName,
				valueType)
		}

		fields = append(fields, field)
	}
	return fields
}
func GenerateEnttyFields(columns []*sql.ColumnType, db *sql.DB, dbName string, tableName string) []string {

	filedsInfo := getFiledInfo(db, dbName, tableName)
	var fields []string
	var field = ""
	for i, c := range columns {
		nullable, _ := c.Nullable()
		key := c.Name()
		valueType := sqlTypeToGoType(strings.ToLower(c.DatabaseTypeName()), nullable, true)
		if valueType == "" { // unknown type
			continue
		}
		fieldName := FmtFieldName(stringifyFirstChar(key))

		var annotations []string
		if i == 0 {
			annotations = append(annotations, fmt.Sprintf("gorm:\"column:%s;primary_key\"", key))
		} else {
			annotations = append(annotations, fmt.Sprintf("gorm:\"column:%s\"", key))
		}

		annotations = append(annotations, fmt.Sprintf("json:\"%s\" form:\"%s\"", key, key))

		if len(annotations) > 0 {
			field = fmt.Sprintf("%s %s `%s` // %s",
				fieldName,
				valueType,
				strings.Join(annotations, " "), filedsInfo[key].FieldDesc)

		} else {
			field = fmt.Sprintf("%s %s",
				fieldName,
				valueType)
		}

		fields = append(fields, field)
	}
	return fields
}
func sqlTypeToProtoType(mysqlType string) string {
	switch mysqlType {
	case "tinyint", "int", "smallint", "mediumint":
		return protoInt
	case "bigint":
		return protoInt
	case "char", "enum", "varchar", "longtext", "mediumtext", "text", "tinytext":
		return protoString
	case "date", "datetime", "time", "timestamp":
		return protoString
	case "decimal", "double":
		return protoFloat
	case "float":
		return protoFloat
	case "binary", "blob", "longblob", "mediumblob", "varbinary":
		return protoString
	}

	return ""
}

func sqlTypeToGoType(mysqlType string, nullable bool, gureguTypes bool) string {
	switch mysqlType {
	case "tinyint", "int", "smallint", "mediumint":
		if nullable {
			if gureguTypes {
				return gureguNullInt
			}
			return sqlNullInt
		}
		return golangInt
	case "bigint":
		if nullable {
			if gureguTypes {
				return gureguNullInt
			}
			return sqlNullInt
		}
		return golangInt64
	case "char", "enum", "varchar", "longtext", "mediumtext", "text", "tinytext":
		if nullable {
			if gureguTypes {
				return gureguNullString
			}
			return sqlNullString
		}
		return "string"
	case "date", "datetime", "time", "timestamp":
		if nullable && gureguTypes {
			return gureguNullTime
		}
		return golangTime
	case "decimal", "double":
		if nullable {
			if gureguTypes {
				return gureguNullFloat
			}
			return sqlNullFloat
		}
		return golangFloat64
	case "float":
		if nullable {
			if gureguTypes {
				return gureguNullFloat
			}
			return sqlNullFloat
		}
		return golangFloat32
	case "binary", "blob", "longblob", "mediumblob", "varbinary":
		return golangByteArray
	}
	return ""
}
func CreateModel(db *sql.DB, dbname string, tableName string, structName string, giturl string) *ModelInfo {
	defer func() {
		if er := recover(); er != nil {
			panic(er)
		}
	}()
	cols, _ := schema.Table(db, tableName)
	fields := GenerateEnttyFields(cols, db, dbname, tableName)
	protoFields := generateProtoFieldsTypes(db, cols, true)

	//fields := generateMysqlTypes(db, columnTypes, 0, jsonAnnotation, gormAnnotation, gureguTypes)
	nullable, _ := cols[0].Nullable()
	var modelInfo = &ModelInfo{
		PackageName:     "",
		StructName:      structName,
		TableName:       tableName,
		ShortStructName: strings.ToLower(string(structName[0])),
		Fields:          fields,
		Columns:         cols,
		ProtoFields:     protoFields,
		SqlPrimaryKey:   cols[0].Name(),
		PrimaryKey:      FmtFieldName(stringifyFirstChar(cols[0].Name())),
		PrimaryType:     sqlTypeToGoType(strings.ToLower(cols[0].DatabaseTypeName()), nullable, true),
		TransferChar:    "`",
		ProjectName:     giturl,
	}

	return modelInfo
}
