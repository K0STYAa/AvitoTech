package service

import (
	"encoding/json"
	"github.com/K0STYAa/AvitoTech"
	"github.com/K0STYAa/AvitoTech/pkg/repository"
	"io/ioutil"
	"math"
	"net/http"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetById(userId int, currency string) (AvitoTech.Users, error) {
	user, err := s.repo.GetById(userId)
	if currency != "" {
		resp, err2 := http.Get("http://api.exchangeratesapi.io/v1/latest?access_key=590e9cb35a5a8f520f145a17d5eddcb0")
		if err2 != nil {
			return user, err2
		}
		body, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(body)
		var res interface{}
		json.Unmarshal([]byte(bodyString), &res)
		rates1 := res.(map[string]interface{})
		rates2 := rates1["rates"]

		var Currency, RUB float64
		if rec, ok := rates2.(map[string]interface{}); ok {
			for key, val := range rec {
				if key == currency {
					Currency = val.(float64)
				}
				if key == "RUB" {
					RUB = val.(float64)
				}
			}
		}
		if currency != "EUR" {
			user.Balance = math.Floor((user.Balance/RUB*Currency)*100) / 100
		} else {
			user.Balance = math.Floor((user.Balance/RUB)*100) / 100
		}
	} else {
		user.Balance = math.Floor(user.Balance*100) / 100
	}
	return user, err
}
