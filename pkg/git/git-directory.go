package git

import (
    "github.com/d3code/pkg/shell"
    "os"
    "os/exec"
    "strings"
)

func Git(directory string) bool {
    path := shell.FullPath(directory)
    if f, err := os.Stat(path + "/.git"); err != nil || !f.IsDir() {
        return false
    }

    command := exec.Command("git", "-C", path, "status", "-s")
    _, err := command.Output()
    if exitErr, ok := err.(*exec.ExitError); ok &&
        strings.Contains(string(exitErr.Stderr), "not a git repository") {
        return false
    } else if err != nil {
        println(err.Error())
        os.Exit(1)
    }

    return true
}
