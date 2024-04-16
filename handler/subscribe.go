package handler

import (
	"project/service"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type SubscribeHandler interface {
}

type subscribehandler struct {
	subscribeService service.SubscribeService
}

func NewSubscribeHandler(subscribeSevice service.SubscribeService) SubscribeHandler {
	return subscribehandler{subscribeService: subscribeSevice}
}
