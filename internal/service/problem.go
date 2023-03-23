package service

import (
	"getcharzp.cn/define"
	"getcharzp.cn/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

// GetProblemList
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Param category_identity query string false "category_identity"
// @Success 200 {string} json "{"code":"200","data":"查询到的数据信息"}"
// @Router /problem-list [get]
func GetProblemList(ctx *gin.Context) {
	// 分页相关
	page, err := strconv.Atoi(ctx.DefaultQuery("page", define.DefaultPage)) // 第几页
	if err != nil {
		log.Println("GetProblemList Page strconv Error: ", err)
	}
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", define.DefaultSize)) // 每页的纪录数
	page = (page - 1) * size                                              // 查询的offset开始序号
	// ^^

	var count int64 // 查询结果的总纪录条数
	keyword := ctx.Query("keyword")
	categoryIdentity := ctx.Query("category_identity")
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

// GetProblemDetail
// @Tags 公共方法
// @Summary 问题详情
// @Param identity query string false "problem identity"
// @Success 200 {string} json "{"code":"200","data":"查询到的数据信息"}"
// @Router /problem-detail [get]
func GetProblemDetail(ctx *gin.Context) {
	identity := ctx.Query("identity")
	if identity == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "问题唯一标识不能为空",
		})
		return
	}

	data := new(models.ProblemBasic)
	err := models.DB.Where("identity = ?", identity).
		Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusNotFound,
				"msg":  "未找到该问题",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "服务器内部错误: " + err.Error(),
			})
		}
		return

	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}
