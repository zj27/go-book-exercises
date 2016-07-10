package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf(" %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	f := n - n/3*3
	if f == 0 {
		f = 3
	}
	var buf bytes.Buffer
	for i, j := 0, f; i < n; i, j = j, j+3 {
		buf.WriteString(s[i:j])
		if j != n {
			buf.WriteString(",")
		}
	}
	return buf.String()
}
