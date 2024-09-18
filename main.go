package main

import (
	"go_demo_test/router"
)

func main() {
	r := router.Router()

	// 启动HTTP服务，监听在8080端口
	r.Run(":8080")
}
