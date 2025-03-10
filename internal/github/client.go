package github

// GitHubClient defines the operations for interacting with GitHub.
type GitHubClient interface {
	// CreateBranch creates a new branch in the given repository.
	CreateBranch(repo string, branchName string, baseBranch string) error

	// CreatePullRequest creates a pull request for the given branch.
	CreatePullRequest(repo string, branchName string, title string, description string) (string, error)

	// DeleteBranch deletes a branch from the given repository.
	DeleteBranch(repo string, branchName string) error

	// ListRepositories fetches a list of repositories the user has access to.
	ListRepositories() ([]string, error)
}
