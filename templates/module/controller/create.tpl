package controller

import (
	"github.com/library/modules/{{ toSnake .Title }}/dto"
	"github.com/library/modules/{{ toSnake .Title }}/mapper"
	"github.com/library/utils"
)

func (c *Controller) Create(requestBody string, authHeader string) (*dto.ResponseDTO, error) {
	authUser, apiErr := utils.DecodeAuthUserFromBearerToken(authHeader)
	if apiErr != nil {
		return nil, utils.NewError(apiErr.Message, apiErr.Code)
	}

	dto, err := mapper.ToCreateDTOFromString(requestBody)
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.StatusBadRequest)
	}

	if !utils.ContainsRole(authUser.Roles, "admin") || dto.Owner == nil {
		dto.Owner = &authUser.Id
	}

	if err := dto.Validate(); err != nil {
		return nil, utils.NewError(err.Error(), utils.StatusBadRequest)
	}

	entity, err := c.service.Create(dto)
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.StatusInternalServerError)
	}

	response := mapper.ToResponseDTOFromEntity(entity)
	return response, nil
}
