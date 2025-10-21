package service

import (
	"business-service/model"
	"business-service/storage"
)

// PersonService 人员服务
type PersonService struct {
	storage storage.PersonStorage
}

// NewPersonService 创建人员服务实例
func NewPersonService(storage storage.PersonStorage) *PersonService {
	return &PersonService{
		storage: storage,
	}
}

// CreatePerson 创建人员
func (s *PersonService) CreatePerson(person *model.Person) error {
	return s.storage.Create(person)
}

// GetPerson 根据工号获取人员
func (s *PersonService) GetPerson(workId string) (*model.Person, error) {
	return s.storage.Get(workId)
}

// ListPersons 获取所有人员
func (s *PersonService) ListPersons() ([]*model.Person, error) {
	return s.storage.List()
}

// UpdatePerson 更新人员信息
func (s *PersonService) UpdatePerson(workId string, person *model.Person) error {
	return s.storage.Update(workId, person)
}

// DeletePerson 删除人员
func (s *PersonService) DeletePerson(workId string) error {
	return s.storage.Delete(workId)
}
