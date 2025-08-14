package {{ toSnake .Title }}

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/library/modules/{{ toSnake .Title }}/controller"
	"github.com/library/modules/{{ toSnake .Title }}/repository"
	"github.com/library/modules/{{ toSnake .Title }}/service"
	"github.com/library/utils"
)

var (
	ctrl *controller.Controller
	path string = "/{{ pluralize (toSnake .Title) }}"
	rep  *repository.MySQLRepo
	srv  *service.MainService
)

func GetBasePath() string {
	return path
}

func RegisterRoutes(rg *gin.RouterGroup) {
	rg.DELETE("/:id", delete)
	rg.GET("", list)
	rg.GET("/:id", get)
	rg.PATCH("/:id", update)
	rg.POST("", create)
}

func init() {
	rep = repository.NewMySQLRepo()
	srv = service.NewService(rep)
	ctrl = controller.NewController(srv)
}

func parseQueryBool(c *gin.Context, key string, defaultValue bool) (bool, error) {
	valueStr := c.DefaultQuery(key, strconv.FormatBool(defaultValue))
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return false, err
	}
	return value, nil
}

func parseQueryInt(c *gin.Context, key string, defaultValue int) (int, error) {
	valueStr := c.DefaultQuery(key, strconv.Itoa(defaultValue))
	value, err := strconv.Atoi(valueStr)
	if err != nil || value < 1 {
		return 0, err
	}
	return value, nil
}

func respondWithApiError(c *gin.Context, err error) {
	if err != nil {
		if apiErr, ok := err.(*utils.APIError); ok {
			respondWithError(c, int(apiErr.Code), apiErr.Message)
		} else {
			respondWithError(c, http.StatusInternalServerError, "internal server error")
		}
	}
}

func respondWithError(c *gin.Context, status int, msg interface{}) {
	c.JSON(status, gin.H{
		"success": false,
		"message": msg,
	})
}

func respondWithSuccess(c *gin.Context, status int, msg interface{}) {
	c.JSON(status, gin.H{
		"success": true,
		"message": msg,
	})
}
