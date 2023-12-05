package github

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/spf13/cast"
	"golang.org/x/oauth2"
)

// OpenIssue
/*
@brief:
	OpenIssue open an issue on in Terrap's repository
@params:
	ghToken - string - the GitHub token to use
	title - string - the issue title
	description - string - the problem / feature request title
	isFeatureRequest - bool - is a feature request
@returns:
	string - the new issue number
	error - if exists
*/
func OpenIssue(ghToken, title, description string, isFeatureRequest bool) (string, error) {
	// create new oauth token
	ctx := context.Background()
	sts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghToken},
	)

	// create client
	nc := oauth2.NewClient(ctx, sts)
	client := github.NewClient(nc)

	// append to title
	if isFeatureRequest {
		title = fmt.Sprintf("[FEATURE] - %s", title)
	}

	// create the request payload
	issueRequest := &github.IssueRequest{
		Title: github.String(title),
		Body:  github.String(description),
	}

	// create the issue
	issue, _, err := client.Issues.Create(ctx, commons.GitHubOwner, commons.GitHubRepo, issueRequest)
	if _, ok := err.(*github.RateLimitError); ok {
		return "", createRateError()
	} else if err != nil {
		return "", errors.New("Failed to create issue: " + err.Error())
	}

	return cast.ToString(*issue.Number), nil
}
