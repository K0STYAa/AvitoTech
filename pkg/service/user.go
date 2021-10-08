package service

import (
	"github.com/K0STYAa/AvitoTech"
	"github.com/K0STYAa/AvitoTech/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService{
	return &UserService{repo: repo}
}

func (s *UserService) GetById(userId int) (AvitoTech.Users, error) {
	return s.repo.GetById(userId)
}
