package main

import (
	"fmt"

	"hoge.com/hoge/add"
)

func AddFunc(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	fmt.Println(AddFunc(1, 1))
	fmt.Println(add.AddFunc2(1, 1))
}
