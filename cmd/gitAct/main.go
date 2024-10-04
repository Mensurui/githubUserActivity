package main

import (
	"flag"
	"fmt"
	"github.com/Mensurui/githubUserActivity"
	"os"
)

func main() {
	url := flag.String("url", " ", "Enter the url of the github user you want to review the activity of.")
	flag.Parse()
	l := &githubuseractivity.List{}

	switch {
	case *url != " ":
		act, err := l.GetUserActivity(*url)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, values := range act {
			fmt.Printf("Activity: %s, RepoName: %s\n", values.Type, values.Repo.Name)
		}

	default:
		fmt.Println("Enter something")

	}
}
