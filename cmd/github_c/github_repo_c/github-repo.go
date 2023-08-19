package github_repo_c

import (
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/internal/github"
    "github.com/d3code/x/internal/github/input"
    "github.com/d3code/x/internal/prompt"
    "github.com/spf13/cobra"
)

const (
    openCode   = "Open in Visual Studio Code"
    openGitHub = "Open in GitHub"
)

var Root = &cobra.Command{
    Use: "repo",
    Run: func(cmd *cobra.Command, args []string) {
        account := input.Account()
        repositories := github.Repositories(account)

        var repoMap = make(map[string]github.SimpleRepo)
        var repoNames []string

        for _, repo := range repositories {
            fullRepositoryName := repo.Owner + "/" + repo.Name
            repoMap[fullRepositoryName] = repo
            repoNames = append(repoNames, fullRepositoryName)
        }

        _, selectedRepository := prompt.Select("Select a repository", repoNames)

        items := GetActions(selectedRepository)
        _, selectedAction := prompt.Select("Select an action", items)

        switch selectedAction {
        case openCode:
            shell.RunCmd(".", false, "code", ".")
        case openGitHub:
            shell.RunCmd(".", false, "open", repoMap[selectedRepository].Url)
        }
    },
}

func GetActions(repo string) []string {
    return []string{"Details", openCode, openGitHub, "Delete", "Update", "Commit", "Exit"}
}
