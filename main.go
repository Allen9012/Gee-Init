package main

import (
	"fmt"
	"gee-Init/config"
	"gee-Init/router"
)

func main() {
	// 从配置文件读取配置
	err := config.Init()
	if err != nil {
		fmt.Printf("init config error : %s \n", err)
		return
	}

	// 装载路由
	r := router.NewRouter()
	r.Run(":3000")
}
