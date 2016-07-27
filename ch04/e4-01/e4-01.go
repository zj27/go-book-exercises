package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println(diffCount(generateHashes()))
}

func generateHashes() ([32]byte, [32]byte) {
	return sha256.Sum256([]byte("x")), sha256.Sum256([]byte("X"))
}

func diffCount(c1 [32]byte, c2 [32]byte) int {
	var c3 [32]byte
	for i := 0; i < 32; i++ {
		c3[i] = c1[i] ^ c2[i]
	}
	return popCount(c3)
}

func popCount(x [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		for j := x[i]; j > 0; {
			count++
			j = j & (j - 1)
		}
	}
	return count
}
