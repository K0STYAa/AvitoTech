package service

import (
	"github.com/K0STYAa/AvitoTech"
	"github.com/K0STYAa/AvitoTech/pkg/repository"
)

type User interface {
	GetById(userId int) (AvitoTech.Users, error)
}

type History interface {
	GetById(historyId int) (AvitoTech.History, error)
}

type Operation interface {
	Accrual(userId int, accrual int) (error)
	WriteDowns(userId int, accrual int) (error)
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