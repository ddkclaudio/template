package dto

import (
	"errors"
	"fmt"
)

type FilterRequestDTO struct {
	Owner    string `json:"owner,omitempty"`
	OnlyMe   bool   `json:"only_me,omitempty"`
	Page     int    `json:"page,omitempty"`
	PageSize int    `json:"page_size,omitempty"`
}

func NewFilterRequestDTO() *FilterRequestDTO {
	return &FilterRequestDTO{
		OnlyMe:   true,
		Page:     1,
		PageSize: 10,
	}
}

func (p *FilterRequestDTO) Validate() (bool, error) {
	if p.Page < 1 {
		return false, errors.New("field 'page' must be >= 1")
	}

	if p.PageSize < 1 || p.PageSize > 100 {
		return false, fmt.Errorf("field 'page_size' must be between 1 and 100")
	}

	return true, nil
}
