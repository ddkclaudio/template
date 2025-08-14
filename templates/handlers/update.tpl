package {{ toSnake .Title }}

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func update(c *gin.Context) {
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

	id := c.Param("id")

	response, apiErr := ctrl.Update(id, requestJSON, authHeader)
	if apiErr != nil {
		respondWithApiError(c, apiErr)
		return
	}

	respondWithSuccess(c, http.StatusOK, response)
}
