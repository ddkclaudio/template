package controller

import (
	"github.com/library/modules/address/service"
)

type Controller struct {
	service *service.MainService
}

func NewController(service *service.MainService) *Controller {
	return &Controller{service: service}
}
