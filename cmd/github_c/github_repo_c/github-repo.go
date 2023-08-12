package github_repo_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/internal/github"
    "github.com/d3code/x/internal/github/input"
    "github.com/d3code/x/internal/prompt"
    "github.com/manifoldco/promptui"
    "github.com/spf13/cobra"
    "strings"
)

const (
    openCode   = "Open in Visual Studio Code"
    openGitHub = "Open in GitHub"
)

var Root = &cobra.Command{
    Use: "repo",
    Run: func(cmd *cobra.Command, args []string) {
        account := input.Account()
        repo := github.Repositories(account)
        for _, simpleRepo := range repo {
            clog.InfoF("%s/%s", simpleRepo.Owner, simpleRepo.Name)
        }

        //config := cfg.Configuration()
        //var localDirectory string
        //for directory, git := range config.Git {
        //    remote := normalizeRemote(git.Remote)
        //}
        //
        //if localDirectory != "" {
        //    clog.InfoF("{{ Local directory | grey }}   {{ %s | green }}", localDirectory)
        //}
        //
        //action, dir := ActionOption(repo[0], localDirectory)
        //switch action {
        //case openCode:
        //    shell.RunCmd(dir, false, "code", ".")
        //}

    },
}

func ActionOption(repo github.SimpleRepo, localDirectory string) (string, string) {

    var items []string
    if localDirectory == repo.Url {
        items = []string{"Details", "Clone", openGitHub, "Exit"}
    } else {
        items = []string{"Details", openCode, openGitHub, "Delete", "Update", "Commit", "Exit"}
    }

    selectPrompt := promptui.Select{
        Label:        "Action",
        Items:        items,
        HideHelp:     true,
        HideSelected: true,
        Templates: &promptui.SelectTemplates{
            Label:    "  {{ . }}",
            Active:   "  {{ . | green }}",
            Inactive: "  {{ . }}",
        },
        Stdout: prompt.NoBellStdout,
    }

    _, action, err := selectPrompt.Run()
    xerr.ExitIfError(err)

    return action, localDirectory
}

func normalizeRemote(remote string) string {
    remote = strings.TrimSuffix(remote, ".git")
    remote = strings.TrimPrefix(remote, "ssh://")
    remote = strings.ReplaceAll(remote, "git@github.com/", "git@github.com:")

    return remote
}
