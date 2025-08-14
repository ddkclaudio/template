package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jaqueline/handlers/address"
)

func main() {
	r := gin.Default()

	address.RegisterRoutes(r.Group(address.GetBasePath()))

	r.Run(":8080")
}
