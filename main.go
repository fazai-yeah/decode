package main

import (
	"decode/common"
	_ "decode/docs" // 这引用生成的docs包
	"decode/route"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	common.SetupLogger()
	r := gin.Default()

	r.POST("/upload", route.HandleSateFile)
	r.GET("/xml", route.HandleAlertFile)
	// 设置Swagger的访问路径，比如"/swagger/*any"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
