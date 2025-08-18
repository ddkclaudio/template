package mapper

import (
	"encoding/json"
	"errors"
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/library/modules/{{ toSnake .Title }}/domain"
	"github.com/library/modules/{{ toSnake .Title }}/dto"
)


func copyStructFields(dst interface{}, src interface{}) error {
	dstVal := reflect.ValueOf(dst)
	if dstVal.Kind() != reflect.Ptr || dstVal.Elem().Kind() != reflect.Struct {
		return errors.New("dst must be pointer to struct")
	}
	dstVal = dstVal.Elem()

	srcVal := reflect.ValueOf(src)
	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}
	if srcVal.Kind() != reflect.Struct {
		return errors.New("src must be struct or pointer to struct")
	}

	srcType := srcVal.Type()
	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		fieldName := srcType.Field(i).Name

		dstField := dstVal.FieldByName(fieldName)
		if !dstField.IsValid() || !dstField.CanSet() {
			continue
		}

		if dstField.Type() == srcField.Type() {
			dstField.Set(srcField)
		}
	}

	return nil
}

func ToCreateDTOFromString(payload string) (*dto.CreateRequestDTO, error) {
	var createDTO dto.CreateRequestDTO
	if err := json.Unmarshal([]byte(payload), &createDTO); err != nil {
		return nil, err
	}

	if err := createDTO.Validate(); err != nil {
		return nil, err
	}

	return &createDTO, nil
}

func ToEntityFromCreateDTO(dto *dto.CreateRequestDTO) *domain.Entity {
	now := time.Now().Unix()
	entity := &domain.Entity{
		Id:        uuid.New().String(),
		CreatedAt: now,
		UpdatedAt: now,
	}
	_ = copyStructFields(entity, dto)
	return entity
}

func ToEntityFromUpdateDTO(updateDTO *dto.UpdateRequestDTO, existingEntity *domain.Entity) *domain.Entity {
	entityVal := reflect.ValueOf(existingEntity).Elem()
	dtoVal := reflect.ValueOf(updateDTO).Elem()

	for i := 0; i < dtoVal.NumField(); i++ {
		dtoField := dtoVal.Field(i)
		fieldName := dtoVal.Type().Field(i).Name

		entityField := entityVal.FieldByName(fieldName)
		if !entityField.IsValid() || !entityField.CanSet() {
			continue
		}

		if dtoField.Kind() == reflect.Ptr && !dtoField.IsNil() {
			val := dtoField.Elem()

			if entityField.Kind() == reflect.Ptr {
				entityField.Set(dtoField)
			} else {
				if val.Type().AssignableTo(entityField.Type()) {
					entityField.Set(val)
				}
			}
		}
	}

	return existingEntity
}

func ToListResponseDTOFromEntityList(entities []*domain.Entity) []*dto.ResponseDTO {
	var dtos []*dto.ResponseDTO
	if entities == nil {
		return dtos
	}
	for _, entity := range entities {
		dtos = append(dtos, ToResponseDTOFromEntity(entity))
	}
	return dtos
}

func ToResponseDTOFromEntity(entity *domain.Entity) *dto.ResponseDTO {
	responseDTO := &dto.ResponseDTO{}
	_ = copyStructFields(responseDTO, entity)
	responseDTO.Id = entity.Id
	return responseDTO
}

func ToUpdateDTOFromString(body string) (*dto.UpdateRequestDTO, error) {
	var updateDTO dto.UpdateRequestDTO
	if err := json.Unmarshal([]byte(body), &updateDTO); err != nil {
		return nil, err
	}

	if err := updateDTO.Validate(); err != nil {
		return nil, err
	}

	return &updateDTO, nil
}
