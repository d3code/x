package git_c

import (
    "fmt"
    "github.com/d3code/pkg/clog"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/git"
    "github.com/spf13/cobra"
    "os"
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
        if len(args) == 1 && args[0] != "" {
            commitMessage = args[0]
        } else {
            commitMessage = "Auto commit"
        }

        if all {
            for path, _ := range configuration.Git {

                msg := fmt.Sprintf("%sChecking {{%s|blue}}", "\n", path)
                clog.Info(msg)
                underline := strings.Repeat("-", len(path)+len("Checking "))
                clog.Info(underline)

                git.StageCommitFetchPullPush(path, commitMessage)
            }
            return
        }

        currentDirectory := shell.CurrentDirectory()
        if !git.Git(currentDirectory) {
            clog.Info("{{ Current directory is not a git repository | red }}")
            os.Exit(1)
        }

        underline := strings.Repeat("-", len(currentDirectory)+len("Checking "))
        clog.Info(underline)

        shell.RunOut("git", "-C", currentDirectory, "status")
        fmt.Println()

        git.StageCommitFetchPullPush(currentDirectory, commitMessage)

    },
}
