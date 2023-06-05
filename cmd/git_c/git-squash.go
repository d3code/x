package git_c

import (
    "github.com/d3code/clog"
    "os"

    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/internal/git"
    "github.com/d3code/x/internal/prompt"
    "github.com/spf13/cobra"
)

func init() {
    Git.AddCommand(squashCmd)
}

var squashCmd = &cobra.Command{
    Use:   "squash",
    Short: "Squash all commits into one",
    PreRun: func(cmd *cobra.Command, args []string) {
        if !git.Git(".") {
            clog.ErrorF("%s is not a git repository", shell.CurrentDirectory())
            os.Exit(1)
        }
    },
    Run: func(cmd *cobra.Command, args []string) {
        shell.RunShell(true, "git reset $(git commit-tree HEAD^{tree} -m 'Initial commit')")

        // Push
        if prompt.Confirm("Force push changes to remote?") {
            shell.RunCmd(".", true, "git", "push", "-f")
            cmd.Println("Pushed to remote")
        }
    },
}
