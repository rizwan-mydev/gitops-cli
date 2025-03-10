package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/rizwan-mydev/gitops-cli/internal/github"
)

var prRepo, prBranch, prTitle, prDescription string

// CreatePullRequestCmd defines the command for creating a pull request
var CreatePullRequestCmd = &cobra.Command{
	Use:   "create-pr",
	Short: "Create a pull request in a GitHub repository",
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure required flags are set
		if prRepo == "" || prBranch == "" || prTitle == "" {
			log.Fatalf("Error: Repository, branch name, and PR title must be provided")
		}

		client := github.NewInMemoryGitHubClient()

		// Call CreatePullRequest method
		prID, err := client.CreatePullRequest(prRepo, prBranch, prTitle, prDescription)
		if err != nil {
			log.Fatalf("Error creating pull request: %v", err)
		}

		// Print the pull request ID
		fmt.Printf("Pull request created: %s\n", prID)
	},
}

func init() {
	// Add flags to the command
	CreatePullRequestCmd.Flags().StringVarP(&prRepo, "repo", "r", "", "Repository name")
	CreatePullRequestCmd.Flags().StringVarP(&prBranch, "branch", "b", "", "Branch name for PR")
	CreatePullRequestCmd.Flags().StringVarP(&prTitle, "title", "t", "", "Pull request title")
	CreatePullRequestCmd.Flags().StringVarP(&prDescription, "desc", "d", "", "Pull request description")
}
