package github_c

import (
    "github.com/d3code/x/cmd/github_c/github_account_c"
    "github.com/d3code/x/cmd/github_c/github_configure_c"
    "github.com/d3code/x/cmd/github_c/github_repo_c"
    "github.com/spf13/cobra"
)

func init() {
    GitHub.AddCommand(github_repo_c.Repo)
    GitHub.AddCommand(github_account_c.Account)
    GitHub.AddCommand(github_configure_c.Configure)
}

var GitHub = &cobra.Command{
    Use:     "github",
    Aliases: []string{"gh"},
}
