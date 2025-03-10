package main

import (
	"log"
	"github.com/rizwan-mydev/gitops-cli/cmd"
	"github.com/spf13/cobra"
)

func main() {
	// Create a root command to group subcommands
	rootCmd := &cobra.Command{
		Use:   "gitops-cli",
		Short: "GitOps CLI tool for managing branches and pull requests",
	}

	// Add subcommands
	rootCmd.AddCommand(cmd.CreateBranchCmd) // Add the new branch command
	rootCmd.AddCommand(cmd.CreatePullRequestCmd) // Add the new PR command

	// Execute the CLI
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
