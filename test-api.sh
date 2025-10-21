#!/bin/bash

echo "================================"
echo "测试业务服务API"
echo "================================"
echo ""

BASE_URL="http://localhost:8080"

echo "1. 创建人员 E001 - 张三"
curl -X POST ${BASE_URL}/api/person \
  -H "Content-Type: application/json" \
  -d '{"workId":"E001","name":"张三","gender":"男","age":30}' \
  -s | jq .
echo ""

echo "2. 创建人员 E002 - 李四"
curl -X POST ${BASE_URL}/api/person \
  -H "Content-Type: application/json" \
  -d '{"workId":"E002","name":"李四","gender":"女","age":28}' \
  -s | jq .
echo ""

echo "3. 查询人员 E001"
curl ${BASE_URL}/api/person/E001 -s | jq .
echo ""

echo "4. 列出所有人员"
curl ${BASE_URL}/api/persons -s | jq .
echo ""

echo "5. 更新人员 E001"
curl -X PUT ${BASE_URL}/api/person/E001 \
  -H "Content-Type: application/json" \
  -d '{"name":"张三","gender":"男","age":31}' \
  -s | jq .
echo ""

echo "6. 再次查询人员 E001（验证更新）"
curl ${BASE_URL}/api/person/E001 -s | jq .
echo ""

echo "================================"
echo "测试MCP服务"
echo "================================"
echo ""

MCP_URL="http://localhost:8081"

echo "7. 获取MCP工具列表"
curl -X POST ${MCP_URL}/mcp/tools/list \
  -H "Content-Type: application/json" \
  -s | jq .
echo ""

echo "8. 通过MCP查询人员 E001"
curl -X POST ${MCP_URL}/mcp/tools/call \
  -H "Content-Type: application/json" \
  -d '{"params":{"name":"get_person","arguments":{"workId":"E001"}}}' \
  -s | jq .
echo ""

echo "9. 通过MCP列出所有人员"
curl -X POST ${MCP_URL}/mcp/tools/call \
  -H "Content-Type: application/json" \
  -d '{"params":{"name":"list_persons","arguments":{}}}' \
  -s | jq .
echo ""

echo "10. 通过MCP创建人员 E003"
curl -X POST ${MCP_URL}/mcp/tools/call \
  -H "Content-Type: application/json" \
  -d '{"params":{"name":"create_person","arguments":{"workId":"E003","name":"王五","gender":"男","age":35}}}' \
  -s | jq .
echo ""

echo "================================"
echo "测试完成"
echo "================================"

