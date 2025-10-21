package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Person 人员信息模型
type Person struct {
	WorkId string `json:"workId"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

// BusinessClient 业务服务客户端
type BusinessClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewBusinessClient 创建业务服务客户端
func NewBusinessClient(baseURL string) *BusinessClient {
	return &BusinessClient{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

// CreatePerson 创建人员
func (c *BusinessClient) CreatePerson(person *Person) (map[string]interface{}, error) {
	data, err := json.Marshal(person)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Post(
		c.baseURL+"/internal/person",
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		if errMsg, ok := result["error"].(string); ok {
			return nil, fmt.Errorf("%s", errMsg)
		}
		return nil, fmt.Errorf("请求失败: %d", resp.StatusCode)
	}

	return result, nil
}

// GetPerson 获取人员信息
func (c *BusinessClient) GetPerson(workId string) (map[string]interface{}, error) {
	resp, err := c.httpClient.Get(c.baseURL + "/internal/person/" + workId)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		if errMsg, ok := result["error"].(string); ok {
			return nil, fmt.Errorf("%s", errMsg)
		}
		return nil, fmt.Errorf("请求失败: %d", resp.StatusCode)
	}

	return result, nil
}

// ListPersons 列出所有人员
func (c *BusinessClient) ListPersons() (map[string]interface{}, error) {
	resp, err := c.httpClient.Get(c.baseURL + "/internal/persons")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		if errMsg, ok := result["error"].(string); ok {
			return nil, fmt.Errorf("%s", errMsg)
		}
		return nil, fmt.Errorf("请求失败: %d", resp.StatusCode)
	}

	return result, nil
}

// UpdatePerson 更新人员信息
func (c *BusinessClient) UpdatePerson(workId string, person *Person) (map[string]interface{}, error) {
	data, err := json.Marshal(person)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPut,
		c.baseURL+"/internal/person/"+workId,
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		if errMsg, ok := result["error"].(string); ok {
			return nil, fmt.Errorf("%s", errMsg)
		}
		return nil, fmt.Errorf("请求失败: %d", resp.StatusCode)
	}

	return result, nil
}

// DeletePerson 删除人员
func (c *BusinessClient) DeletePerson(workId string) (map[string]interface{}, error) {
	req, err := http.NewRequest(
		http.MethodDelete,
		c.baseURL+"/internal/person/"+workId,
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		if errMsg, ok := result["error"].(string); ok {
			return nil, fmt.Errorf("%s", errMsg)
		}
		return nil, fmt.Errorf("请求失败: %d", resp.StatusCode)
	}

	return result, nil
}
