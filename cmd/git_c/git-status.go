package git_c

import (
    "github.com/d3code/pkg/shell"
    "github.com/d3code/pkg/xerr"
    "github.com/spf13/cobra"
    "os/exec"
)

func init() {

    Git.AddCommand(gitStatus)
}

var gitStatus = &cobra.Command{
    Use:   "status",
    Short: "Show the working tree status",
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
        shell.Installed("git")
    },
    Run: func(cmd *cobra.Command, args []string) {
        command := exec.Command("git", "status")
        command.Stdout = cmd.OutOrStdout()
        command.Stderr = cmd.ErrOrStderr()

        err := command.Run()
        xerr.ExitIfError(err)

    },
}

func ChatGPT() {

}
