package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Auth struct {
	Username string
	Password string
}

type IssueForWrite struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func ListIssues(user string, repo string) ([]*Issue, error) {
	listIssuesURL := "https://api.github.com/repos/" + user +
		"/" + repo + "/issues"
	resp, err := http.Get(listIssuesURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("List issues failed: %s", resp.Status)
	}

	var result []*Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil

}

func ReadIssue(user string, repo string, id string) (*Issue, error) {
	getIssueURL := "https://api.github.com/repos/" + user +
		"/" + repo + "/issues/" + id
	resp, err := http.Get(getIssueURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Get issue failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func CreateIssue(user string, repo string, auth Auth, issue *IssueForWrite) (*Issue, error) {
	createIssueURL := "https://api.github.com/repos/" + user +
		"/" + repo + "/issues"

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(issue)

	client := &http.Client{}
	req, err := http.NewRequest("POST", createIssueURL, buf)
	if err != nil {
		return nil, fmt.Errorf("Create issue failed: %s", err)
	}

	req.SetBasicAuth(auth.Username, auth.Password)
	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusCreated {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		return nil, fmt.Errorf("Create issue failed: %s\n %s", resp.Status, bodyString)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
