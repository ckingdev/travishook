package travishook

import (
	"encoding/json"
	"fmt"
)

type Payload struct {
	ID             string     `json:"id"`
	Number         string     `json:"number"`
	Status         int        `json:"status"`
	StartedAt      int        `json:"started_at"`
	FinishedAt     int        `json:"finished_at"`
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
	Notifications map[string][]string `json:"notifications"`
}

type Build struct {
	ID             int    `json:"id"`
	RepositoryID   int    `json:"repository_id"`
	Number         string `json:"number"`
	State          string `json:"state"`
	StartedAt      int    `json:"started_at"`
	FinishedAt     int    `json:"finished_at"`
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
	if len(raw) <= len("payload={}") || string(raw[0:len("payload={")]) != "payload={" {
		fmt.Printf("%s\n", raw)
		return nil, fmt.Errorf("Invalid payload format.")
	}
	raw = raw[len("payload="):]
	var p Payload
	if err := json.Unmarshal(raw, &p); err != nil {
		return nil, err
	}
	return &p, nil
}
