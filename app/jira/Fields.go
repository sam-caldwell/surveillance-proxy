package jira

type Fields struct {
	Project     Project `json:"project"`
	Summary     string  `json:"summary"`
	Description string  `json:"description"`
	Issuetype   Type    `json:"issuetype"`
}
