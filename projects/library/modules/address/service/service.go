package service

import (
	"github.com/library/modules/address/domain"
	"github.com/library/modules/address/dto"
	"github.com/library/modules/address/mapper"
)

type MainService struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *MainService {
	return &MainService{repo: repo}
}

func (s *MainService) Create(input *dto.CreateRequestDTO) (*domain.Entity, error) {
	newEntity := mapper.ToEntityFromCreateDTO(input)
	return s.repo.Create(newEntity)
}

func (s *MainService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *MainService) Get(id string) (*domain.Entity, error) {
	return s.repo.Get(id)
}

func (s *MainService) List(filter dto.FilterRequestDTO) ([]*domain.Entity, error) {
	list, err := s.repo.List(filter)
	if err != nil {
		return nil, err
	}

	if list == nil {
		list = []*domain.Entity{}
	}

	return list, nil
}

func (s *MainService) Update(input *domain.Entity) (*domain.Entity, error) {
	return s.repo.Update(input)
}
