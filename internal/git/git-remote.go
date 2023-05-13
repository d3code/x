package git

import (
    "github.com/d3code/pkg/shell"
    "github.com/d3code/pkg/xerr"
    "os/exec"
    "strings"
)

func Remote(path string) (string, error) {
    command := exec.Command("git", "-C", shell.FullPath(path), "remote", "get-url", "origin")
    output, err := command.Output()
    if exitErr, ok := err.(*exec.ExitError); ok &&
        strings.Contains(string(exitErr.Stderr), "No such remote") {
        return "", nil
    } else if err != nil {
        return "", err
    }

    remote := string(output)
    return strings.TrimSuffix(remote, "\n"), nil
}

func SetRemote(path string, repo string) {
    url := FormatRepositoryUrl(repo)

    command := exec.Command("git", "-C", shell.FullPath(path), "remote", "set-url", "origin", url)
    _, err := command.Output()
    xerr.ExitIfError(err)
}

func RemoteBehind(path string) bool {
    command := exec.Command("git", "-C", shell.FullPath(path), "status")
    output, err := command.Output()
    xerr.ExitIfError(err)

    if strings.Contains(string(output), "branch is ahead") {
        return true
    }

    return false
}

func RemoteAhead(path string) bool {
    command := exec.Command("git", "-C", shell.FullPath(path), "status")
    output, err := command.Output()
    xerr.ExitIfError(err)

    if strings.Contains(string(output), "branch is behind") || strings.Contains(string(output), "use \"git pull\"") {
        return true
    }

    return false
}
