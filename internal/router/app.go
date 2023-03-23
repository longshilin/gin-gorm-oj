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

	// 问题列表
	r.GET("/problem-list", service.GetProblemList)
	// 问题详情
	r.GET("/problem-detail", service.GetProblemDetail)

	// 用户列表
	r.GET("/user-detail", service.GetUserDetail)
	// 提交纪录
	r.GET("/submit-list", service.GetSubmitList)
	// 分类列表
	r.GET("/category-list", service.GetCategoryList)

	return r
}
