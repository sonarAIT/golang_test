package repository

import "hoge.com/hoge/domain/entity"

type User interface {
	Fetch() ([]entity.User, error)
}
