package service

import "github.com/K0STYAa/AvitoTech/pkg/repository"

type User interface {

}

type History interface {

}

type Operation interface {

}

type Service struct {
	User
	History
	Operation
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}