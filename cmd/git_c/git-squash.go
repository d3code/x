package git_c

import (
    "github.com/d3code/pkg/clog"
    "os"

    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/internal/cobra_util"
    "github.com/d3code/x/internal/git"
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
        shell.RunOut("/bin/bash", "-c", "git reset $(git commit-tree HEAD^{tree} -m 'Initial commit')")

        // Push
        if cobra_util.PromptConfirm("Force push changes to remote?") {
            shell.RunOut("git", "push", "-f")
            cmd.Println("Pushed to remote")
        }
    },
}
