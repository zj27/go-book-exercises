package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
