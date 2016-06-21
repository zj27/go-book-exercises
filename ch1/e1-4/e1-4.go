package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]string)
	files := os.Args[1:]
	if len(files) != 0 {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup4: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, names := range counts {
		fmt.Printf("%s\t%s\n", line, names)
	}
}

func countLines(f *os.File, counts map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] += (" " + f.Name())
	}
}
