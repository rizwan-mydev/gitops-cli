package main

import (
	"log"
	"github.com/rizwan-mydev/gitops-cli/cmd"
)

func main() {
	if err := cmd.CreateBranchCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
