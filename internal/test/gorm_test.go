package test

import (
	"fmt"
	"getcharzp.cn/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGormTest(t *testing.T) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:eden123@tcp(172.27.89.219:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	data := make([]*models.ProblemBasic, 0)
	err = db.Find(&data).Error

	if err != nil {
		t.Fatal(err)
	}

	for _, v := range data {
		fmt.Printf("Problem ===> %v \n", v)
	}
}
