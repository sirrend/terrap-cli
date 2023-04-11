/*
Copyright © 2023 Sirrend
*/

package cmd

import (
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/github_issue"
	"github.com/spf13/cobra"
	"os"
)

// createIssue
/*
@brief:
	createIssue calls the GitHub issue creation method
@params:
	token - string - the GitHub token to use
	title - string - the issue title
	description - string - the problem / feature request title
	isFeature - bool - is a feature request
*/
func createIssue(token, title, description string, isFeature bool) {
	issue, err := github_issue.OpenIssue(token, title, description, isFeature)
	if err != nil {
		if _, ok := err.(*github_issue.RateError); ok {
			_, _ = commons.YELLOW.Println(err.Error())
			os.Exit(0)
		}

		_, _ = commons.RED.Printf("Oops.. something went wrong while trying to create your issue: %s", err.Error())
		os.Exit(0)
	}

	_, _ = commons.SIRREND.Printf("Issue created successfully: https://github.com/sirrend/terrap-cli/issues/%s\n", issue)
}

// openIssueCmd represents the open-issue command
var openIssueCmd = &cobra.Command{
	Use:   "open-issue",
	Short: "An easy way to open a GitHub issue in sirrend/terrap-cli",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("title").Changed && cmd.Flag("description").Changed {
			token := os.Getenv("GH_TOKEN")
			if cmd.Flag("token").Changed {
				createIssue(cmd.Flag("token").Value.String(), cmd.Flag("title").Value.String(), cmd.Flag("description").Value.String(), cmd.Flag("feature").Changed)
			} else if token != "" {
				createIssue(token, cmd.Flag("title").Value.String(), cmd.Flag("description").Value.String(), cmd.Flag("feature").Changed)
			} else {
				_, _ = commons.YELLOW.Println("No GitHub token supplied, please pass it using the environment variable 'GH_TOKEN' or the flag '--token'.")
			}

		} else {
			_, _ = commons.YELLOW.Println("Both Title and Description flags are required.")
		}
	},
}

func init() {
	rootCmd.AddCommand(openIssueCmd)
	openIssueCmd.Flags().StringP("title", "t", "", "The issue title.")
	openIssueCmd.Flags().StringP("description", "d", "", "Issue's description")
	openIssueCmd.Flags().String("token", "", "The token to use to open the issue")
	openIssueCmd.Flags().BoolP("feature", "f", false, "Is this issue a feature request")
}
