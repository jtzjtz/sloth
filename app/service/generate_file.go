package service

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
	"sloth/app/entity"
	"text/template"
)

//生成文件，成功返回result，报错抛出string错误（请补获）
func genFile(filePath string, data interface{}, template *template.Template, dir string, isFormat bool) (result entity.Result) {
	var buf bytes.Buffer
	err := template.Execute(&buf, data)
	if err != nil {
		panic("Error in rendering model: " + err.Error())

	}
	fileDataStr := buf.Bytes()
	var errFormat error
	if isFormat {
		if fileDataStr, errFormat = format.Source(buf.Bytes()); errFormat != nil {
			panic("Error in formating source: " + errFormat.Error())
		}
	}

	if dir != "" {
		errDir := os.MkdirAll(dir, 0777)
		if errDir != nil {
			panic(dir + "文件夹创建失败：" + errDir.Error())

		}
	}

	errFile := ioutil.WriteFile(filePath, fileDataStr, 0777)
	if errFile != nil {
		panic(filePath + "文件创建失败：" + errFile.Error())

	}
	result.SetCode(entity.CODE_SUCCESS)
	return result
}
