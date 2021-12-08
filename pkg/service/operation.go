package service

import (
	"github.com/K0STYAa/AvitoTech/pkg/repository"
)

type OperationService struct {
	repo repository.Operation
}

func NewOperationService(repo repository.Operation) *OperationService{
	return &OperationService{repo: repo}
}

func (s *OperationService) Accrual(userId int, amount float64) (error) {
	return s.repo.Accrual(userId, amount)
}

func (s *OperationService) WriteDowns(userId int, amount float64) (error) {
	return s.repo.WriteDowns(userId, amount)
}

func (s *OperationService) Transfer(senderId int, receiver_id int, amount float64) (error) {
	return s.repo.Transfer(senderId, receiver_id, amount)
}
