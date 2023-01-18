package usecases

import (
	"hoge.com/hoge/domain/entity"
	"hoge.com/hoge/domain/repository"
)

type Users interface {
	FetchUsers() ([]entity.User, error)
}

type usersUsecase struct {
	usersRepository repository.User
}

func NewUsersUsecase(usersRepository repository.User) Users {
	return &usersUsecase{usersRepository}
}

func (u *usersUsecase) FetchUsers() ([]entity.User, error) {
	users, err := u.usersRepository.Fetch()
	if err != nil {
		return []entity.User{}, err
	}

	return users, nil
}
