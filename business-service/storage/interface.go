package storage

import "business-service/model"

// PersonStorage 人员存储接口，为后续切换到MySQL做准备
type PersonStorage interface {
	// Create 创建人员信息
	Create(person *model.Person) error

	// Get 根据工号获取人员信息
	Get(workId string) (*model.Person, error)

	// List 获取所有人员信息
	List() ([]*model.Person, error)

	// Update 更新人员信息
	Update(workId string, person *model.Person) error

	// Delete 删除人员信息
	Delete(workId string) error

	// Exists 检查工号是否存在
	Exists(workId string) bool
}
