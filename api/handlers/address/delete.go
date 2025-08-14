package address

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func delete(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		respondWithError(c, http.StatusUnauthorized, "missing Authorization header")
		return
	}

	id := c.Param("id")
	if id == "" {
		respondWithError(c, http.StatusBadRequest, "missing id parameter")
		return
	}

	err := ctrl.Delete(id, authHeader)
	if err != nil {
		respondWithApiError(c, err)
		return
	}

	respondWithSuccess(c, http.StatusNoContent, gin.H{
		"message": "record deleted successfully",
		"id":      id,
	})
}
