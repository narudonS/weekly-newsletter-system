package service

import (
	"project/repository"
)

type SubscribeService interface {
}

type subscribeService struct {
	subscribeRepo repository.SubscribeRepository
}

func NewSubscribeService(subscribeRepository repository.SubscribeRepository) SubscribeService {
	return &subscribeService{subscribeRepo: subscribeRepository}
}
