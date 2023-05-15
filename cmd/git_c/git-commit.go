package git_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/pkg/cfg"
    "github.com/d3code/x/pkg/git"
    "github.com/spf13/cobra"
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
            for path := range configuration.Git {
                clog.UnderlineF("Checking {{ %s | blue }}", path)
                git.StageCommitFetchPullPush(path, commitMessage)
            }
            return
        }

        git.StageCommitFetchPullPush(".", commitMessage)
    },
}
