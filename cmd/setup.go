package cmd

import (
    "fmt"
    "github.com/d3code/pkg/shell"
    "github.com/spf13/cobra"
)

func init() {
    Root.AddCommand(Setup)
}

var Setup = &cobra.Command{
    Use: "setup",
    Run: func(cmd *cobra.Command, args []string) {

        git := shell.Installed("git")
        if git {
            fmt.Println("Git is installed")
        }
    },
}
