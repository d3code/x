package github_repo_c

import (
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/internal/github"
    "github.com/d3code/x/internal/input"
    "github.com/manifoldco/promptui"
    "github.com/spf13/cobra"
    "sort"
    "strings"
)

func init() {

    Root.AddCommand(Task)
}

var Task = &cobra.Command{
    Use:     "task",
    Aliases: []string{"t"},

    Run: func(cmd *cobra.Command, args []string) {
        account := Account()
        repositories := github.Repositories(account)
        sort.Sort(github.RepositoryList(repositories))

        Repo(repositories)
    },
}

func Repo(repos []github.RepoResponse) github.RepoResponse {
    prompt := promptui.Select{
        Label:             "Select repository",
        Items:             repos,
        HideHelp:          true,
        HideSelected:      true,
        StartInSearchMode: true,
        Templates: &promptui.SelectTemplates{
            Label:    "  {{ .FullName }}",
            Active:   "  {{ .FullName | green }}",
            Inactive: "  {{ .FullName }}",
            Details:  "\n{{ .Description }}\n{{ .HtmlUrl }}\n",
            Help:     "",
        },
        Searcher: func(input string, index int) bool {
            repo := repos[index]
            name := repo.Name + "/" + repo.Owner.Login
            if name == "" {
                return false
            }
            return strings.Contains(name, input)
        },
        Stdout: input.NoBellStdout,
    }

    run, _, err := prompt.Run()
    xerr.ExitIfError(err)

    return repos[run]
}
