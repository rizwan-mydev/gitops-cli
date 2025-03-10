package cmd

import (
	"bytes"
	"testing"

	"github.com/rizwan-mydev/gitops-cli/internal/github"
	"github.com/stretchr/testify/assert"
)

func TestCreatePullRequestCmd(t *testing.T) {
	// Set up an in-memory GitHub client
	client := github.NewInMemoryGitHubClient()

	// Define flags and test variables
	repoName := "test-repo"
	branchName := "feature-branch"
	prTitle := "New Feature PR"
	prDescription := "This is a test pull request."

	// Initialize repository with base branch
	client.Repositories[repoName] = []string{baseBranch}
	
	// Prepare command execution with in-memory client
	var output bytes.Buffer
	CreatePullRequestCmd.SetOut(&output) // Redirect command output

	// Set command flags
	CreatePullRequestCmd.Flags().Set("repo", repoName)
	CreatePullRequestCmd.Flags().Set("branch", branchName)
	CreatePullRequestCmd.Flags().Set("title", prTitle)
	CreatePullRequestCmd.Flags().Set("desc", prDescription)

	// Execute the command with the injected client
	err := CreatePullRequestCmd.Execute()
	assert.NoError(t, err, "Error executing CreatePullRequestCmd")

	// Capture command output
	outputStr := output.String()
	t.Logf("Command Output:\n%s", outputStr)

	// Validate PR ID appears in output
	expectedPRID := repoName + "-pr-" + branchName
	assert.Contains(t, outputStr, expectedPRID, "Expected PR ID in output")
}
