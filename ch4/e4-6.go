package main

import (
	"fmt"
	"unicode"
)

func dedupSpace(strings []byte) []byte {
	i := 0

	for _, s := range strings {
		if unicode.IsSpace(rune(s)) {
			if i == 0 || !unicode.IsSpace(rune(strings[i-1])) {
				strings[i] = ' '
				i++
			}
		} else {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	data := []byte{'a', ' ', 'b', ' ', ' ', 'c', '\t', 'd', ' ', ' ', '\t', '\t', 'e'}
	fmt.Printf("%q\n", data)
	fmt.Printf("%q\n", dedupSpace(data))
}
