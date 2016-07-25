package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/zj27/go-book-exercises/ch4/github"
)

func main() {
	listMode := flag.Bool("list", false, "List issues")
	readMode := flag.Bool("read", false, "Read a issue")
	issueID := flag.Int("id", -1, "ID of the issue to read")

	flag.Parse()

	user, repo := "zj27", "helloworld"
	if *listMode {
		result, err := github.ListIssues(user, repo)
		if err != nil {
			log.Fatal(err)
		}
		//	fmt.Printf("%d issues:\n", result.TotalCount)

		for _, item := range result {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	} else if *readMode {
		item, err := github.ReadIssue(user, repo, strconv.Itoa(*issueID))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("#%d\n", item.Number)
		fmt.Printf("User:\t%s\n", item.User.Login)
		fmt.Printf("Title:\t%s\n", item.Title)
		fmt.Printf("State:\t%s\n", item.State)
		fmt.Printf("Body:\n%s\n", item.Body)
	} else {
		fmt.Println("Usage:")
		flag.PrintDefaults()
	}

}
