package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseUTF8(s []byte) {
	for i, j := 0, 1; j <= len(s); j++ {
		if utf8.Valid(s[i:j]) {
			reverse(s[i:j])
			i = j
		}
	}
	reverse(s)
}

func main() {
	data := []byte{'a', 'b', 'c'}
	reverseUTF8(data[:])
	fmt.Printf("%q\n", data)
	s := "一二三"
	d := []byte(s)
	reverseUTF8(d)
	fmt.Printf("%q\n", d)
}
