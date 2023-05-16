package github_repo_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/github"
    "github.com/d3code/x/internal/input"
    "github.com/manifoldco/promptui"
    "github.com/spf13/cobra"
    "sort"
    "strings"
)

const (
    openCode   = "Open in Visual Studio Code"
    openGitHub = "Open in GitHub"
)

var Root = &cobra.Command{
    Use: "repo",
    Run: func(cmd *cobra.Command, args []string) {
        account := github.Account()
        repositories := github.Repositories(account)
        sort.Sort(github.RepositoryList(repositories))

        repo := github.Repo(repositories)

        clog.InfoL(
            "{{ Repository | grey }} "+repo.FullName,
            "{{ Clone URL | grey }}  "+repo.SshUrl,
            "{{ URL | grey }}        "+repo.HtmlUrl,
        )

        if repo.Language != nil {
            clog.InfoF("{{ Language | grey }}   {{ %s | green }}", *repo.Language)
        }

        config := cfg.Configuration()
        remoteLocal := normalizeRemote(repo.SshUrl)

        var localDirectory string
        for directory, git := range config.Git {
            remote := normalizeRemote(git.Remote)
            if remote == remoteLocal {
                localDirectory = directory
            }
        }

        if localDirectory != "" {
            clog.InfoF("{{ Local directory | grey }}   {{ %s | green }}", localDirectory)
        }

        action, dir := ActionOption(repo, localDirectory)
        switch action {
        case openCode:
            shell.RunCmd(dir, false, "code", ".")
        case openGitHub:
            shell.RunCmd(dir, false, "open", repo.HtmlUrl)
        }
    },
}

func ActionOption(repo github.RepoResponse, localDirectory string) (string, string) {

    var items []string
    if localDirectory == "" {
        items = []string{"Details", "Clone", "Exit"}
    } else {
        items = []string{"Details", openCode, openGitHub, "Delete", "Update", "Commit", "Exit"}
    }

    prompt := promptui.Select{
        Label:        "Action",
        Items:        items,
        HideHelp:     true,
        HideSelected: true,
        Templates: &promptui.SelectTemplates{
            Label:    "  {{ . }}",
            Active:   "  {{ . | green }}",
            Inactive: "  {{ . }}",
        },
        Stdout: input.NoBellStdout,
    }

    _, action, err := prompt.Run()
    xerr.ExitIfError(err)

    return action, localDirectory
}

func normalizeRemote(remote string) string {
    remote = strings.TrimSuffix(remote, ".git")
    remote = strings.TrimPrefix(remote, "ssh://")
    remote = strings.ReplaceAll(remote, "git@github.com/", "git@github.com:")

    return remote
}
