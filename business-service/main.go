package main

import (
	"business-service/handler"
	"business-service/service"
	"business-service/storage"
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 解析命令行参数
	port := flag.Int("p", 8080, "服务监听端口")
	flag.Parse()

	// 初始化存储
	personStorage := storage.NewMemoryStorage()

	// 初始化服务
	personService := service.NewPersonService(personStorage)

	// 初始化处理器
	frontendHandler := handler.NewFrontendHandler(personService)
	internalHandler := handler.NewInternalHandler(personService)

	// 创建路由
	router := gin.Default()

	// 注册路由
	frontendHandler.RegisterRoutes(router)
	internalHandler.RegisterRoutes(router)

	// 启动服务
	addr := fmt.Sprintf(":%d", *port)
	log.Printf("业务服务启动在 %s 端口\n", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
