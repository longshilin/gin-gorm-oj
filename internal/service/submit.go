package service

import (
	"getcharzp.cn/define"
	"getcharzp.cn/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetSubmitList
// @Tags 公共方法
// @Summary 提交列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param problem_identity query string false "problem_identity"
// @Param user_identity query string false "user_identity"
// @Param status query int false "status"
// @Success 200 {string} json "{"code":"200","data":"查询到的数据信息"}"
// @Router /submit-list [get]
func GetSubmitList(ctx *gin.Context) {
	// 分页相关
	page, err := strconv.Atoi(ctx.DefaultQuery("page", define.DefaultPage)) // 第几页
	if err != nil {
		log.Println("GetProblemList Page strconv Error: ", err)
	}
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", define.DefaultSize)) // 每页的纪录数
	page = (page - 1) * size                                              // 查询的offset开始序号
	// ^^

	var count int64 // 查询结果的总纪录条数
	list := make([]*models.SubmitBasic, 0)
	problemIdentity := ctx.Query("problem_identity")
	userIdentity := ctx.Query("user_identity")
	status, _ := strconv.Atoi(ctx.Query("status"))

	tx := models.GetSubmitList(problemIdentity, userIdentity, status)
	err = tx.Count(&count).Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("Get Problem List Error: ", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Submit List Error: " + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}
