#!/bin/bash

echo "启动MCP服务..."
cd mcp-service

# 如果需要指定业务服务地址，请设置环境变量
# export BUSINESS_SERVICE_URL=http://localhost:8080

go run main.go

