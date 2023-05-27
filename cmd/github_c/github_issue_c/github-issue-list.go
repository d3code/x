package github_issue_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/x/internal/github"
    "github.com/spf13/cobra"
)

func init() {
    Root.AddCommand(List)
}

var List = &cobra.Command{
    Use:     "list",
    Aliases: []string{"l"},
    Run: func(cmd *cobra.Command, args []string) {

        issues := github.ListIssue()

        for _, list := range issues {
            clog.Infof("[ %v ] %s", list.Id, list.Title)
        }

    },
}
