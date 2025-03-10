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

// TestDeleteBranch verifies that a branch is deleted successfully
func TestDeleteBranch(t *testing.T) {
	client := NewInMemoryGitHubClient()

	// Test: Deleting a branch in a repository
	repo := "test-repo"
	branchToDelete := "feature-branch"
	branchToKeep := "main"

	// Initialize repository with branches
	client.Repositories[repo] = []string{branchToKeep, branchToDelete}
	client.PullRequests[branchToDelete] = "test-repo-pr-feature-branch" // Simulating an existing PR

	// Call DeleteBranch
	err := client.DeleteBranch(repo, branchToDelete)
	assert.NoError(t, err, "Expected no error when deleting branch")

	// Check if the branch was deleted
	branches, exists := client.Repositories[repo]
	assert.True(t, exists, "Expected repository %s to exist", repo)

	// Verify the branch is no longer present
	assert.NotContains(t, branches, branchToDelete, "Expected branch %s to be deleted", branchToDelete)

	// Verify the pull request was deleted
	_, prExists := client.PullRequests[branchToDelete]
	assert.False(t, prExists, "Expected pull request for branch %s to be deleted", branchToDelete)

	// Verify the other branch remains
	assert.Contains(t, branches, branchToKeep, "Expected branch %s to remain", branchToKeep)
}
