package main

import (
	"harbor-image-delete/route"
)

func main() {
	// 初始化路由
	r := route.InitRouter()
	r.Run(":8090")

}
