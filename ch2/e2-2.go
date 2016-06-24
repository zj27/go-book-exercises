package main

import (
	"fmt"
	"github.com/tgpl/ch2/tempconv"
	"os"
	"strconv"
)

func main() {
	numbers := os.Args[1:]
	if len(numbers) == 0 {
	} else {
		for _, number := range numbers {
			n, err := strconv.ParseFloat(number, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			ShowConv(n)
		}
	}
}

func ShowConv(n float64) {
	f := tempconv.Fahrenheit(n)
	c := tempconv.Celsius(n)
	k := tempconv.Kelvin(n)
	fmt.Printf("%s = %s = %s\n", f, tempconv.FToC(f), tempconv.FToK(f))
	fmt.Printf("%s = %s = %s\n", c, tempconv.CToF(c), tempconv.CToK(c))
	fmt.Printf("%s = %s = %s\n", k, tempconv.KToC(k), tempconv.KToF(k))
}
