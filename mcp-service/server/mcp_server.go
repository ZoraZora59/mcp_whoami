package server

import (
	"encoding/json"
	"log"
	"mcp-service/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MCPServer MCP HTTP服务器
type MCPServer struct {
	personTools *tools.PersonTools
	router      *gin.Engine
}

// NewMCPServer 创建MCP服务器实例
func NewMCPServer(personTools *tools.PersonTools) *MCPServer {
	server := &MCPServer{
		personTools: personTools,
		router:      gin.Default(),
	}
	server.setupRoutes()
	return server
}

// setupRoutes 设置路由
func (s *MCPServer) setupRoutes() {
	// MCP协议端点
	s.router.POST("/mcp/tools/list", s.handleToolsList)
	s.router.POST("/mcp/tools/call", s.handleToolsCall)
	s.router.POST("/mcp/initialize", s.handleInitialize)

	// 健康检查
	s.router.GET("/health", s.handleHealth)
}

// handleInitialize 处理初始化请求
func (s *MCPServer) handleInitialize(c *gin.Context) {
	var request map[string]interface{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"protocolVersion": "2024-11-05",
		"capabilities": map[string]interface{}{
			"tools": map[string]interface{}{},
		},
		"serverInfo": map[string]interface{}{
			"name":    "mcp-whoami-server",
			"version": "1.0.0",
		},
	}

	c.JSON(http.StatusOK, response)
}

// handleToolsList 处理工具列表请求
func (s *MCPServer) handleToolsList(c *gin.Context) {
	toolsList := s.personTools.GetToolsList()

	response := map[string]interface{}{
		"tools": toolsList,
	}

	c.JSON(http.StatusOK, response)
}

// handleToolsCall 处理工具调用请求
func (s *MCPServer) handleToolsCall(c *gin.Context) {
	var request struct {
		Params struct {
			Name      string                 `json:"name"`
			Arguments map[string]interface{} `json:"arguments"`
		} `json:"params"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("解析请求失败: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("调用工具: %s, 参数: %v", request.Params.Name, request.Params.Arguments)

	// 调用工具
	result, err := s.personTools.CallTool(request.Params.Name, request.Params.Arguments)

	// 格式化响应
	response := s.personTools.FormatResult(result, err)

	c.JSON(http.StatusOK, response)
}

// handleHealth 健康检查
func (s *MCPServer) handleHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// Run 启动服务器
func (s *MCPServer) Run(addr string) error {
	log.Printf("MCP服务启动在 %s", addr)
	return s.router.Run(addr)
}

// PrintToolsInfo 打印工具信息（用于调试）
func (s *MCPServer) PrintToolsInfo() {
	tools := s.personTools.GetToolsList()
	toolsJSON, _ := json.MarshalIndent(tools, "", "  ")
	log.Printf("可用工具:\n%s", string(toolsJSON))
}
