package model

// Person 人员信息模型
type Person struct {
	WorkId string `json:"workId"` // 工号（唯一标识/密令）
	Name   string `json:"name"`   // 姓名
	Gender string `json:"gender"` // 性别
	Age    int    `json:"age"`    // 年龄
}
