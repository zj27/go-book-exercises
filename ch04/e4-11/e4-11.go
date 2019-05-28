package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/zj27/go-book-exercises/ch04/github"
)

var user = "zj27"
var repo = "helloworld"

func main() {
	listMode := flag.Bool("list", false, "List issues")
	readMode := flag.Bool("read", false, "Read a issue")
	createMode := flag.Bool("create", false, "Create a issue")
	issueID := flag.Int("id", -1, "ID of the issue to read")
	username := flag.String("username", "", "Username for auth")
	token := flag.String("token", "", "Token for auth")

	flag.Parse()

	if *listMode {
		ListIssues()
	} else if *readMode {
		ReadIssue(issueID)
	} else if *createMode {
		CreateIssue(username, token)
	} else {
		fmt.Println("Usage:")
		flag.PrintDefaults()
	}

}

func ListIssues() {
	result, err := github.ListIssues(user, repo)
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

}

func ReadIssue(issueID *int) {
	item, err := github.ReadIssue(user, repo, strconv.Itoa(*issueID))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("#%d\n", item.Number)
	fmt.Printf("Title:\t%s\n", item.Title)
	fmt.Printf("URL:\t%s\n", item.HTMLURL)
	fmt.Printf("User:\t%s\t%s\n", item.User.Login, item.User.HTMLURL)
	fmt.Printf("State:\t%s\n", item.State)
	fmt.Printf("CreateTime:\t%s\n", item.CreatedAt)
	fmt.Printf("Body:\n%s\n", item.Body)

}

func CreateIssue(username *string, token *string) {
	title := "testTitle"
	body := "testBody"
	var newIssue github.IssueForWrite
	newIssue.Title = title
	newIssue.Body = body
	var auth github.Auth
	auth.Username = *username
	auth.Password = *token
	item, err := github.CreateIssue(user, repo, auth, &newIssue)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Issue #%d is created\n", item.Number)
}
