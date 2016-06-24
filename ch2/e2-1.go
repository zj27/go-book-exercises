package main

import (
	"fmt"
	"github.com/tgpl/ch2/tempconv"
)

func main() {

	fmt.Println(tempconv.CToK(tempconv.FreezingC))
	var k tempconv.Kelvin = 0
	fmt.Println(tempconv.KToC(k))

}
