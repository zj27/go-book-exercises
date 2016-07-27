package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Not enough args.")
		return
	}
	fmt.Println(isAnagrams(os.Args[1], os.Args[2]))
}

func isAnagrams(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, j := 0, len(s2)-1; i < j; {
		if s1[i] != s2[j] {
			return false
		}
		i++
		j--
	}
	return true
}
