package repository

type User interface {

}

type History interface {

}

type Operation interface {

}

type Repository struct {
	User
	History
	Operation
}

func NewRepository() *Repository {
	return &Repository{}
}