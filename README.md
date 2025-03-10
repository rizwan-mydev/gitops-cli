# GitOps CLI

A simple GitOps CLI tool built using Go and Cobra, designed to manage repositories and branches on local GitHub InMemory.

## Project Structure

- **cmd/**: Contains all the commands for the GitOps CLI tool (e.g., listing repositories, creating branches).
- **internal/github/**: Contains the GitHub client interface and in-memory stub used for testing.
- **main.go**: The entry point for the application, where the root command is defined and initialized.

## Features

- List repositories on local in-memory GitHub.
- Create branches in local in-memory GitHub repositories.
- Create PRs
- Support for filtering repositories by name.
- Unit tests using an in-memory GitHub client for isolated testing.

## Setup & Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/gitops-cli.git
   cd gitops-cli

2. **Install dependencies:**
    go mod tidy

3. **Build Application:**
    go build -o gitops-cli


## Project Requirements

**Create Branch:**
    ./gitops-cli create-branch --repo "repo1" --branch "feature-xyz"

**Create PR:**
    ./gitops-cli create-pr --repo my-repo --branch feature-1 --title "New Feature" --desc "Implementing a new feature"

**Delete Branch:**
    ./gitops-cli delete-branch --repo=test-repo --branch=feature-branch

**Clone the repository:**
    ./gitops-cli list-repositories

**Run Test Cases:**
    go test ./cmd