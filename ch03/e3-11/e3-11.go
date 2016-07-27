package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf(" %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	d := strings.Index(s, ".")
	if d != -1 {
		return comma(s[:d]) + s[d:]
	}
	if s[0] == '+' || s[0] == '-' {
		return s[0:1] + comma(s[1:])
	}
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
