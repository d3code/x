package cmd

import (
    "fmt"
    "github.com/d3code/pkg/shell"
    "github.com/spf13/cobra"
)

func init() {
    RootCmd.AddCommand(config)
}

var config = &cobra.Command{
    Use: "config",
    Run: func(cmd *cobra.Command, args []string) {

        git := shell.Installed("git")
        if git {
            fmt.Println("Git is installed")
        }
    },
}
