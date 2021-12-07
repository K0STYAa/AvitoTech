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

func (s *HistoryService) GetById(historyId int) (AvitoTech.History, error) {
	return s.repo.GetById(historyId)
}
