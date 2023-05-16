package github_issue_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/files"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/internal/github"
    "github.com/d3code/x/internal/input"
    "github.com/google/uuid"
    "github.com/spf13/cobra"
    "os"
    "os/exec"
    "strings"
)

func init() {
    Root.AddCommand(Create)
}

var Create = &cobra.Command{
    Use:     "create",
    Aliases: []string{"c"},
    Run: func(cmd *cobra.Command, args []string) {

        title := input.PromptString("Issue title", true)
        issuedDetails := IssueDetails(cmd, title)

        github.CreateIssue(title, issuedDetails, []string{"enhancement"})
        clog.Info(issuedDetails)
    },
}

func IssueDetails(cmd *cobra.Command, title string) string {
    file := CreateTempFile(title)

    command := exec.Command("vim", file)
    command.Stdin = cmd.InOrStdin()
    command.Stdout = cmd.OutOrStdout()
    command.Stderr = cmd.OutOrStderr()

    err := command.Run()
    xerr.ExitIfError(err)

    result, err := os.ReadFile(file)
    xerr.ExitIfError(err)

    return string(result)
}

func CreateTempFile(contents string) string {
    tempId := uuid.New().String()
    tempDie := os.TempDir()
    file := strings.Join([]string{tempDie, tempId}, "")

    err := files.Save(tempDie, tempId, []byte(contents), true)
    xerr.ExitIfError(err)

    return file
}
