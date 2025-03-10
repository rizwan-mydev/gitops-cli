package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/rizwan-mydev/gitops-cli/internal/github"
)

var deleteRepo, deleteBranch string

// DeleteBranchCmd defines the command for deleting a branch
var DeleteBranchCmd = &cobra.Command{
	Use:   "delete-branch",
	Short: "Delete a branch from a GitHub repository",
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure flags are set
		if deleteRepo == "" || deleteBranch == "" {
			log.Fatalf("Error: Repository and branch name must be provided")
		}

		client := github.NewInMemoryGitHubClient()

		// Initialize the repository with some branches
		client.Repositories[deleteRepo] = []string{"main", "feature-branch"}

		// Check if the branch exists before attempting to delete
		branches, exists := client.Repositories[deleteRepo]
		if !exists {
			log.Fatalf("Error: Repository %s not found", deleteRepo)
		}

		// Ensure the branch exists in the repository before trying to delete it
		var branchExists bool
		for _, branch := range branches {
			if branch == deleteBranch {
				branchExists = true
				break
			}
		}
		if !branchExists {
			log.Fatalf("Error: Branch %s not found in repository %s", deleteBranch, deleteRepo)
		}

		// Call DeleteBranch method
		err := client.DeleteBranch(deleteRepo, deleteBranch)
		if err != nil {
			log.Fatalf("Error deleting branch '%s': %v", deleteBranch, err)
		}

		// Confirm branch deletion
		branches, exists = client.Repositories[deleteRepo]
		if !exists {
			log.Fatalf("Repository %s does not exist", deleteRepo)
		}

		// Print remaining branches
		fmt.Fprintf(cmd.OutOrStdout(), "Remaining branches in repository '%s':\n", deleteRepo)
		for _, branch := range branches {
			fmt.Fprintf(cmd.OutOrStdout(), "%s\n", branch)
		}
	},
}

func init() {
	// Add flags to the command
	DeleteBranchCmd.Flags().StringVarP(&deleteRepo, "repo", "r", "", "Repository name")
	DeleteBranchCmd.Flags().StringVarP(&deleteBranch, "branch", "b", "", "Branch name to delete")
}
