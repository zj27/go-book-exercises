package main

import "fmt"

// Rotate s left by k positions
func rotate(s []int, k int) {
	for i := 0; i < len(s)-k; i++ {
		j := i + k
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	rotate(a[:], 3)
	fmt.Println(a)
}
