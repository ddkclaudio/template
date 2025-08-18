package dto

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
	return true, nil
}
