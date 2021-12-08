package service

import (
	"github.com/K0STYAa/AvitoTech"
	"github.com/K0STYAa/AvitoTech/pkg/repository"
)

type User interface {
	GetById(userId int, currency int) (AvitoTech.Users, error)
}

type History interface {
	GetById(historyId int) (AvitoTech.History, error)
}

type Operation interface {
	Accrual(userId int, amount int) (error)
	WriteDowns(userId int, amount int) (error)
	Transfer(senderId int, receiver_id int, amount int) (error)
}

type Service struct {
	User
	History
	Operation
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
		History: NewHistoryService(repos.History),
		Operation: NewOperationService(repos.Operation),
	}
}