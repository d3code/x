package git

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/files"
    "github.com/d3code/pkg/shell"
    "os"
    "strings"
)

func StageCommitFetchPullPush(path string, commitMessage string) {
    path = shell.FullPath(path)
    clog.Debug("Path:", path)

    if !Git(path) {
        clog.Error(path, "is not a git repository")
        return
    }

    file := path + "/.gitignore"
    if !files.Exist(file) {
        clog.InfoF("Creating {{ .gitignore | green }} with defaults at {{ %s | blue }}", file)
        GitignoreCreate(path)
    }

    Stage(path)
    Commit(path, commitMessage)

    shell.RunCmd(path, false, "git", "fetch", "-n")

    Pull(path)
    Push(path)
}

func Commit(path string, commitMessage string) bool {
    status := shell.RunCmd(path, false, "git", "status", "--porcelain")
    if len(status.Stdout) == 0 {
        clog.Debug("Nothing to commit")
        return false
    }

    if len(commitMessage) == 0 {
        clog.Debug("No commit message given")
        return false
    }

    shell.RunCmd(path, false, "git", "commit", "-m", commitMessage)
    return true
}

func Pull(path string) {
    e, err := shell.RunCmdE(path, false, "git", "pull", "--ff-only")

    if strings.Contains(e.Stderr, "no such ref was fetched") {
        clog.Debug("Branch does not exist on remote")
    } else if err != nil {
        clog.Error(e.Stderr)
        os.Exit(1)
    }
}

func Push(path string) {
    branch := shell.RunCmd(path, false, "git", "branch", "--show-current")
    shell.RunCmd(path, true, "git", "push", "-u", "origin", branch.Stdout)
}

func Stage(path string) (bool, string) {
    status := shell.RunCmd(path, false, "git", "status", "--porcelain")
    if len(status.Stdout) == 0 {
        clog.Debug("No unstaged changes")
        return false, status.Stdout
    }

    shell.RunCmd(path, false, "git", "add", ".")
    return true, status.Stdout
}
