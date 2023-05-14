package git_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/pkg/xerr"
    "github.com/spf13/cobra"
    "os"
    "os/exec"
)

func init() {

    Git.AddCommand(gitStatus)
}

var gitStatus = &cobra.Command{
    Use: "status",
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
        err := shell.RunOutE("which", "git")
        if err != nil {
            clog.Info("{{ ERROR:|red}} git is not installed ")
            os.Exit(1)
        }
    },
    Run: func(cmd *cobra.Command, args []string) {
        command := exec.Command("git", "status")
        command.Stdout = cmd.OutOrStdout()
        command.Stderr = cmd.ErrOrStderr()

        err := command.Run()
        xerr.ExitIfError(err)
    },
}
