package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo2() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	//fmt.Printf("%.2fs  %s\n", time.Since(start).Seconds(), s)
	fmt.Println(s, time.Since(start).Seconds())
}

func echo3() {
	start := time.Now()
	//fmt.Printf("%.2fs  %s\n", time.Since(start).Seconds(),
	//	strings.Join(os.Args[1:], " "))
	fmt.Println(strings.Join(os.Args[1:], " "), time.Since(start).Seconds())
}

func main() {
	echo2()
	echo3()
}
