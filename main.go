package main

import (
	"github.com/gin-gonic/gin"
	"sloth/app/route"
)

//启动服务
func main() {
	var port = ":8000"
	var engin *gin.Engine
	gin.SetMode(gin.DebugMode)
	engin = gin.New()

	route.Init(engin) //路由初始化
	_ = engin.Run(port)
}
