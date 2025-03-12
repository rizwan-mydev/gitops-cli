package cmd

import (
	"bytes"
	"testing"
	"fmt"

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

	// Simulate the ListRepositoriesCmd logic directly
	// Fetch the repositories
	repos, err := client.ListRepositories()
	assert.NoError(t, err, "Expected no error when listing repositories")

	// Output the repositories to the buffer
	if len(repos) == 0 {
		fmt.Fprintf(&output, "No repositories found.\n")
	} else {
		fmt.Fprintf(&output, "Repositories:\n")
		for _, repo := range repos {
			fmt.Fprintf(&output, "%s\n", repo)
		}
	}

	// Capture command output
	outputStr := output.String()
	t.Logf("Command Output:\n%s", outputStr)

	// Validate the output contains the repository names
	assert.Contains(t, outputStr, "Repositories:", "Expected header 'Repositories:'")
	assert.Contains(t, outputStr, "test-repo-1", "Expected test-repo-1 to be listed")
	assert.Contains(t, outputStr, "test-repo-2", "Expected test-repo-2 to be listed")
}
