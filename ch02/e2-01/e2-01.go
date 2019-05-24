package main

import (
	"fmt"

	"github.com/zj27/go-book-exercises/ch02/tempconv"
)

func main() {

	fmt.Println(tempconv.CToK(tempconv.FreezingC))
	var k tempconv.Kelvin = 0
	fmt.Println(tempconv.KToC(k))

}
