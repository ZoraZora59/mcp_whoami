package tools

import (
	"encoding/json"
	"mcp-service/client"
)

// PersonTools MCP工具定义
type PersonTools struct {
	businessClient *client.BusinessClient
}

// NewPersonTools 创建MCP工具实例
func NewPersonTools(businessClient *client.BusinessClient) *PersonTools {
	return &PersonTools{
		businessClient: businessClient,
	}
}

// GetToolsList 获取工具列表
func (t *PersonTools) GetToolsList() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"name":        "create_person",
			"description": "创建新的人员信息",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"workId": map[string]interface{}{
						"type":        "string",
						"description": "工号（唯一标识）",
					},
					"name": map[string]interface{}{
						"type":        "string",
						"description": "姓名",
					},
					"gender": map[string]interface{}{
						"type":        "string",
						"description": "性别",
					},
					"age": map[string]interface{}{
						"type":        "integer",
						"description": "年龄",
					},
				},
				"required": []string{"workId", "name", "gender", "age"},
			},
		},
		{
			"name":        "get_person",
			"description": "根据工号（密令）查询人员信息",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"workId": map[string]interface{}{
						"type":        "string",
						"description": "工号（唯一标识/密令）",
					},
				},
				"required": []string{"workId"},
			},
		},
		{
			"name":        "list_persons",
			"description": "列出所有人员信息",
			"inputSchema": map[string]interface{}{
				"type":       "object",
				"properties": map[string]interface{}{},
			},
		},
		{
			"name":        "update_person",
			"description": "更新人员信息",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"workId": map[string]interface{}{
						"type":        "string",
						"description": "工号（唯一标识）",
					},
					"name": map[string]interface{}{
						"type":        "string",
						"description": "姓名",
					},
					"gender": map[string]interface{}{
						"type":        "string",
						"description": "性别",
					},
					"age": map[string]interface{}{
						"type":        "integer",
						"description": "年龄",
					},
				},
				"required": []string{"workId", "name", "gender", "age"},
			},
		},
		{
			"name":        "delete_person",
			"description": "删除人员信息",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"workId": map[string]interface{}{
						"type":        "string",
						"description": "工号（唯一标识）",
					},
				},
				"required": []string{"workId"},
			},
		},
	}
}

// CallTool 调用工具
func (t *PersonTools) CallTool(name string, arguments map[string]interface{}) (interface{}, error) {
	switch name {
	case "create_person":
		return t.createPerson(arguments)
	case "get_person":
		return t.getPerson(arguments)
	case "list_persons":
		return t.listPersons(arguments)
	case "update_person":
		return t.updatePerson(arguments)
	case "delete_person":
		return t.deletePerson(arguments)
	default:
		return nil, nil
	}
}

func (t *PersonTools) createPerson(args map[string]interface{}) (interface{}, error) {
	person := &client.Person{
		WorkId: args["workId"].(string),
		Name:   args["name"].(string),
		Gender: args["gender"].(string),
		Age:    int(args["age"].(float64)),
	}
	return t.businessClient.CreatePerson(person)
}

func (t *PersonTools) getPerson(args map[string]interface{}) (interface{}, error) {
	workId := args["workId"].(string)
	return t.businessClient.GetPerson(workId)
}

func (t *PersonTools) listPersons(args map[string]interface{}) (interface{}, error) {
	return t.businessClient.ListPersons()
}

func (t *PersonTools) updatePerson(args map[string]interface{}) (interface{}, error) {
	person := &client.Person{
		WorkId: args["workId"].(string),
		Name:   args["name"].(string),
		Gender: args["gender"].(string),
		Age:    int(args["age"].(float64)),
	}
	workId := args["workId"].(string)
	return t.businessClient.UpdatePerson(workId, person)
}

func (t *PersonTools) deletePerson(args map[string]interface{}) (interface{}, error) {
	workId := args["workId"].(string)
	return t.businessClient.DeletePerson(workId)
}

// FormatResult 格式化结果为MCP响应格式
func (t *PersonTools) FormatResult(result interface{}, err error) map[string]interface{} {
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": "错误: " + err.Error(),
				},
			},
			"isError": true,
		}
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	return map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": string(resultJSON),
			},
		},
	}
}
