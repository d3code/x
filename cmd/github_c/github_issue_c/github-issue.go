package github_issue_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/files"
    "github.com/d3code/pkg/xerr"
    "github.com/google/uuid"
    "github.com/spf13/cobra"
    "os"
    "os/exec"
    "strings"
)

var Root = &cobra.Command{
    Use:     "issue",
    Aliases: []string{"i"},
    Run: func(cmd *cobra.Command, args []string) {
        uuid := uuid.New().String()
        dir := os.TempDir()
        file := strings.Join([]string{dir, uuid}, "")

        err := files.Save(dir, uuid, []byte("Hello World!"), true)
        xerr.ExitIfError(err)

        command := exec.Command("vim", file)
        command.Stdin = cmd.InOrStdin()
        command.Stdout = cmd.OutOrStdout()
        command.Stderr = cmd.OutOrStderr()

        err = command.Run()
        xerr.ExitIfError(err)

        result, err := os.ReadFile(file)
        xerr.ExitIfError(err)

        clog.Info(file)
        clog.Info(string(result))
    },
}
