package travishook

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Payload struct {
	ID             int     `json:"id"`
	Number         string     `json:"number"`
	Status         int        `json:"status"`
	StartedAt      string        `json:"started_at"`
	FinishedAt     string        `json:"finished_at"`
	StatusMessage  string     `json:"status_message"`
	Commit         string     `json:"commit"`
	Branch         string     `json:"branch"`
	Message        string     `json:"message"`
	CompareURL     string     `json:"compare_url"`
	CommittedAt    string     `json:"committed_at"`
	CommitterName  string     `json:"committer_name"`
	CommitterEmail string     `json:"committer_email"`
	AuthorName     string     `json:"author_name"`
	AuthorEmail    string     `json:"author_email"`
	Type           string     `json:"type"`
	BuildURL       string     `json:"build_url"`
	Repository     Repository `json:"repository"`
	Config         Config     `json:"config"`
	Matrix         []Build    `json:"matrix"`
}

type Repository struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	OwnerName string `json:"owner_name"`
	URL       string `json:"url"`
}

type Config struct {
	Notifications map[string]string `json:"notifications"`
}

type Build struct {
	ID             int    `json:"id"`
	RepositoryID   int    `json:"repository_id"`
	Number         string `json:"number"`
	State          string `json:"state"`
	StartedAt      string    `json:"started_at"`
	FinishedAt     string    `json:"finished_at"`
	Config         Config `json:"config"`
	Status         int    `json:"status"`
	Log            string `json:"log"`
	ParentID       int    `json:"parent_id"`
	Commit         string `json:"commit"`
	Branch         string `json:"branch"`
	Message        string `json:"message"`
	CommittedAt    string `json:"committed_at"`
	CommitterName  string `json:"committer_name"`
	CommitterEmail string `json:"committer_email"`
	AuthorName     string `json:"author_name"`
	AuthorEmail    string `json:"author_email"`
	CompareURL     string `json:"compare_url"`
}

func makePayload(raw []byte) (*Payload, error) {
	unescaped, err := url.QueryUnescape(string(raw))
	if err != nil {
		return nil, err
	}
	if len(unescaped) <= len("payload={}") || string(unescaped[0:len("payload={")]) != "payload={" {
		fmt.Printf("%s\n", unescaped)
		return nil, fmt.Errorf("Invalid payload format.")
	}
	unescaped = unescaped[len("payload="):]
	var p Payload
	if err := json.Unmarshal([]byte(unescaped), &p); err != nil {
		return nil, err
	}
	return &p, nil
}
