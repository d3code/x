package git_c

import (
    "github.com/d3code/clog"
    "os"

    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/pkg/cobra_util"
    "github.com/d3code/x/pkg/git"
    "github.com/spf13/cobra"
)

func init() {

    Git.AddCommand(squashCmd)
}

var squashCmd = &cobra.Command{
    Use: "squash",
    PreRun: func(cmd *cobra.Command, args []string) {
        if !git.Git(".") {
            clog.Info("{{ERROR|red}} Current directory is not a git repository")
            os.Exit(1)
        }
    },
    Run: func(cmd *cobra.Command, args []string) {
        shell.RunShell(true, "git reset $(git commit-tree HEAD^{tree} -m 'Initial commit')")

        // Push
        if cobra_util.PromptConfirm("Force push changes to remote?") {
            shell.RunCmd(".", true, "git", "push", "-f")
            cmd.Println("Pushed to remote")
        }
    },
}
