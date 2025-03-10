package github

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// The test checks if the branch is added to the repository and confirms its existence.
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

// TestCreatePullRequest verifies that a pull request is created successfully
func TestCreatePullRequest(t *testing.T) {
	client := NewInMemoryGitHubClient()

	// Define test input
	repo := "test-repo"
	branchName := "feature-branch"
	prTitle := "New Feature PR"
	prDescription := "This is a test pull request."

	// Call CreatePullRequest
	prID, err := client.CreatePullRequest(repo, branchName, prTitle, prDescription)
	assert.NoError(t, err, "Expected no error when creating pull request")

	// Verify the pull request exists in the client's PullRequests map
	storedPRID, exists := client.PullRequests[branchName]
	assert.True(t, exists, "Expected pull request for branch %s to exist", branchName)
	assert.Equal(t, prID, storedPRID, "Expected PR ID to match")
}
