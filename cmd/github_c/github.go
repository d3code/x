package github_c

import (
    "github.com/d3code/x/cmd/github_c/github_config_c"
    "github.com/d3code/x/cmd/github_c/github_issue_c"
    "github.com/d3code/x/cmd/github_c/github_repo_c"
    "github.com/spf13/cobra"
)

func init() {
    GitHub.AddCommand(github_repo_c.Root)
    GitHub.AddCommand(github_config_c.Root)
    GitHub.AddCommand(github_issue_c.Root)
}

var GitHub = &cobra.Command{
    Use:     "github",
    Aliases: []string{"gh"},
}
