package domain

import "github.com/library/modules/address/dto"

type Repository interface {
	Create(entity *Entity) (*Entity, error)
	Delete(id string) error
	Get(id string) (*Entity, error)
	List(filter dto.FilterRequestDTO) ([]*Entity, error)
	Update(entity *Entity) (*Entity, error)
}
