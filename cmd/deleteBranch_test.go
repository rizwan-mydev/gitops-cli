package cmd

import (
	"testing"

	"github.com/rizwan-mydev/gitops-cli/internal/github"
	"github.com/stretchr/testify/assert"
)

func TestDeleteBranchCmd(t *testing.T) {
	client := github.NewInMemoryGitHubClient()

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
