package controller

import (
	"github.com/library/utils"
)

func (c *Controller) Delete(id string, authHeader string) error {
	authUser, apiErr := utils.DecodeAuthUserFromBearerToken(authHeader)
	if apiErr != nil {
		return utils.NewError(apiErr.Message, apiErr.Code)
	}

	id = utils.SanitizeString(id)
	if id == "" {
		return utils.NewError("id not provided", utils.StatusBadRequest)
	}

	entity, err := c.service.Get(id)
	if err != nil {
		return utils.NewError(err.Error(), utils.StatusInternalServerError)
	}

	if entity == nil {
		return utils.NewError("entity not found", utils.StatusForbidden)
	}

	if !utils.HasPermission(authUser, entity.Owner) {
		return utils.NewError("permission denied", utils.StatusForbidden)
	}

	if err := c.service.Delete(id); err != nil {
		return utils.NewError(err.Error(), utils.StatusInternalServerError)
	}

	return nil
}
