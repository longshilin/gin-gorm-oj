package main

import (
	"getcharzp.cn/router"
)

func main() {
	r := router.Router()
	// 启动gin.Engine服务，监听端口是10020
	r.Run("localhost:10020")
}
