package github_repo_c

import (
    "github.com/d3code/x/internal/github"
    "github.com/spf13/cobra"
    "sort"
)

func init() {

    Root.AddCommand(Task)
}

var Task = &cobra.Command{
    Use:     "task",
    Aliases: []string{"t"},

    Run: func(cmd *cobra.Command, args []string) {
        account := github.Account()
        repositories := github.Repositories(account)
        sort.Sort(github.RepositoryList(repositories))

        github.Repo(repositories)
    },
}
