package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/rizwan-mydev/gitops-cli/internal/github"
)

var repo, baseBranch string
var branchNames []string

// CreateBranchCmd defines the command for creating branches
var CreateBranchCmd = &cobra.Command{
	Use:   "create-branch",
	Short: "Create branches in a GitHub repository",
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure flags are set
		if repo == "" || len(branchNames) == 0 {
			log.Fatalf("Error: Repository and at least one branch name must be provided")
		}
		client := github.NewInMemoryGitHubClient()

		// Iterate over branch names and create each branch
		for _, branchName := range branchNames {
			// Call CreateBranch method
			err := client.CreateBranch(repo, branchName, baseBranch)
			if err != nil {
				log.Fatalf("Error creating branch '%s': %v", branchName, err)
			}
		}

		// Confirm branch creation
		branches, exists := client.Repositories[repo]
		if !exists {
			log.Fatalf("Repository %s does not exist", repo)
		}

		// Print all branches
		fmt.Fprintf(cmd.OutOrStdout(), "Branches in repository '%s':\n", repo)
		for _, branch := range branches {
			fmt.Fprintf(cmd.OutOrStdout(), "%s\n", branch)
		}
		
	},
}

func init() {
	// Add flags to the command
	CreateBranchCmd.Flags().StringVarP(&repo, "repo", "r", "", "Repository name")
	CreateBranchCmd.Flags().StringSliceVarP(&branchNames, "branch", "b", []string{}, "List of branch names to create")
	CreateBranchCmd.Flags().StringVarP(&baseBranch, "base", "s", "main", "Base branch name")
}
