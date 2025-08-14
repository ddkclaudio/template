package controller

import (
	"github.com/library/modules/address/dto"
	"github.com/library/modules/address/mapper"
	"github.com/library/utils"
)

func (c *Controller) Update(id, requestBody, authHeader string) (*dto.ResponseDTO, error) {
	authUser, apiErr := utils.DecodeAuthUserFromBearerToken(authHeader)
	if apiErr != nil {
		return nil, utils.NewError(apiErr.Message, apiErr.Code)
	}

	dto, err := mapper.ToUpdateDTOFromString(requestBody)
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.StatusBadRequest)
	}

	entity, err := c.service.Get(id)
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.StatusInternalServerError)
	}

	if !utils.HasPermission(authUser, entity.Owner) {
		return nil, utils.NewError("permission denied", utils.StatusForbidden)
	}

	entity = mapper.ToEntityFromUpdateDTO(dto, entity)

	if !utils.ContainsRole(authUser.Roles, "admin") {
		dto.Owner = &authUser.Id
	}

	if dto.Owner == nil {
		dto.Owner = entity.Owner
	}

	updatedEntity, err := c.service.Update(entity)
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.StatusInternalServerError)
	}

	return mapper.ToResponseDTOFromEntity(updatedEntity), nil
}
