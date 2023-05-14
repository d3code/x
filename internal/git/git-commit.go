package git

import (
    "fmt"
    "github.com/d3code/pkg/clog"
    "github.com/d3code/pkg/files"
    "github.com/d3code/pkg/shell"
)

func StageCommitFetchPullPush(path string, commitMessage string) {
    path = shell.FullPath(path)

    if !Git(path) {
        clog.Info("{{ Current directory is not a git repository | red }}")
        return
    }

    file := path + "/.gitignore"
    if !files.Exist(file) {
        clog.InfoF("Creating {{ %s | green }} file with defaults at {{ %s | blue }}", " .gitignore", file)
        GitignoreCreate(path)
    }

    Stage(path)
    Commit(path, commitMessage)

    shell.RunDir(path, "git", "fetch", "-n")
    Pull(path)
    Push(path)
}

func Commit(path string, commitMessage string) bool {
    status := shell.RunDir(path, "git", "status", "--porcelain")
    if len(status) == 0 {
        clog.Info("No changes to commit")
        return false
    }

    if len(commitMessage) == 0 {
        fmt.Println("No commit message given")
        return false
    }

    shell.RunOutDir(path, "git", "commit", "-m", commitMessage)
    return true
}

func Pull(path string) {
    shell.RunOutDir(path, "git", "pull", "--rebase")
}

func Push(path string) {
    branch := shell.RunDir(path, "git", "branch", "--show-current")
    shell.RunOutDir(path, "git", "push", "-u", "origin", branch)
}
