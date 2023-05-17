package git_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/pkg/slice_utils"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/input"
    "github.com/manifoldco/promptui"
    "github.com/spf13/cobra"
)

func init() {
    Git.AddCommand(repo)
}

const (
    openCode   = "Open in Visual Studio Code"
    openGitHub = "Open in GitHub"
)

var repo = &cobra.Command{
    Use:     "repo",
    Aliases: []string{"r"},
    Short:   "Open a local git repository",
    Run: func(cmd *cobra.Command, args []string) {
        config := cfg.Configuration()

        var directories = slice_utils.Keys(config.Git)
        if len(directories) == 0 {
            clog.Error("No git repositories found in config")
        }

        _, directory := input.PromptSelect("Select a git repository", directories)
        action, dir := ActionOption(directory)
        switch action {
        case openCode:
            shell.RunCmd(dir, false, "code", ".")
        }
    },
}

func ActionOption(localDirectory string) (string, string) {
    items := []string{"Details", openCode, openGitHub, "Delete", "Update", "Commit", "Exit"}
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
