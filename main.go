package main

import (
	"fmt"
	"gee-Init/config/config_init"
	"gee-Init/docs"
	"gee-Init/router"
)

func main() {
	// 从配置文件读取配置
	err := config_init.Init()
	if err != nil {
		fmt.Printf("config_init config error : %s \n", err)
		return
	}
	docs.SwaggerInfo.BasePath = "/api/v1"
	// 装载路由
	r := router.NewRouter()
	r.Run(":3000")
}
