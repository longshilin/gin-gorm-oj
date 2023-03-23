package router

import (
	_ "getcharzp.cn/docs"
	"getcharzp.cn/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	// Swagger 配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 配置路由规则
	r.GET("/ping", service.Ping)
	r.GET("/problem-list", service.GetProblemListTest)

	return r
}
