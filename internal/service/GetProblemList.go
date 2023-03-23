package service

import (
	"getcharzp.cn/define"
	"getcharzp.cn/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetProblemListTest
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Param category_identity query string false "category_identity"
// @Success 200 {string} json "{"code":"200","data":"查询到的数据信息"}"
// @Router /problem-list [get]
func GetProblemListTest(ctx *gin.Context) {
	// 分页相关
	page, err := strconv.Atoi(ctx.DefaultQuery("page", define.DefaultPage)) // 第几页
	if err != nil {
		log.Println("GetProblemList Page strconv Error: ", err)
	}
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", define.DefaultSize)) // 每页的纪录数
	page = (page - 1) * size                                              // 查询的offset开始序号
	var count int64                                                       // 查询结果的总纪录条数
	keyword := ctx.Query("keyword")
	categoryIdentity := ctx.Query("category_identity")
	// ^^

	list := make([]*models.ProblemBasic, 0)
	tx := models.GetProblemList(keyword, categoryIdentity)
	err = tx.Count(&count).Omit("content").Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("Get Problem List Error: ", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})

}
