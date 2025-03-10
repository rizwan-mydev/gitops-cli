package github

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
