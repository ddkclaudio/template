package address

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func get(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		respondWithError(c, http.StatusUnauthorized, "missing Authorization header")
		return
	}

	id := c.Param("id")

	response, err := ctrl.Get(id, authHeader)
	if err != nil {
		respondWithApiError(c, err)
		return
	}

	respondWithSuccess(c, http.StatusOK, response)
}
