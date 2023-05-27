package github_repo_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/x/internal/github"
    "github.com/spf13/cobra"
    "sort"
)

func init() {

    Root.AddCommand(List)
}

var List = &cobra.Command{
    Use: "list",

    Run: func(cmd *cobra.Command, args []string) {
        account := github.Account()
        repositories := github.Repositories(account)
        sort.Sort(github.RepositoryList(repositories))

        for _, repo := range repositories {
            var private string
            if repo.Private {
                private = "[p]"
            } else {
                private = "[ ]"
            }

            clog.Infof("{{ %s | grey }}  {{ %s/%s | blue }}", private, repo.Owner.Login, repo.Name)
        }
    },
}
