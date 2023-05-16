package git_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/cmd/go_c"
    "github.com/d3code/x/pkg/cfg"
    "github.com/d3code/x/pkg/git"
    "github.com/spf13/cobra"
    "strings"
)

func init() {
    Git.AddCommand(commitCmd)
    commitCmd.Flags().BoolP("all", "a", false, "commit all repositories")
}

var commitCmd = &cobra.Command{
    Use: "commit",
    Run: func(cmd *cobra.Command, args []string) {
        all, err := cmd.Flags().GetBool("all")
        xerr.ExitIfError(err)

        configuration := cfg.Configuration()

        var commitMessage string
        if len(args) == 1 {
            commitMessage = args[0]
        }

        if all {
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
