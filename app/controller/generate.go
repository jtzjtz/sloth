package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sloth/app/entity"
	"sloth/app/service"
)

func Create(ctx *gin.Context) {

	var result entity.Result
	var form entity.GenerateForm
	er := ctx.ShouldBind(&form)
	if er != nil {
		result.Msg = er.Error()
		ctx.JSON(200, result)
		return
	}
	fmt.Printf("执行人：%s  项目：%s version:%s \n", form.GitUser, form.GitUrl, form.Tag)
	result = service.Genrate(form)
	ctx.JSON(200, result)
	return

}
func Index(ctx *gin.Context) {

	ctx.HTML(200, "index.html", nil)
}
