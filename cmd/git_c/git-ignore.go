package git_c

import (
    "github.com/d3code/x/internal/git"
    "github.com/spf13/cobra"
)

func init() {
    Git.AddCommand(Ignore)
}

var Ignore = &cobra.Command{
    Use:   "ignore",
    Short: "Create .gitignore file for the current directory",
    Run: func(cmd *cobra.Command, args []string) {
        git.GitignoreCreate(".")
    },
}
