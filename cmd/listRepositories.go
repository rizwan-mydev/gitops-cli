package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/rizwan-mydev/gitops-cli/internal/github"
	"log"
	"strings"
)

// ListRepositoriesCmd defines the command for listing repositories
var ListRepositoriesCmd = &cobra.Command{
	Use:   "list-repositories",
	Short: "List all repositories in the GitHub client",
	Run: func(cmd *cobra.Command, args []string) {
		client := github.NewInMemoryGitHubClient()

		// Add some test repositories to simulate data
		client.Repositories["test-repo-1"] = []string{"main", "feature-branch"}
		client.Repositories["test-repo-2"] = []string{"main", "dev-branch"}

		// Get filter flag
		filter, _ := cmd.Flags().GetString("filter")

		// Fetch the list of repositories
		repos, err := client.ListRepositories()
		if err != nil {
			log.Fatalf("Error listing repositories: %v", err)
		}

		// Print the list of repositories, applying filter if necessary
		if len(repos) == 0 {
			fmt.Fprintf(cmd.OutOrStdout(), "No repositories found.\n")
		} else {
			fmt.Fprintf(cmd.OutOrStdout(), "Repositories:\n")
			for _, repo := range repos {
				// If filter is set, match exactly with the repository name
				if filter == "" || strings.EqualFold(repo, filter) {
					fmt.Fprintf(cmd.OutOrStdout(), "%s\n", repo)
				}
			}
		}
	},
}

func init() {
	// Add flags to the ListRepositoriesCmd
	ListRepositoriesCmd.Flags().StringP("filter", "f", "", "Filter repositories by name (exact match)")

	// Add the ListRepositoriesCmd to the root command
	rootCmd.AddCommand(ListRepositoriesCmd)
}
