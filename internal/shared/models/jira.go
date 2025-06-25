package models

import "fmt"

type JiraIssue struct {
	Key     string
	Project string
	Title   string
	BaseUrl string
}

func (j *JiraIssue) GetUrl() string {
	return fmt.Sprintf("%s/browse/%s", j.BaseUrl, j.Key)
}
