package address

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		respondWithError(c, http.StatusUnauthorized, "missing Authorization header")
		return
	}

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "failed to read request body")
		return
	}
	requestJSON := string(bodyBytes)

	response, err := ctrl.Create(requestJSON, authHeader)
	if err != nil {
		respondWithApiError(c, err)
		return
	}

	respondWithSuccess(c, http.StatusCreated, response)
}
