package fetch

import (
	"hoge.com/hoge/model"
	"hoge.com/hoge/repository"
)

type Users interface {
	FetchUsers() ([]model.User, error)
}

type usersUsecase struct {
	usersRepository repository.User
}

func (u *usersUsecase) FetchUsers() ([]model.User, error) {
	users, err := u.usersRepository.Fetch()
	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}
