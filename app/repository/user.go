package repository

import "hoge.com/hoge/model"

type User interface {
	Fetch() ([]model.User, error)
}
