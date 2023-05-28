package git_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/internal/git"
    "github.com/d3code/x/internal/gpt"
    "github.com/d3code/x/internal/input"
    "strings"

    "github.com/d3code/pkg/slice_utils"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/cmd/go_c"
    "github.com/d3code/x/internal/cfg"
    "github.com/spf13/cobra"
)

func init() {
    Git.AddCommand(Commit)
    Commit.Flags().BoolP("all", "a", false, "all scanned repositories")
    Commit.Flags().BoolP("interactive", "i", false, "interactive mode")
}

var Commit = &cobra.Command{
    Use:     "commit",
    Aliases: []string{"c"},
    Short:   "Commit and push changes to the remote repository",
    Run: func(cmd *cobra.Command, args []string) {
        all, err := cmd.Flags().GetBool("all")
        xerr.ExitIfError(err)

        interactive, err := cmd.Flags().GetBool("interactive")
        xerr.ExitIfError(err)

        if all {
            repositories := slice_utils.Keys(cfg.Configuration().Git)

            for _, repository := range repositories {
                CommitDirectory(repository, interactive)
            }
        }

        directory := shell.CurrentDirectory()
        UpdateGoProject(directory)
        CommitDirectory(directory, interactive)
    },
}

func CommitDirectory(repository string, interactive bool) {

    clog.UnderlineF("Checking {{ %s | blue }}", repository)

    // is git repo
    if !git.Is(repository) {
        clog.Warn("Not a git repository")
        repositories := slice_utils.Keys(cfg.Configuration().Git)
        inConfig := slice_utils.ContainsString(repositories, repository)

        if interactive && inConfig && input.PromptConfirm("Remove directory from config?") {
            cfg.Configuration().DeleteGitDirectory(repository)
        } else if !interactive {
            cfg.Configuration().DeleteGitDirectory(repository)
        }

        return
    }

    // is clean, next
    commitMessage, changes := gpt.GenerateCommitMessage(repository)
    if !changes {
        clog.Info("No changes detected")
        return
    }

    if interactive {
        clog.Info(commitMessage, "\n")
        if !input.PromptConfirm("Commit changes?") {
            return
        }
    }

    git.StageCommitFetchPullPush(repository, commitMessage)
}

func UpdateGoProject(updatePath string) {

    configuration := cfg.Configuration()
    for path, _ := range configuration.Golang {
        if strings.HasPrefix(updatePath, path) {
            go_c.UpdateGo(path)
        }
    }
}
