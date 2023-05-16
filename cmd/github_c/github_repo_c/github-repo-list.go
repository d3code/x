package github_repo_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/slice_utils"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/github"
    "github.com/d3code/x/internal/input"
    "github.com/manifoldco/promptui"
    "github.com/spf13/cobra"
    "os"
    "sort"
)

func init() {

    Root.AddCommand(List)
}

var List = &cobra.Command{
    Use: "list",

    Run: func(cmd *cobra.Command, args []string) {
        account := Account()
        repositories := github.Repositories(account)
        sort.Sort(github.RepositoryList(repositories))

        for _, repo := range repositories {
            var private string
            if repo.Private {
                private = "[p]"
            } else {
                private = "[ ]"
            }

            clog.InfoF("{{ %s | grey }}  {{ %s/%s | blue }}", private, repo.Owner.Login, repo.Name)
        }
    },
}

func Account() string {
    configuration := cfg.Configuration()
    accounts := configuration.GitHub
    keys := slice_utils.Keys(accounts)

    if len(accounts) == 0 {
        clog.Error("No accounts configured")
        os.Exit(1)
    } else if len(accounts) == 1 {
        return keys[0]
    }

    prompt := promptui.Select{
        Label:        "Account",
        Items:        keys,
        HideHelp:     true,
        HideSelected: true,
        Templates: &promptui.SelectTemplates{
            Label:    "{{ . }}",
            Active:   "{{ . | green }}",
            Inactive: "  {{ . }}",
            Details:  "",
            Help:     "",
        },
        Stdout: input.NoBellStdout,
    }

    run, _, err := prompt.Run()
    xerr.ExitIfError(err)

    clog.InfoF("Selected account: %s", keys[run])
    return keys[run]
}
