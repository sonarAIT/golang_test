package usecases

import "hoge.com/hoge/domain/static"

type Add interface {
	Add(int, int) int
}

type add struct{}

func NewAddUsecase() Add {
	return &add{}
}

func (add *add) Add(n1 int, n2 int) int {
	return static.AddFunc(n1, n2)
}
