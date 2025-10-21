#!/bin/bash

set -e

echo "开始构建项目..."

# 构建业务服务
echo "构建 business-service..."
cd business-service
go build -o business-service main.go
echo "✓ business-service 构建完成"
cd ..

# 构建MCP服务
echo "构建 mcp-service..."
cd mcp-service
go build -o mcp-service main.go
echo "✓ mcp-service 构建完成"
cd ..

echo ""
echo "所有服务构建完成！"
echo "二进制文件位置："
echo "  - business-service/business-service"
echo "  - mcp-service/mcp-service"

