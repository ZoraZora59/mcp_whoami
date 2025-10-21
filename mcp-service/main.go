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
	port := flag.Int("p", 8081, "服务监听端口（HTTP模式）")
	businessServicePort := flag.Int("b", 8080, "业务服务监听端口")
	stdio := flag.Bool("stdio", false, "使用stdio模式（用于Cursor等客户端）")
	flag.Parse()

	// 获取业务服务地址
	businessServiceURL := os.Getenv("BUSINESS_SERVICE_URL")
	if businessServiceURL == "" {
		businessServiceURL = fmt.Sprintf("http://localhost:%d", *businessServicePort)
	}

	// 创建业务服务客户端
	businessClient := client.NewBusinessClient(businessServiceURL)

	// 创建MCP工具
	personTools := tools.NewPersonTools(businessClient)

	// 根据模式启动不同的服务器
	if *stdio {
		// stdio模式（用于Cursor）
		stdioServer := server.NewStdioServer(personTools)
		if err := stdioServer.Run(); err != nil {
			log.Fatal("stdio服务启动失败:", err)
		}
	} else {
		// HTTP模式（用于测试和调试）
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
}
