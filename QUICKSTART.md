# 快速开始指南

## 一、启动服务

### 1. 启动业务服务（终端1）

```bash
cd /root/gitee/mcp_whoami
./start-business.sh
```

你会看到类似输出：
```
[GIN-debug] Listening and serving HTTP on :8080
业务服务启动在 :8080 端口
```

**自定义端口启动：**
```bash
cd /root/gitee/mcp_whoami/business-service
go run main.go -p 9090
```

### 2. 启动MCP服务（终端2）

```bash
cd /root/gitee/mcp_whoami
./start-mcp.sh
```

你会看到类似输出：
```
可用工具:
[
  {
    "name": "create_person",
    "description": "创建新的人员信息",
    ...
  },
  ...
]
MCP服务启动在 :8081 端口
```

**自定义端口启动：**
```bash
cd /root/gitee/mcp_whoami/mcp-service
go run main.go -p 9091
```

## 二、测试服务

### 在第三个终端运行测试脚本：

```bash
cd /root/gitee/mcp_whoami
./test-api.sh
```

## 三、手动测试示例

### 测试业务服务

1. **创建人员**
```bash
curl -X POST http://localhost:8080/api/person \
  -H "Content-Type: application/json" \
  -d '{"workId":"E001","name":"张三","gender":"男","age":30}'
```

2. **查询人员（通过密令/工号）**
```bash
curl http://localhost:8080/api/person/E001
```

3. **列出所有人员**
```bash
curl http://localhost:8080/api/persons
```

### 测试MCP服务

1. **获取可用工具列表**
```bash
curl -X POST http://localhost:8081/mcp/tools/list
```

2. **通过MCP查询人员**
```bash
curl -X POST http://localhost:8081/mcp/tools/call \
  -H "Content-Type: application/json" \
  -d '{
    "params": {
      "name": "get_person",
      "arguments": {
        "workId": "E001"
      }
    }
  }'
```

3. **通过MCP创建人员**
```bash
curl -X POST http://localhost:8081/mcp/tools/call \
  -H "Content-Type: application/json" \
  -d '{
    "params": {
      "name": "create_person",
      "arguments": {
        "workId": "E002",
        "name": "李四",
        "gender": "女",
        "age": 28
      }
    }
  }'
```

## 四、服务说明

### 业务服务（默认端口：8080）
- 前端API路径：`/api/*`
- 内部API路径：`/internal/*`（供MCP服务调用）
- 可通过 `-p` 参数自定义端口

### MCP服务（默认端口：8081）
- MCP协议端点：`/mcp/*`
- 健康检查：`/health`
- 可通过 `-p` 参数自定义端口

## 五、核心概念

1. **密令**：即人员的工号（workId），作为唯一标识符
2. **人员信息**：包含工号、姓名、性别、年龄四个字段
3. **内存存储**：数据保存在内存中，服务重启后数据会丢失
4. **MCP协议**：为大模型Agent提供标准化的工具调用接口

## 六、常见问题

**Q: 服务启动失败，提示端口被占用？**
A: 可以使用 `-p` 参数指定其他端口启动服务，例如：
```bash
# 业务服务使用 9090 端口
go run main.go -p 9090

# MCP 服务使用 9091 端口
go run main.go -p 9091
```

**Q: 数据会持久化吗？**
A: 当前使用内存存储，服务重启后数据会丢失。如需持久化，可按README中的说明切换到MySQL。

**Q: MCP服务无法连接业务服务？**
A: 确保业务服务已启动，且地址为 http://localhost:8080。如有需要，通过环境变量 BUSINESS_SERVICE_URL 指定正确地址。

**Q: 如何在Claude Desktop中使用MCP服务？**
A: 在Claude Desktop的MCP配置文件中添加本服务的HTTP端点配置，具体配置方式请参考MCP官方文档。

