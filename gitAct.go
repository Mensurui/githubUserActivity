package githubuseractivity

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Activity struct {
	Type string `json:"type"`
	Repo Repo   `json:"repo"`
}

type List []Activity

func (l *List) GetUserActivity(url string) (List, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user activity: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("response failed with status code: %d, body: %s", resp.StatusCode, body)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var activities List
	err = json.Unmarshal(body, &activities)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	return activities, nil
}
