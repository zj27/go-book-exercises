package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/zj27/go-book-exercises/ch04/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	var cat_1m []*github.Issue
	var cat_1y []*github.Issue
	var cat_more []*github.Issue

	now := time.Now()
	for _, item := range result.Items {
		d := now.Sub(item.CreatedAt)
		if d.Hours() < 24.0*30 {
			cat_1m = append(cat_1m, item)
		} else if d.Hours() < 24.0*365 {
			cat_1y = append(cat_1y, item)
		} else {
			cat_more = append(cat_more, item)
		}
	}

	fmt.Printf("\nWithin a month\n")
	for _, item := range cat_1m {
		fmt.Printf("#%-5d %9.9s %s %.55s\n",
			item.Number, item.User.Login, item.CreatedAt.String(), item.Title)
	}

	fmt.Printf("\nWithin a year\n")
	for _, item := range cat_1y {
		fmt.Printf("#%-5d %9.9s %s %.55s\n",
			item.Number, item.User.Login, item.CreatedAt.String(), item.Title)
	}

	fmt.Printf("\nMore than a year\n")
	for _, item := range cat_more {
		fmt.Printf("#%-5d %9.9s %s %.55s\n",
			item.Number, item.User.Login, item.CreatedAt.String(), item.Title)
	}

}
