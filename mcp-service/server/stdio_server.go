package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"mcp-service/tools"
	"os"
)

// StdioServer MCP stdio模式服务器
type StdioServer struct {
	personTools *tools.PersonTools
}

// NewStdioServer 创建stdio模式服务器实例
func NewStdioServer(personTools *tools.PersonTools) *StdioServer {
	return &StdioServer{
		personTools: personTools,
	}
}

// Run 启动stdio服务器
func (s *StdioServer) Run() error {
	log.SetOutput(os.Stderr) // 日志输出到stderr，避免干扰stdio通信
	log.Println("MCP stdio服务器已启动")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// 解析JSON-RPC请求
		var request map[string]interface{}
		if err := json.Unmarshal([]byte(line), &request); err != nil {
			log.Printf("解析请求失败: %v, 原始内容: %s", err, line)
			s.sendError(-32700, "Parse error", nil)
			continue
		}

		// 处理请求
		s.handleRequest(request)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("读取输入失败: %v", err)
	}

	return nil
}

// handleRequest 处理JSON-RPC请求
func (s *StdioServer) handleRequest(request map[string]interface{}) {
	method, ok := request["method"].(string)
	if !ok {
		s.sendError(-32600, "Invalid Request", request["id"])
		return
	}

	id := request["id"]

	switch method {
	case "initialize":
		s.handleInitialize(id, request)
	case "tools/list":
		s.handleToolsList(id)
	case "tools/call":
		s.handleToolsCall(id, request)
	case "ping":
		s.handlePing(id)
	default:
		s.sendError(-32601, "Method not found", id)
	}
}

// handleInitialize 处理初始化请求
func (s *StdioServer) handleInitialize(id interface{}, request map[string]interface{}) {
	response := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      id,
		"result": map[string]interface{}{
			"protocolVersion": "2024-11-05",
			"capabilities": map[string]interface{}{
				"tools": map[string]interface{}{},
			},
			"serverInfo": map[string]interface{}{
				"name":    "mcp-whoami-server",
				"version": "1.0.0",
			},
		},
	}
	s.sendResponse(response)
}

// handleToolsList 处理工具列表请求
func (s *StdioServer) handleToolsList(id interface{}) {
	toolsList := s.personTools.GetToolsList()

	response := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      id,
		"result": map[string]interface{}{
			"tools": toolsList,
		},
	}
	s.sendResponse(response)
}

// handleToolsCall 处理工具调用请求
func (s *StdioServer) handleToolsCall(id interface{}, request map[string]interface{}) {
	params, ok := request["params"].(map[string]interface{})
	if !ok {
		s.sendError(-32602, "Invalid params", id)
		return
	}

	name, ok := params["name"].(string)
	if !ok {
		s.sendError(-32602, "Missing tool name", id)
		return
	}

	arguments, ok := params["arguments"].(map[string]interface{})
	if !ok {
		arguments = make(map[string]interface{})
	}

	log.Printf("调用工具: %s, 参数: %v", name, arguments)

	// 调用工具
	result, err := s.personTools.CallTool(name, arguments)

	// 格式化响应
	toolResult := s.personTools.FormatResult(result, err)

	response := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      id,
		"result":  toolResult,
	}
	s.sendResponse(response)
}

// handlePing 处理ping请求
func (s *StdioServer) handlePing(id interface{}) {
	response := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      id,
		"result":  map[string]interface{}{},
	}
	s.sendResponse(response)
}

// sendResponse 发送响应
func (s *StdioServer) sendResponse(response map[string]interface{}) {
	data, err := json.Marshal(response)
	if err != nil {
		log.Printf("序列化响应失败: %v", err)
		return
	}
	fmt.Println(string(data))
}

// sendError 发送错误响应
func (s *StdioServer) sendError(code int, message string, id interface{}) {
	response := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      id,
		"error": map[string]interface{}{
			"code":    code,
			"message": message,
		},
	}
	s.sendResponse(response)
}
