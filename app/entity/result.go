package entity

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
	RESULTORDERERROR int = 500
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (t Result) GetCode() int {
	return t.Code
}
func (t *Result) SetCode(code int) Result {
	t.Code = code
	return *t
}
func (t Result) GetMessage() string {
	return t.Msg
}
func (t *Result) SetMessage(message string) Result {
	t.Msg += message
	return *t
}
func (t Result) GetData() interface{} {
	return t.Data
}
func (t *Result) SetData(data interface{}) Result {
	t.Data = data
	return *t
}
func (t Result) ToJson() string {
	jsons, _ := json.Marshal(t)
	return string(jsons)
}
