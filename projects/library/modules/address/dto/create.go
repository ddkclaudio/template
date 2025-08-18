package dto

import (
	"fmt"

	"github.com/library/utils"
)

type CreateRequestDTO struct {
	BannerImage  *string  `json:"banner_image"`
	City         string   `json:"city"`
	Complement   *string  `json:"complement"`
	Country      string   `json:"country"`
	Description  *string  `json:"description"`
	IsActive     *bool    `json:"is_active"`
	IsDeleted    *bool    `json:"is_deleted"`
	IsPrimary    *bool    `json:"is_primary"`
	Latitude     *float64 `json:"latitude"`
	Longitude    *float64 `json:"longitude"`
	Neighborhood string   `json:"neighborhood"`
	Number       string   `json:"number"`
	Owner        *string  `json:"owner"`
	ProfileImage *string  `json:"profile_image"`
	State        string   `json:"state"`
	Street       string   `json:"street"`
	Type         *string  `json:"type"`
	Zip          string   `json:"zip"`
}

func (e CreateRequestDTO) Validate() error {
	if e.City == "" {
		return fmt.Errorf("field 'city' is required")
	}
	if err := utils.ValidateMinLength(e.City, 1); err != nil {
		return fmt.Errorf("field 'city': %w", err)
	}
	if e.Country == "" {
		return fmt.Errorf("field 'country' is required")
	}
	if err := utils.ValidateMinLength(e.Country, 1); err != nil {
		return fmt.Errorf("field 'country': %w", err)
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
	if e.Neighborhood == "" {
		return fmt.Errorf("field 'neighborhood' is required")
	}
	if err := utils.ValidateMinLength(e.Neighborhood, 1); err != nil {
		return fmt.Errorf("field 'neighborhood': %w", err)
	}
	if e.Number == "" {
		return fmt.Errorf("field 'number' is required")
	}
	if err := utils.ValidateMinLength(e.Number, 1); err != nil {
		return fmt.Errorf("field 'number': %w", err)
	}
	if e.State == "" {
		return fmt.Errorf("field 'state' is required")
	}
	if err := utils.ValidateMinLength(e.State, 2); err != nil {
		return fmt.Errorf("field 'state': %w", err)
	}
	if err := utils.ValidateMaxLength(e.State, 2); err != nil {
		return fmt.Errorf("field 'state': %w", err)
	}
	if e.Street == "" {
		return fmt.Errorf("field 'street' is required")
	}
	if err := utils.ValidateMinLength(e.Street, 1); err != nil {
		return fmt.Errorf("field 'street': %w", err)
	}
	if e.Zip == "" {
		return fmt.Errorf("field 'zip' is required")
	}
	if err := utils.ValidateMinLength(e.Zip, 1); err != nil {
		return fmt.Errorf("field 'zip': %w", err)
	}

	return nil
}
