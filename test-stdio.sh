#!/bin/bash

# 测试MCP stdio模式

echo "测试MCP stdio模式..."
echo ""

# 确保业务服务正在运行
echo "检查业务服务..."
if ! curl -s http://localhost:8080/health > /dev/null; then
    echo "错误: 业务服务未运行，请先启动业务服务"
    echo "运行: ./start-business.sh"
    exit 1
fi
echo "✓ 业务服务正在运行"
echo ""

# 测试初始化
echo "1. 测试初始化..."
echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","clientInfo":{"name":"test","version":"1.0.0"}}}' | \
    cd mcp-service && ./mcp-service --stdio 2>/dev/null | head -1 &
sleep 1
echo ""

# 测试工具列表
echo "2. 测试获取工具列表..."
echo '{"jsonrpc":"2.0","id":2,"method":"tools/list"}' | \
    cd mcp-service && ./mcp-service --stdio 2>/dev/null | head -1 | jq .
echo ""

# 测试创建人员
echo "3. 测试创建人员..."
echo '{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"create_person","arguments":{"workId":"E999","name":"测试用户","gender":"男","age":25}}}' | \
    cd mcp-service && ./mcp-service --stdio 2>/dev/null | head -1 | jq .
echo ""

# 测试查询人员
echo "4. 测试查询人员..."
echo '{"jsonrpc":"2.0","id":4,"method":"tools/call","params":{"name":"get_person","arguments":{"workId":"E999"}}}' | \
    cd mcp-service && ./mcp-service --stdio 2>/dev/null | head -1 | jq .
echo ""

echo "测试完成！"
echo ""
echo "如果所有测试都通过，说明 stdio 模式工作正常。"
echo "现在可以在 Cursor 中使用此 MCP 服务了。"

