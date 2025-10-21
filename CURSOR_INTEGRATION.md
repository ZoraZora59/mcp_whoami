# Cursor 集成指南

## 配置步骤

### 1. 确保业务服务已启动

MCP 服务需要连接业务服务，所以首先要启动业务服务：

```bash
cd /root/gitee/mcp_whoami
./start-business.sh
```

或者使用后台方式启动：

```bash
cd /root/gitee/mcp_whoami/business-service
nohup ./business-service > /tmp/business-service.log 2>&1 &
```

### 2. 配置 Cursor MCP

MCP 服务已经配置在 `/root/.cursor/mcp.json` 文件中：

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

### 3. 重启 Cursor

修改 `mcp.json` 后，需要重启 Cursor 才能生效。

### 4. 验证连接

在 Cursor 中，你可以：

1. 打开 Cursor 的 MCP 设置页面，查看 `whoami` 服务是否已连接
2. 在聊天中尝试使用工具，例如：
   - "帮我创建一个工号为 E001 的员工，叫张三，男，30岁"
   - "查询工号 E001 的员工信息"
   - "列出所有员工"

## 可用工具

MCP 服务提供以下 5 个工具：

### 1. create_person
创建新的人员信息

**参数：**
- `workId` (string): 工号（唯一标识）
- `name` (string): 姓名
- `gender` (string): 性别
- `age` (integer): 年龄

### 2. get_person
根据工号查询人员信息

**参数：**
- `workId` (string): 工号

### 3. list_persons
列出所有人员信息

**参数：** 无

### 4. update_person
更新人员信息

**参数：**
- `workId` (string): 工号
- `name` (string): 姓名
- `gender` (string): 性别
- `age` (integer): 年龄

### 5. delete_person
删除人员信息

**参数：**
- `workId` (string): 工号

## 使用示例

在 Cursor 聊天中，你可以这样使用：

```
你: 帮我创建一个员工信息：工号 E001，姓名张三，男，30岁

AI: [会调用 create_person 工具]

你: 查询 E001 的信息

AI: [会调用 get_person 工具]

你: 列出所有员工

AI: [会调用 list_persons 工具]

你: 把 E001 的年龄改成 31 岁

AI: [会调用 update_person 工具]
```

## 故障排查

### MCP 服务无法连接

1. **检查业务服务是否运行：**
   ```bash
   curl http://localhost:8080/health
   ```
   
2. **检查 MCP 可执行文件是否存在：**
   ```bash
   ls -l /root/gitee/mcp_whoami/mcp-service/mcp-service
   ```
   
3. **重新编译 MCP 服务：**
   ```bash
   cd /root/gitee/mcp_whoami
   ./build.sh
   ```

4. **查看 Cursor 日志：**
   在 Cursor 中打开开发者工具（Help > Toggle Developer Tools），查看 Console 中的错误信息

### 业务服务端口被占用

如果 8080 端口被占用，可以修改业务服务端口：

1. 修改启动业务服务的端口：
   ```bash
   cd /root/gitee/mcp_whoami/business-service
   ./business-service -p 9090
   ```

2. 更新 `/root/.cursor/mcp.json` 中的配置：
   ```json
   {
     "mcpServers": {
       "whoami": {
         "command": "/root/gitee/mcp_whoami/mcp-service/mcp-service",
         "args": ["--stdio"],
         "env": {
           "BUSINESS_SERVICE_URL": "http://localhost:9090"
         }
       }
     }
   }
   ```

3. 重启 Cursor

## 部署到生产环境

如果要将服务部署到生产环境：

1. **构建服务：**
   ```bash
   cd /root/gitee/mcp_whoami
   ./build.sh
   ```

2. **部署：**
   ```bash
   sudo ./deploy.sh
   ```

3. **启动业务服务（后台运行）：**
   ```bash
   cd /www/wwwroot/whoami.zorazora.cn
   nohup ./business-service > business-service.log 2>&1 &
   ```

4. **更新 MCP 配置指向生产服务：**
   ```json
   {
     "mcpServers": {
       "whoami": {
         "command": "/www/wwwroot/whoami.zorazora.cn/mcp-service",
         "args": ["--stdio"],
         "env": {
           "BUSINESS_SERVICE_URL": "http://localhost:8080"
         }
       }
     }
   }
   ```

## 注意事项

1. **数据持久化：** 当前使用内存存储，业务服务重启后数据会丢失
2. **业务服务必须先启动：** MCP 服务依赖业务服务，必须确保业务服务在运行
3. **stdio 模式：** Cursor 使用的是 stdio 模式，不是 HTTP 模式
4. **日志位置：** stdio 模式的日志会输出到 stderr，Cursor 可能会捕获这些日志

