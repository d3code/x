package git_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/cmd/go_c"
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/git"
    "github.com/spf13/cobra"
    "strings"
)

func init() {
    Git.AddCommand(commitCmd)
    commitCmd.Flags().BoolP("all", "a", false, "commit all local cloned repositories")
}

var commitCmd = &cobra.Command{
    Use:   "commit",
    Short: "Stage, commit, fetch, pull and push changes to a git repository",
    Run: func(cmd *cobra.Command, args []string) {
        all, err := cmd.Flags().GetBool("all")
        xerr.ExitIfError(err)

        var commitMessage string
        if len(args) > 0 {
            if len(args) > 1 {
                clog.Warn("Too many arguments, only the first one will be used for the commit message [", args[0], "]")
            }
            commitMessage = args[0]
        }

        if all {
            configuration := cfg.Configuration()
            for path := range configuration.Git {
                clog.UnderlineF("Checking {{ %s | blue }}", path)
                git.StageCommitFetchPullPush(path, commitMessage)
            }
            return
        }

        directory := shell.CurrentDirectory()
        UpdateGoProject(directory)

        git.StageCommitFetchPullPush(directory, commitMessage)
    },
}

func UpdateGoProject(updatePath string) {

    configuration := cfg.Configuration()
    for path, _ := range configuration.Golang {
        if strings.HasPrefix(updatePath, path) {
            go_c.UpdateGo(path)
        }
    }
}
