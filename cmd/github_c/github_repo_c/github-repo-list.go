package github_repo_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/x/internal/github"
    "github.com/d3code/x/internal/github/input"
    "github.com/spf13/cobra"
)

func init() {

    Root.AddCommand(List)
}

var List = &cobra.Command{
    Use: "list",
    Run: func(cmd *cobra.Command, args []string) {
        account := input.Account()
        repositories := github.Repositories(account)

        for _, repo := range repositories {
            clog.InfoF(" {{ %s/%s | blue }}", repo.Owner, repo.Name)
        }
    },
}
