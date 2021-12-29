package template

var ResultTmpl = `package entity

import (
    "encoding/json"
)


var (
	CODE_SUCCESS int = 1
	CODE_ERROR   int = -1
	CODE_UPDATE  int = 2
	CODE_NOAUTH  int = 401
)

const (
	RESULTSUCCESSINT32 int32 = 1
	RESULTERRORINT32   int32 = -1
	RESULTNOAUTHINT32  int32 = 401
	RESULTUPDATEINT32  int32 = 2

	RESULTSUCCESSINT int = 1
	RESULTERRORINT   int = -1
	RESULTNOAUTHINT  int = 401
	RESULTUPDATEINT  int = 2
)

type Result struct {
	Code int         {{.TransferChar}}json:"code"{{.TransferChar}}
	Msg  string      {{.TransferChar}}json:"msg"{{.TransferChar}}
	Data interface{} {{.TransferChar}}json:"data"{{.TransferChar}}
}

func (this *Result) GetCode() int {
	return this.Code
}
func (this *Result) SetCode(code int) *Result {
	this.Code = code
	return this
}
func (this *Result) GetMessage() string {
	return this.Msg
}
func (this *Result) SetMessage(message string) *Result {
	this.Msg = message
	return this
}
func (this *Result) GetData() interface{} {
	return this.Data
}
func (this *Result) SetData(data interface{}) *Result {
	this.Data = data
	return this
}
func (this *Result) ToJson() string {
	jsons, _ := json.Marshal(this)
	return string(jsons)
}
type PageData struct {
	CurrentPage int         {{.TransferChar}}json:"current_page"{{.TransferChar}}
	Count       int         {{.TransferChar}}json:"count"{{.TransferChar}}
	List        interface{} {{.TransferChar}}json:"list"{{.TransferChar}}
}

`
