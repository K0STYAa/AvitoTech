package service

import (
	"errors"
	"github.com/K0STYAa/AvitoTech/pkg/repository"
	"math"
)

type OperationService struct {
	repo repository.Operation
}

func NewOperationService(repo repository.Operation) *OperationService{
	return &OperationService{repo: repo}
}

func (s *OperationService) Accrual(userId int, amount float64) (error) {
	if math.Abs(math.Round(amount*100)-(amount*100)) < 0.00000000001 {
		return s.repo.Accrual(userId, int(math.Round(amount*100)))
	}
	return errors.New("invalid amount")
}

func (s *OperationService) WriteDowns(userId int, amount float64) (error) {
	if math.Abs(math.Round(amount*100)-(amount*100)) < 0.00000000001 {
		return s.repo.WriteDowns(userId, int(math.Round(amount*100)))
	}
	return errors.New("invalid amount")
}

func (s *OperationService) Transfer(senderId int, receiverId int, amount float64) (error) {
	if math.Abs(math.Round(amount*100)-(amount*100)) < 0.00000000001 {
		return s.repo.Transfer(senderId, receiverId, int(math.Round(amount*100)))
	}
	return errors.New("invalid amount")
}
