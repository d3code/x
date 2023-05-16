package git_c

import (
    "fmt"
    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/internal/git"
    "github.com/spf13/cobra"
    "os"
)

func init() {

    Git.AddCommand(remoteCmd)
}

var remoteCmd = &cobra.Command{
    Use: "remote",

    PreRun: func(cmd *cobra.Command, args []string) {
        if !git.Git(".") {
            clog.Info("{{ERROR|red}} Current directory is not a git repository")
            os.Exit(1)
        }
    },
    Run: func(cmd *cobra.Command, args []string) {
        dir := shell.CurrentDirectory()
        if len(args) > 0 {
            git.SetRemote(dir, args[0])
        }

        output, err := git.Remote(dir)
        xerr.ExitIfError(err)
        fmt.Println("Remote repository [", output, "]")
    },
}
