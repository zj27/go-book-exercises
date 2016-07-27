package main

import "fmt"

func dedup(strings []string) []string {
	i := 0
	t := ""
	for _, s := range strings {
		if i == 0 || s != t {
			strings[i] = s
			i++
			t = s
		}
	}
	return strings[:i]
}

func main() {
	data := []string{"one", "one", "two", "two", "one", "three"}
	fmt.Printf("%q\n", dedup(data))
}
