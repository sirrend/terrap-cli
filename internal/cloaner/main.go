package cloaner

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"os"
	"strings"
)

// formatTagRef
/*
@brief:
	formatTagRef formats a given tag in a ref format to fit GitHub cloning
@params:
	tag string - the tag to format
@returns:
	string - tag in ref format
*/
func formatTagRef(tag string) string {
	if strings.Contains(tag, "v") {
		return fmt.Sprintf("refs/tags/%s", tag)
	}

	return fmt.Sprintf("refs/tags/v%s", tag)
}

// CloneRepoByTag
/*
@brief:
	CloneRepoByTag clones a repo by a given tag
@params:
	repo string - repo URL to clone
	tag string - the tag to clone
@returns:
	string - the temporary location
	error - if exists, otherwise nil
*/
func CloneRepoByTag(repoURL, tag string) (string, error) {
	tempLocation, _ := os.MkdirTemp("", "*-"+tag)
	if _, err := git.PlainClone(tempLocation, false, &git.CloneOptions{
		URL:             repoURL,
		SingleBranch:    true,
		ReferenceName:   plumbing.ReferenceName(formatTagRef(tag)),
		Tags:            git.NoTags,
		InsecureSkipTLS: true,
		Depth:           1,
	}); err == nil {
		return tempLocation, nil
	}

	return "", errors.New("failed to clone repository")
}

// CloneLatest
/*
@brief:
	CloneLatest clones the latest commit in the main branch
@params:
	repo string - repo URL to clone
@returns:
	string - the temporary location
	error - if exists, otherwise nil
*/
func CloneLatest(repoURL string) (string, error) {
	tempLocation, _ := os.MkdirTemp("", "*")
	if _, err := git.PlainClone(tempLocation, false, &git.CloneOptions{
		URL:             repoURL,
		SingleBranch:    true,
		ReferenceName:   "refs/heads/main",
		Tags:            git.NoTags,
		InsecureSkipTLS: true,
		Depth:           1,
	}); err == nil {
		return tempLocation, nil
	} else {
		fmt.Println(err)
	}

	return "", errors.New("failed to clone repository")
}
