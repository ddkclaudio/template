package controller

import (
	"github.com/library/modules/{{ toSnake .Title }}/service"
)

type Controller struct {
	service *service.MainService
}

func NewController(service *service.MainService) *Controller {
	return &Controller{service: service}
}
