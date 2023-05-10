package git

import (
    "github.com/d3code/pkg/errors"
    "github.com/d3code/pkg/shell"
    "os/exec"
    "strings"
)

func UpToDate(path string) bool {
    command := exec.Command("git", "-C", shell.FullPath(path), "status")
    output, err := command.Output()
    errors.ExitIfError(err)

    if strings.Contains(string(output), "branch is up to date") {
        return true
    }

    return false
}

func Untracked(path string) []string {
    command := exec.Command("git", "-C", shell.FullPath(path), "status", "-s")
    output, err := command.Output()
    errors.ExitIfError(err)

    var listing []string
    lines := strings.Split(string(output), "\n")
    for _, line := range lines {
        if strings.HasPrefix(line, "??") {
            listing = append(listing, line)
        }
    }

    return listing
}

func Modified(path string) []string {
    command := exec.Command("git", "-C", shell.FullPath(path), "status", "-s")
    output, err := command.Output()
    errors.ExitIfError(err)

    var listing []string
    lines := strings.Split(string(output), "\n")

    for _, line := range lines {
        if len(line) < 2 {
            continue
        }

        if line[1:2] != " " && line[0:2] != "??" {
            listing = append(listing, line)
        }
    }

    return listing
}

func Staged(path string) []string {
    command := exec.Command("git", "-C", shell.FullPath(path), "status", "-s")
    output, err := command.Output()
    errors.ExitIfError(err)

    var listing []string
    lines := strings.Split(string(output), "\n")

    for _, line := range lines {
        if len(line) < 2 {
            continue
        }

        if line[0:1] != " " && line[1:2] == " " && line[0:2] != "??" {
            listing = append(listing, line)
        }
    }

    return listing
}
