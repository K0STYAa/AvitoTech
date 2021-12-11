package service

import (
	"github.com/K0STYAa/AvitoTech"
	"github.com/K0STYAa/AvitoTech/pkg/repository"
)

type HistoryService struct {
	repo repository.History
}

func NewHistoryService(repo repository.History) *HistoryService{
	return &HistoryService{repo: repo}
}

func (s *HistoryService) GetById(historyId int, sort string, typeSort string, limit string, offset int) ([]AvitoTech.History, error) {
	history, err := s.repo.GetById(historyId, sort, typeSort, limit, offset)
	for index, element := range history {
		history[index].Amount = element.Amount / 100
	}
	return history, err
}
