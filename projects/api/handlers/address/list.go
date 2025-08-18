package address

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/library/modules/address/dto"
)

func list(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		respondWithError(c, http.StatusUnauthorized, "missing Authorization header")
		return
	}

	page, err := parseQueryInt(c, "page", 1)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "invalid page parameter")
		return
	}

	pageSize, err := parseQueryInt(c, "page_size", 20)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "invalid page_size parameter")
		return
	}

	onlyMe, err := parseQueryBool(c, "only_me", true)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "invalid only_me parameter")
		return
	}

	filterDTO := dto.FilterRequestDTO{
		OnlyMe:   onlyMe,
		Owner:    c.Query("owner"),
		Page:     page,
		PageSize: pageSize,
	}

	if valid, err := filterDTO.Validate(); !valid || err != nil {
		respondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	filterJSONBytes, err := json.Marshal(filterDTO)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "failed to marshal filter DTO")
		return
	}

	response, err := ctrl.List(string(filterJSONBytes), authHeader)
	if err != nil {
		respondWithApiError(c, err)
		return
	}

	respondWithSuccess(c, http.StatusOK, response)
}
