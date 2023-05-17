package github_issue_c

import (
    "github.com/d3code/x/internal/github"
    "github.com/spf13/cobra"
)

func init() {
    Root.AddCommand(Projects)
}

var Projects = &cobra.Command{
    Use:     "projects",
    Aliases: []string{"p"},
    Run: func(cmd *cobra.Command, args []string) {

        github.GetProjects()

    },
}
