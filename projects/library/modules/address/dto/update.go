package dto

import (
	"fmt"

	"github.com/library/utils"
)

type UpdateRequestDTO struct {
	BannerImage  *string  `json:"banner_image,omitempty"`
	City         *string  `json:"city,omitempty"`
	Complement   *string  `json:"complement,omitempty"`
	Country      *string  `json:"country,omitempty"`
	Description  *string  `json:"description,omitempty"`
	IsActive     *bool    `json:"is_active,omitempty"`
	IsDeleted    *bool    `json:"is_deleted,omitempty"`
	IsPrimary    *bool    `json:"is_primary,omitempty"`
	Latitude     *float64 `json:"latitude,omitempty"`
	Longitude    *float64 `json:"longitude,omitempty"`
	Neighborhood *string  `json:"neighborhood,omitempty"`
	Number       *string  `json:"number,omitempty"`
	Owner        *string  `json:"owner,omitempty"`
	ProfileImage *string  `json:"profile_image,omitempty"`
	State        *string  `json:"state,omitempty"`
	Street       *string  `json:"street,omitempty"`
	Type         *string  `json:"type,omitempty"`
	Zip          *string  `json:"zip,omitempty"`
}

func (e UpdateRequestDTO) Validate() error {
	if e.City != nil {
		if err := utils.ValidateMinLength(*e.City, 1); err != nil {
			return fmt.Errorf("field 'city': %w", err)
		}
	}
	if e.Country != nil {
		if err := utils.ValidateMinLength(*e.Country, 1); err != nil {
			return fmt.Errorf("field 'country': %w", err)
		}
	}
	if e.Latitude != nil {
		if err := utils.ValidateMinimum(float64(*e.Latitude), -90); err != nil {
			return fmt.Errorf("field 'latitude': %w", err)
		}
		if err := utils.ValidateMaximum(float64(*e.Latitude), 90); err != nil {
			return fmt.Errorf("field 'latitude': %w", err)
		}
	}
	if e.Longitude != nil {
		if err := utils.ValidateMinimum(float64(*e.Longitude), -180); err != nil {
			return fmt.Errorf("field 'longitude': %w", err)
		}
		if err := utils.ValidateMaximum(float64(*e.Longitude), 180); err != nil {
			return fmt.Errorf("field 'longitude': %w", err)
		}
	}
	if e.Neighborhood != nil {
		if err := utils.ValidateMinLength(*e.Neighborhood, 1); err != nil {
			return fmt.Errorf("field 'neighborhood': %w", err)
		}
	}
	if e.Number != nil {
		if err := utils.ValidateMinLength(*e.Number, 1); err != nil {
			return fmt.Errorf("field 'number': %w", err)
		}
	}
	if e.State != nil {
		if err := utils.ValidateMinLength(*e.State, 2); err != nil {
			return fmt.Errorf("field 'state': %w", err)
		}
		if err := utils.ValidateMaxLength(*e.State, 2); err != nil {
			return fmt.Errorf("field 'state': %w", err)
		}
	}
	if e.Street != nil {
		if err := utils.ValidateMinLength(*e.Street, 1); err != nil {
			return fmt.Errorf("field 'street': %w", err)
		}
	}
	if e.Zip != nil {
		if err := utils.ValidateMinLength(*e.Zip, 1); err != nil {
			return fmt.Errorf("field 'zip': %w", err)
		}
	}

	return nil
}
