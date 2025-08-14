package controller

import (
	"github.com/library/modules/{{ toSnake .Title }}/dto"
	"github.com/library/modules/{{ toSnake .Title }}/mapper"
	"github.com/library/utils"
)

func (c *Controller) Get(id string, authHeader string) (*dto.ResponseDTO, error) {
	authUser, apiErr := utils.DecodeAuthUserFromBearerToken(authHeader)
	if apiErr != nil {
		return nil, utils.NewError(apiErr.Message, apiErr.Code)
	}

	id = utils.SanitizeString(id)
	if id == "" {
		return nil, utils.NewError("id not provided", utils.StatusBadRequest)
	}

	entity, err := c.service.Get(id)
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.StatusInternalServerError)
	}

	if entity == nil {
		return nil, utils.NewError("entity not found", utils.StatusNotFound)
	}

	if !utils.HasPermission(authUser, entity.Owner) {
		return nil, utils.NewError("permission denied", utils.StatusForbidden)
	}

	response := mapper.ToResponseDTOFromEntity(entity)

	return response, nil
}
