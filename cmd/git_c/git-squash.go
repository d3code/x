package git_c

import (
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/internal/git"
    "github.com/d3code/x/internal/terminal"
    "github.com/spf13/cobra"
    "os"
)

func init() {

    Git.AddCommand(squashCmd)
}

var squashCmd = &cobra.Command{
    Use: "squash",

    PreRun: func(cmd *cobra.Command, args []string) {
        if !git.IsGitDirectory(".") {
            shell.Println("{{ERROR|red}} Current directory is not a git repository")
            os.Exit(1)
        }
    },
    Run: func(cmd *cobra.Command, args []string) {
        shell.RunOut("/bin/bash", "-c", "git reset $(git commit-tree HEAD^{tree} -m 'Initial commit')")

        // Push
        if terminal.PromptYesNo("Force push changes to remote?") {
            shell.RunOut("git", "push", "-f")
            cmd.Println("Pushed to remote")
        }
    },
}
