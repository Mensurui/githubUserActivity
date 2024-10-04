package githubuseractivity

import (
	"testing"
)

func TestMain(t *testing.T) {
	url := "https://api.github.com/users/mensurui/events"
	ls := &List{}
	typeAct := "PushEvent"
	nameAct := "Mensurui/taskTracker"

	result, err := ls.GetUserActivity(url)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(result) == 0 {
		t.Fatal("Expected at least one activity, but got none")
	}

	if result[0].Repo.Name != nameAct {
		t.Errorf("Expected: %s, Got: %s", nameAct, result[0].Repo.Name)
	}

	if result[0].Type != typeAct {
		t.Errorf("Expected: %s, Got: %s", typeAct, result[0].Type)
	}
}
