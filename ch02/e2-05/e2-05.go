package main

import (
	"fmt"

	"github.com/zj27/go-book-exercises/ch02/popcount"
)

func main() {
	fmt.Println(popcount.PopCount25(0))
	fmt.Println(popcount.PopCount25(1))
	fmt.Println(popcount.PopCount25(2))
	fmt.Println(popcount.PopCount25(3))
	fmt.Println(popcount.PopCount25(4))
	fmt.Println(popcount.PopCount25(7))
}
