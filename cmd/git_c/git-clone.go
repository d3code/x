package git_c

import (
    "fmt"
    "github.com/d3code/x/internal/git"
    "github.com/d3code/x/internal/terminal"
    "github.com/spf13/cobra"
    "os"
    "os/exec"
)

func init() {

    Git.AddCommand(cloneCmd)
    cloneCmd.Flags().StringP("dir", "d", "", "change parent directory")
}

var cloneCmd = &cobra.Command{
    Use: "clone",

    Run: func(cmd *cobra.Command, args []string) {
        if git.IsGitDirectory(".") {
            fmt.Println("Current directory is already a git repository")
            return
        }

        directory := cmd.Flag("dir").Value.String()
        if ChangeDirectory(directory) {
            fmt.Println("Clone to directory: " + directory)
        }

        repository := getRepository(args)
        url := git.FormatRepositoryUrl(repository)

        command := exec.Command("git", "clone", url)
        command.Stdout = cmd.OutOrStdout()
        command.Stderr = cmd.ErrOrStderr()

        err := command.Run()
        ExitIfError(err, "Error cloning repository")
    },
}

func getRepository(args []string) string {
    if len(args) > 0 {
        return args[0]
    }

    return terminal.PromptString("Repository to clone", true)
}

func ChangeDirectory(directory string) bool {
    if directory == "" {
        return false
    }

    err := os.Chdir(directory)
    if err != nil {
        fmt.Println(err)
        return false
    }

    return true
}
