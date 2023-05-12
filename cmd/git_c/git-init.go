package git_c

import (
    _ "embed"
    "fmt"
    "github.com/d3code/pkg/errors"
    "github.com/d3code/pkg/files"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/internal/git"
    "github.com/d3code/x/pkg/terminal"
    "github.com/spf13/cobra"
    "os"
    "os/exec"
)

//go:embed gitignore.txt
var gitignoreTemplate string

func init() {

    Git.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
    Use: "init",

    PreRun: func(cmd *cobra.Command, args []string) {
        if git.IsGitDirectory(".") {
            shell.Println("{{ERROR|red}} Current directory is already a git repository")
            os.Exit(1)
        }
    },
    Run: func(cmd *cobra.Command, args []string) {

        initialize(cmd)
        gitignore(cmd)

        command := exec.Command("git", "add", ".")
        _, err := command.Output()
        errors.ExitIfError(err)

        remote(cmd)

        command = exec.Command("git", "branch", "-M", "master")
        _, err = command.Output()
        errors.ExitIfError(err)

        command = exec.Command("git", "commit", "-m", "Initial commit")
        _, err = command.Output()
        errors.ExitIfError(err)

        if terminal.PromptYesNo("Push to remote?") {
            command := exec.Command("git", "push", "-u", "origin", "master", "--force")
            _, err := command.Output()
            errors.ExitIfError(err)
        }
    },
}

func ExitIfError(err error, message string) {
    if err != nil {
        fmt.Println(message, err)
        os.Exit(1)
    }
}

func initialize(cmd *cobra.Command) {
    if files.Exist(".git") {
        cmd.Println("Git repository is already initialized")
        os.Exit(0)
    } else {
        cmd.Println("Initializing Git repository")
        command := exec.Command("git", "init")
        _, err := command.Output()
        errors.ExitIfError(err)
    }
}

func gitignore(cmd *cobra.Command) {
    overwrite, _ := cmd.Flags().GetBool("gitignore")
    if overwrite {
        gitignoreCreate()
        return
    }

    if !files.Exist(".gitignore") {
        gitignoreCreate()
    } else {
        if terminal.PromptYesNo("Overwrite .gitignore with defaults?") {
            gitignoreCreate()
        }
    }
}

func gitignoreCreate() {
    fmt.Println("Creating .gitignore file with defaults")
    gitignore := []byte(gitignoreTemplate)
    err := os.WriteFile(".gitignore", gitignore, 0666)
    ExitIfError(err, "Error creating .gitignore file")
}

func remote(cmd *cobra.Command) {
    repository := terminal.PromptString("Remote github repository (leave empty for none)", false)
    if repository != "" {
        gitUrl := fmt.Sprintf("git@github.com:d3code/%s.git", repository)
        cmd.Println("Setting remote: " + gitUrl)

        command := exec.Command("git", "remote", "add", "origin", gitUrl)
        _, err := command.Output()
        errors.ExitIfError(err)
    }
}
