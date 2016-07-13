package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var m = flag.String("s", "256", "sha")

func main() {
	flag.Parse()
	var s []byte
	fmt.Scanln(&s)
	var c []byte
	if *m == "256" {
		c = sha256.Sum256(s)
	} else if *m == "384" {
		c = sha512.Sum384(s)
	} else if *m == "512" {
		c = sha512.Sum512(s)
	}
	fmt.Printf("%x\n", c)
}
