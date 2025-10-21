package main

import (
	"flag"
	"fmt"
	"log"
	"mcp-service/client"
	"mcp-service/server"
	"mcp-service/tools"
	"os"
)

func main() {
	// 解析命令行参数
	port := flag.Int("p", 8081, "服务监听端口")
	flag.Parse()

	// 获取业务服务地址
	businessServiceURL := os.Getenv("BUSINESS_SERVICE_URL")
	if businessServiceURL == "" {
		businessServiceURL = "http://localhost:8080"
	}

	// 创建业务服务客户端
	businessClient := client.NewBusinessClient(businessServiceURL)

	// 创建MCP工具
	personTools := tools.NewPersonTools(businessClient)

	// 创建MCP服务器
	mcpServer := server.NewMCPServer(personTools)

	// 打印工具信息
	mcpServer.PrintToolsInfo()

	// 启动服务器
	addr := fmt.Sprintf(":%d", *port)
	log.Printf("MCP服务启动在 %s 端口\n", addr)
	if err := mcpServer.Run(addr); err != nil {
		log.Fatal("MCP服务启动失败:", err)
	}
}
