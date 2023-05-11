package git_c

import (
    "fmt"
    "github.com/d3code/pkg/errors"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/internal/git"
    "github.com/d3code/x/internal/terminal"
    "github.com/spf13/cobra"
    "os"
    "os/exec"
    "strings"
)

func init() {
    Git.AddCommand(commitCmd)
    commitCmd.Flags().BoolP("interactive", "i", false, "interactively commit changes    (default: false)")
}

var commitCmd = &cobra.Command{
    Use: "commit",

    Run: func(cmd *cobra.Command, args []string) {
        interactive, err := cmd.Flags().GetBool("interactive")
        errors.ExitIfError(err)

        directory := shell.CurrentDirectory()

        var gitDirs []string
        if !git.IsGitDirectory(directory) {
            gitDirs = getGitSubdirectories(directory)
            for _, dir := range gitDirs {
                shell.Println("[git] {{" + dir + "| grey }}")
            }
        } else {
            gitDirs = append(gitDirs, directory)
        }

        for _, dir := range gitDirs {
            underline := strings.Repeat("-", len(dir)+len("Checking "))
            shell.Println("\nChecking {{ " + dir + " | blue }}")
            shell.Println(underline)

            var commitMessage string
            if len(args) > 0 {
                commitMessage = args[0]
            } else if !interactive {
                commitMessage = "Commit changes"
            }

            Commit(dir, interactive, commitMessage)
        }
    },
}

func Commit(dir string, interactive bool, commitMessage string) {

    shell.RunOut("git", "-C", dir, "status")
    fmt.Println()

    StageUntracked(dir, !interactive)
    StageModified(dir, !interactive)
    CommitChanges(dir, commitMessage)
    Pull(dir, !interactive)
    Push(dir, !interactive)
}

func getGitSubdirectories(currentDirectory string) []string {
    var dirs []string

    dir, err := os.ReadDir(currentDirectory)
    errors.ExitIfError(err)

    for _, d := range dir {
        if d.IsDir() {
            fullDirectory := currentDirectory + "/" + d.Name()
            if git.IsGitDirectory(fullDirectory) {
                dirs = append(dirs, fullDirectory)
            }
        }
    }
    return dirs
}

func StageUntracked(path string, auto bool) {
    untracked := git.Untracked(path)
    if len(untracked) > 0 {

        if auto || terminal.PromptYesNo("Add untracked changes?") {
            fmt.Println("Staging untracked changes")
            for _, file := range untracked {
                fileName := file[3:]

                cmd := exec.Command("git", "-C", path, "add", fileName)
                _, err := cmd.Output()
                errors.ExitIfError(err)

            }
        }
    }
}

func StageModified(path string, auto bool) {
    modified := git.Modified(path)
    if len(modified) > 0 {

        if auto || terminal.PromptYesNo("Stage changes?") {
            fmt.Println("Staging changes")
            for _, file := range modified {
                changed := file[3:]
                var fileName string

                strings.Split(changed, " -> ")
                if len(strings.Split(changed, " -> ")) > 1 {
                    fileName = strings.Split(changed, " -> ")[1]
                } else {
                    fileName = strings.Split(changed, " -> ")[0]
                }

                cmd := exec.Command("git", "-C", path, "add", fileName)
                _, err := cmd.Output()
                errors.ExitIfError(err)
            }
        }
    }
}

func CommitChanges(path string, args string) {
    staged := git.Staged(path)
    if len(staged) > 0 {

        var commitMessage string
        if len(args) > 0 {
            commitMessage = args
        } else {
            commitMessage = terminal.PromptString("Commit message", true)
        }

        shell.Println("Committing changes with commit message {{ " + commitMessage + " | blue }}")
        cmd := exec.Command("git", "-C", shell.FullPath(path), "commit", "-m", commitMessage)
        _, err := cmd.Output()
        errors.ExitIfError(err)
    }
}

func Pull(path string, auto bool) {
    cmdFetch := exec.Command("git", "-C", shell.FullPath(path), "fetch")
    _, err := cmdFetch.Output()
    errors.ExitIfError(err)

    if git.RemoteAhead(path) {
        fmt.Println("Pulling changes")
        if auto || terminal.PromptYesNo("Remote branch is ahead. Pull?") {
            shell.RunOut("git", "-C", shell.FullPath(path), "config", "pull.rebase", "true")
            shell.RunOut("git", "-C", shell.FullPath(path), "pull")
        }
    }
}

func Push(path string, auto bool) {
    if git.RemoteBehind(path) {

        fmt.Println("Pushing changes")
        if auto || terminal.PromptYesNo("Push staged changes to remote?") {
            cmd := exec.Command("git", "-C", shell.FullPath(path), "push")
            _, err := cmd.Output()
            errors.ExitIfError(err)
        }
    }
}
