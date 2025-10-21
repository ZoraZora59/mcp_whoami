package storage

import (
	"business-service/model"
	"errors"
	"sync"
)

// MemoryStorage 内存存储实现
type MemoryStorage struct {
	data map[string]*model.Person
	mu   sync.RWMutex
}

// NewMemoryStorage 创建内存存储实例
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]*model.Person),
	}
}

// Create 创建人员信息
func (s *MemoryStorage) Create(person *model.Person) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.data[person.WorkId]; exists {
		return errors.New("工号已存在")
	}

	s.data[person.WorkId] = person
	return nil
}

// Get 根据工号获取人员信息
func (s *MemoryStorage) Get(workId string) (*model.Person, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	person, exists := s.data[workId]
	if !exists {
		return nil, errors.New("人员不存在")
	}

	return person, nil
}

// List 获取所有人员信息
func (s *MemoryStorage) List() ([]*model.Person, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	persons := make([]*model.Person, 0, len(s.data))
	for _, person := range s.data {
		persons = append(persons, person)
	}

	return persons, nil
}

// Update 更新人员信息
func (s *MemoryStorage) Update(workId string, person *model.Person) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.data[workId]; !exists {
		return errors.New("人员不存在")
	}

	s.data[workId] = person
	return nil
}

// Delete 删除人员信息
func (s *MemoryStorage) Delete(workId string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.data[workId]; !exists {
		return errors.New("人员不存在")
	}

	delete(s.data, workId)
	return nil
}

// Exists 检查工号是否存在
func (s *MemoryStorage) Exists(workId string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, exists := s.data[workId]
	return exists
}
