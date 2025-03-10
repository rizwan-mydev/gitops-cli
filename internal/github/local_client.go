package github

import (
	"fmt"
)

// InMemoryGitHubClient is a stub implementation of GitHubClient for testing.
type InMemoryGitHubClient struct {
	Repositories map[string][]string // Map of repo -> branches
	PullRequests map[string]string   // Map of branch -> PR ID
}

// NewInMemoryGitHubClient creates a new instance of the in-memory client.
func NewInMemoryGitHubClient() *InMemoryGitHubClient {
	return &InMemoryGitHubClient{
		Repositories: make(map[string][]string),
		PullRequests: make(map[string]string),
	}
}

// CreateBranch creates a new branch in the given repository.
func (c *InMemoryGitHubClient) CreateBranch(repo string, branchName string, baseBranch string) error {
	if _, exists := c.Repositories[repo]; !exists {
		c.Repositories[repo] = []string{}
	}
	c.Repositories[repo] = append(c.Repositories[repo], branchName)
	return nil
}

// CreatePullRequest creates a pull request for the given branch.
func (c *InMemoryGitHubClient) CreatePullRequest(repo string, branchName string, title string, description string) (string, error) {
	prID := fmt.Sprintf("%s-pr-%s", repo, branchName)
	c.PullRequests[branchName] = prID
	return prID, nil
}

// DeleteBranch deletes a branch from the given repository.
func (c *InMemoryGitHubClient) DeleteBranch(repo string, branchName string) error {
	branches, exists := c.Repositories[repo]
	if !exists {
		return fmt.Errorf("repository %s not found", repo)
	}

	// Remove the branch
	newBranches := []string{}
	for _, b := range branches {
		if b != branchName {
			newBranches = append(newBranches, b)
		}
	}
	c.Repositories[repo] = newBranches
	delete(c.PullRequests, branchName)
	return nil
}
