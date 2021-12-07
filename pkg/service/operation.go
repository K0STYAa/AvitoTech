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

func (s *OperationService) Accrual(userId int, accrual int) (error) {
	return s.repo.Accrual(userId, accrual)
}

func (s *OperationService) WriteDowns(userId int, accrual int) (error) {
	return s.repo.WriteDowns(userId, accrual)
}
