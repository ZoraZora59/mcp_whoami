#!/bin/bash

echo "清理构建产物..."

# 清理业务服务的二进制文件
if [ -f business-service/business-service ]; then
    rm business-service/business-service
    echo "✓ 已删除 business-service/business-service"
fi

# 清理MCP服务的二进制文件
if [ -f mcp-service/mcp-service ]; then
    rm mcp-service/mcp-service
    echo "✓ 已删除 mcp-service/mcp-service"
fi

# 清理可能的日志文件
if [ -f nohup.out ]; then
    rm nohup.out
    echo "✓ 已删除 nohup.out"
fi

echo ""
echo "清理完成！"

