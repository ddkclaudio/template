package main

import (
	"github.com/gin-gonic/gin"
	// "github.com/jaqueline/handlers/settings"
)

func main() {
	r := gin.Default()

	// settings.RegisterRoutes(r.Group(settings.GetBasePath()))

	r.Run(":8080")
}
