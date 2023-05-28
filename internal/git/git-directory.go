package git

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "os"
    "os/exec"
    "strings"
)

func Is(directory string) bool {
    path := shell.FullPath(directory)
    if f, err := os.Stat(path + "/.git"); err != nil || !f.IsDir() {
        clog.Debug(path, "no .git directory found")
        return false
    }

    _, err := shell.RunCmdE(path, false, "git", "status", "-s")
    if exitErr, ok := err.(*exec.ExitError); ok && strings.Contains(string(exitErr.Stderr), "not a git repository") {
        clog.Debug(path, string(exitErr.Stderr))
        return false
    } else if err != nil {
        clog.Error(err.Error())
        os.Exit(1)
    }

    return true
}
