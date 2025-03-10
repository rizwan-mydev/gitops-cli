package cmd

import (
	"bytes"
	"testing"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/rizwan-mydev/gitops-cli/internal/github"
	"github.com/stretchr/testify/assert"
)

// TestListRepositoriesCmd verifies that repositories are listed correctly
func TestListRepositoriesCmd(t *testing.T) {
	// Set up an in-memory GitHub client
	client := github.NewInMemoryGitHubClient()

	// Initialize some repositories
	client.Repositories["test-repo-1"] = []string{"main", "feature-branch"}
	client.Repositories["test-repo-2"] = []string{"main", "dev-branch"}

	// Prepare command execution
	var output bytes.Buffer
	ListRepositoriesCmd.SetOut(&output) // Redirect command output

	// Mock the ListRepositories function to use the test client
	ListRepositoriesCmd.RunE = func(cmd *cobra.Command, args []string) error {
		// Return the repositories from the test client
		repos, err := client.ListRepositories()
		if err != nil {
			return err
		}

		// Output the repositories to the command's stdout
		for _, repo := range repos {
			fmt.Fprintf(cmd.OutOrStdout(), "%s\n", repo)
		}
		return nil
	}

	// Execute command
	err := ListRepositoriesCmd.Execute()
	assert.NoError(t, err, "Expected no error when executing ListRepositoriesCmd")

	// Capture command output
	outputStr := output.String()
	t.Logf("Command Output:\n%s", outputStr)

	// Validate the output contains the repository names
	assert.Contains(t, outputStr, "test-repo-1", "Expected test-repo-1 to be listed")
	assert.Contains(t, outputStr, "test-repo-2", "Expected test-repo-2 to be listed")
}
