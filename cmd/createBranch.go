package cmd

import (
	"fmt"
	"log"
	"github.com/spf13/cobra"
	"github.com/rizwan-mydev/gitops-cli/internal/github"
)

var repo, branchName, baseBranch string

// CreateBranchCmd defines the command for creating a branch
var CreateBranchCmd = &cobra.Command{
	Use:   "create-branch",
	Short: "Create a branch in a GitHub repository",
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure flags are set
		if repo == "" || branchName == "" {
			log.Fatalf("Error: Repository and branch name must be provided")
		}
		client := github.NewInMemoryGitHubClient()

		// Call CreateBranch method
		err := client.CreateBranch(repo, branchName, baseBranch)
		if err != nil {
			log.Fatalf("Error creating branch: %v", err)
		}

		// Confirm branch creation
		branches, exists := client.Repositories[repo]
		if !exists {
			log.Fatalf("Repository %s does not exist", repo)
		}

		// Print all branches
		fmt.Printf("Branches in repository '%s':\n", repo)
		for _, branch := range branches {
			fmt.Println(branch)
		}
	},
}

func init() {
	// Add flags to the command
	CreateBranchCmd.Flags().StringVarP(&repo, "repo", "r", "", "Repository name")
	CreateBranchCmd.Flags().StringVarP(&branchName, "branch", "b", "", "Branch name")
	CreateBranchCmd.Flags().StringVarP(&baseBranch, "base", "s", "main", "Base branch name")
}
