package github

import (
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/internal/prompt"
    "github.com/manifoldco/promptui"
    "strings"
)

func Repo(repos []RepoResponse) RepoResponse {
    selectPrompt := promptui.Select{
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
        },
        Searcher: func(input string, index int) bool {
            repo := repos[index]
            name := repo.Name + "/" + repo.Owner.Login
            if name == "" {
                return false
            }
            return strings.Contains(name, input)
        },
        Stdout: prompt.NoBellStdout,
    }

    run, _, err := selectPrompt.Run()
    xerr.ExitIfError(err)

    return repos[run]
}
