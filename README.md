# MCP WhoAmI 服务

这是一个基于Go语言构建的人员信息管理系统，包含业务服务和MCP（Model Context Protocol）服务两个部分。

## 项目结构

```
mcp_whoami/
├── business-service/          # 业务服务
│   ├── main.go               # 服务入口
│   ├── model/                # 数据模型
│   │   └── person.go
│   ├── storage/              # 存储层
│   │   ├── interface.go      # 存储接口（为MySQL预留）
│   │   └── memory.go         # 内存存储实现
│   ├── service/              # 业务逻辑层
│   │   └── person_service.go
│   ├── handler/              # HTTP处理器
│   │   ├── frontend.go       # 前端API
│   │   └── internal.go       # 内部API（供MCP服务调用）
│   └── go.mod
├── mcp-service/              # MCP服务
│   ├── main.go               # 服务入口
│   ├── server/               # MCP服务器
│   │   └── mcp_server.go
│   ├── tools/                # MCP工具定义
│   │   └── person_tools.go
│   ├── client/               # 业务服务客户端
│   │   └── business_client.go
│   └── go.mod
└── README.md
```

## 功能特性

### 业务服务

业务服务提供完整的人员信息管理功能：

- **数据模型**：包含工号、姓名、性别、年龄四个字段
- **内存存储**：使用map结构存储数据，支持并发访问（读写锁保护）
- **前端API**：提供RESTful接口供前端调用
- **内部API**：独立的内部接口供MCP服务调用
- **可扩展设计**：存储层采用接口设计，便于后续切换到MySQL

### MCP服务

MCP服务实现标准的MCP HTTP协议，为大模型Agent提供工具调用能力：

- **标准协议**：遵循MCP协议规范
- **5个工具**：创建、查询、列出、更新、删除人员信息
- **HTTP通信**：通过HTTP与业务服务通信
- **JSON格式**：所有数据使用JSON格式传输

## 快速开始

### 前置要求

- Go 1.21 或更高版本
- 8080端口（业务服务默认端口）和8081端口（MCP服务默认端口）可用，或通过 `-p` 参数指定其他端口

### 启动业务服务

使用启动脚本：
```bash
./start-business.sh
```

或手动启动：
```bash
cd business-service
go run main.go
```

业务服务将在 `http://localhost:8080` 启动（默认端口）

**自定义端口：**
```bash
# 指定端口 9090
go run main.go -p 9090
```

### 启动MCP服务

使用启动脚本：
```bash
./start-mcp.sh
```

或手动启动：
```bash
cd mcp-service
go run main.go
```

MCP服务将在 `http://localhost:8081` 启动（默认端口）

**自定义端口：**
```bash
# 指定端口 9091
go run main.go -p 9091
```

**指定业务服务地址：**

如果业务服务不在默认地址，可以通过环境变量指定：

```bash
BUSINESS_SERVICE_URL=http://your-host:port go run main.go
```

### 运行测试

项目提供了一个测试脚本 `test-api.sh`，可以快速测试所有API功能：

```bash
# 先确保两个服务都已启动，然后在新终端运行
./test-api.sh
```

测试脚本将执行以下操作：
1. 创建多个人员
2. 查询人员信息
3. 列出所有人员
4. 更新人员信息
5. 测试MCP工具调用

## API文档

### 业务服务 - 前端接口

基础路径：`/api`

#### 创建人员

```
POST /api/person
Content-Type: application/json

{
  "workId": "E001",
  "name": "张三",
  "gender": "男",
  "age": 30
}
```

#### 查询人员

```
GET /api/person/:workId
```

#### 列出所有人员

```
GET /api/persons
```

#### 更新人员

```
PUT /api/person/:workId
Content-Type: application/json

{
  "name": "李四",
  "gender": "女",
  "age": 28
}
```

#### 删除人员

```
DELETE /api/person/:workId
```

### 业务服务 - 内部接口

基础路径：`/internal`

接口与前端接口完全相同，只是路径前缀不同。供MCP服务内部调用。

### MCP服务接口

#### 初始化

```
POST /mcp/initialize
Content-Type: application/json

{
  "protocolVersion": "2024-11-05",
  "clientInfo": {
    "name": "client-name",
    "version": "1.0.0"
  }
}
```

#### 获取工具列表

```
POST /mcp/tools/list
```

#### 调用工具

```
POST /mcp/tools/call
Content-Type: application/json

{
  "params": {
    "name": "get_person",
    "arguments": {
      "workId": "E001"
    }
  }
}
```

## MCP工具说明

MCP服务提供以下5个工具：

### 1. create_person

创建新的人员信息

**参数：**
- `workId` (string, 必需)：工号（唯一标识）
- `name` (string, 必需)：姓名
- `gender` (string, 必需)：性别
- `age` (integer, 必需)：年龄

### 2. get_person

根据工号查询人员信息

**参数：**
- `workId` (string, 必需)：工号（密令）

### 3. list_persons

列出所有人员信息

**参数：** 无

### 4. update_person

更新人员信息

**参数：**
- `workId` (string, 必需)：工号（唯一标识）
- `name` (string, 必需)：姓名
- `gender` (string, 必需)：性别
- `age` (integer, 必需)：年龄

### 5. delete_person

删除人员信息

**参数：**
- `workId` (string, 必需)：工号（唯一标识）

## 使用示例

### 使用curl测试业务服务

创建人员：
```bash
curl -X POST http://localhost:8080/api/person \
  -H "Content-Type: application/json" \
  -d '{"workId":"E001","name":"张三","gender":"男","age":30}'
```

查询人员：
```bash
curl http://localhost:8080/api/person/E001
```

### 使用curl测试MCP服务

获取工具列表：
```bash
curl -X POST http://localhost:8081/mcp/tools/list
```

调用get_person工具：
```bash
curl -X POST http://localhost:8081/mcp/tools/call \
  -H "Content-Type: application/json" \
  -d '{"params":{"name":"get_person","arguments":{"workId":"E001"}}}'
```

## 技术栈

- **语言**：Go 1.21+
- **Web框架**：Gin
- **协议**：HTTP, MCP
- **存储**：内存存储（可扩展至MySQL）

## 设计特点

1. **分层架构**：模型层、存储层、服务层、处理器层清晰分离
2. **接口设计**：存储层使用接口，便于后续切换数据库
3. **并发安全**：内存存储使用读写锁保护
4. **独立服务**：业务服务和MCP服务独立部署和运行
5. **标准协议**：MCP服务遵循MCP协议规范

## Cursor 集成

MCP 服务支持两种模式：

1. **HTTP 模式**（用于测试和调试）
2. **stdio 模式**（用于 Cursor 等 AI 客户端）

### 在 Cursor 中使用

1. **启动业务服务：**
   ```bash
   ./start-business.sh
   ```

2. **配置 Cursor MCP（已自动配置）：**
   配置文件位于 `/root/.cursor/mcp.json`：
   ```json
   {
     "mcpServers": {
       "whoami": {
         "command": "/root/gitee/mcp_whoami/mcp-service/mcp-service",
         "args": ["--stdio"],
         "env": {
           "BUSINESS_SERVICE_URL": "http://localhost:8080"
         }
       }
     }
   }
   ```

3. **重启 Cursor**

4. **在 Cursor 中使用：**
   - "帮我创建一个工号为 E001 的员工，叫张三，男，30岁"
   - "查询工号 E001 的员工信息"
   - "列出所有员工"

详细说明请参考 [CURSOR_INTEGRATION.md](CURSOR_INTEGRATION.md)

## 后续扩展

### 切换到MySQL

1. 实现 `storage.PersonStorage` 接口的MySQL版本
2. 在 `main.go` 中替换存储实例：
   ```go
   // 替换
   personStorage := storage.NewMemoryStorage()
   // 为
   personStorage := storage.NewMySQLStorage(db)
   ```

### 添加更多字段

1. 修改 `model/person.go` 添加字段
2. 更新对应的存储层和API处理

### 添加MCP工具

1. 在 `tools/person_tools.go` 中添加新工具定义
2. 在 `CallTool` 方法中添加处理逻辑
3. 在业务服务中添加对应的接口支持

## 许可证

MIT License

