package main

import (
	"fmt"

	"hoge.com/hoge/usecases"
)

func main() {
	fmt.Println(usecases.NewAddUsecase().Add(1, 1))
}
