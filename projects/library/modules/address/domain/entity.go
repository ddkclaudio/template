package domain

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/library/utils"
)

type Entity struct {
	BannerImage  *string  `json:"banner_image"`
	City         string   `json:"city"`
	Complement   *string  `json:"complement"`
	Country      string   `json:"country"`
	CreatedAt    *int64   `json:"created_at"`
	Description  *string  `json:"description"`
	Id           string   `json:"id"`
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
	UpdatedAt    *int64   `json:"updated_at"`
	Zip          string   `json:"zip"`
}

func (e Entity) Validate() error {
	if e.BannerImage != nil {
	}
	if e.City == "" {
		return fmt.Errorf("field 'city' is required")
	}
	if err := utils.ValidateMinLength(e.City, 1); err != nil {
		return fmt.Errorf("field 'city': %w", err)
	}
	if e.Complement != nil {
	}
	if e.Country == "" {
		return fmt.Errorf("field 'country' is required")
	}
	if err := utils.ValidateMinLength(e.Country, 1); err != nil {
		return fmt.Errorf("field 'country': %w", err)
	}
	if e.CreatedAt != nil {
	}
	if e.Description != nil {
	}
	if e.Id == uuid.New().String() {
		return fmt.Errorf("field 'id' is required")
	}
	if e.IsActive != nil {
	}
	if e.IsDeleted != nil {
	}
	if e.IsPrimary != nil {
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
	if e.Owner != nil {
	}
	if e.ProfileImage != nil {
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
	if e.Type != nil {
	}
	if e.UpdatedAt != nil {
	}
	if e.Zip == "" {
		return fmt.Errorf("field 'zip' is required")
	}
	if err := utils.ValidateMinLength(e.Zip, 1); err != nil {
		return fmt.Errorf("field 'zip': %w", err)
	}
	if _, err := uuid.Parse(e.Id); err != nil {
		return fmt.Errorf("field 'id' must be a valid UUID")
	}

	return nil
}

func (*Entity) TableName() string {
	return "addresses"
}
