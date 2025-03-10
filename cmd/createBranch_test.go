package cmd

import (
	"bytes"
	"testing"

	"github.com/rizwan-mydev/gitops-cli/internal/github"
	"github.com/stretchr/testify/assert"
)

// TestCreateBranchCmd verifies that multiple branches can be created successfully
func TestCreateBranchCmd(t *testing.T) {
	// Set up an in-memory GitHub client
	client := github.NewInMemoryGitHubClient()
	repoName := "test-repo"
	branchNames := []string{"feature-1", "feature-2", "feature-3"}
	baseBranch := "main"

	// Initialize repository with base branch
	client.Repositories[repoName] = []string{baseBranch}

	// Prepare command execution
	var output bytes.Buffer
	CreateBranchCmd.SetOut(&output)  // Redirect command output

	// Set command flags
	CreateBranchCmd.Flags().Set("repo", repoName)
	CreateBranchCmd.Flags().Set("base", baseBranch)

	// Add all branches to the flag
	for _, branch := range branchNames {
		CreateBranchCmd.Flags().Set("branch", branch)
	}

	// Execute command
	err := CreateBranchCmd.Execute()
	assert.NoError(t, err, "Error executing CreateBranchCmd")

	// Capture command output
	outputStr := output.String()
	t.Logf("Command Output:\n%s", outputStr)

	// Validate each branch appears in output
	for _, branch := range branchNames {
		assert.Contains(t, outputStr, branch, "Expected branch %s in output", branch)
	}
}
