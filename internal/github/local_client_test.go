package github

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateBranch(t *testing.T) {
	client := NewInMemoryGitHubClient()

	// Test: Creating a branch in a new repository
	repo := "test-repo"
	branchName := "feature-branch"
	baseBranch := "main"

	// Call CreateBranch
	err := client.CreateBranch(repo, branchName, baseBranch)
	assert.NoError(t, err, "Expected no error when creating branch")

	// Check if the branch was added
	branches, exists := client.Repositories[repo]
	assert.True(t, exists, "Expected repository %s to exist", repo)

	// Verify the branch name is present
	assert.Contains(t, branches, branchName, "Expected branch %s to be created", branchName)
}
