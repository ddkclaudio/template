package controller

import (
	"encoding/json"

	"github.com/library/modules/{{ toSnake .Title }}/dto"
	"github.com/library/modules/{{ toSnake .Title }}/mapper"
	"github.com/library/utils"
)

func (c *Controller) List(filter string, authHeader string) ([]*dto.ResponseDTO, error) {
	authUser, apiErr := utils.DecodeAuthUserFromBearerToken(authHeader)
	if apiErr != nil {
		return nil, utils.NewError(apiErr.Message, apiErr.Code)
	}

	var dto dto.FilterRequestDTO
	if err := json.Unmarshal([]byte(filter), &dto); err != nil {
		return nil, utils.NewError(err.Error(), utils.StatusBadRequest)
	}

	if ok, err := dto.Validate(); !ok || err != nil {
		return nil, utils.NewError(err.Error(), utils.StatusBadRequest)
	}

	if !utils.ContainsRole(authUser.Roles, "admin") {
		if dto.Owner != "" && dto.Owner != authUser.Id {
			return nil, utils.NewError("You do not have permission to access another user's data", utils.StatusForbidden)
		}
	}

	if dto.OnlyMe {
		dto.Owner = authUser.Id
	}

	list, err := c.service.List(dto)
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.StatusInternalServerError)
	}

	return mapper.ToListResponseDTOFromEntityList(list), nil
}
