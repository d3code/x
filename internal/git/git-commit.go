package git

import (
    "fmt"
    "github.com/d3code/pkg/shell"
)

func StageCommitFetchPullPush(path string, commitMessage string) {

    path = shell.FullPath(path)

    Stage(path)
    CommitChanges(path, commitMessage)

    shell.RunDir(path, "git", "fetch", "-n")
    Pull(path)
    Push(path)
}

func CommitChanges(path string, commitMessage string) bool {
    status := shell.RunDir(path, "git", "status", "--porcelain")
    if len(status) == 0 {
        shell.Println("No changes to commit")
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
    shell.RunDir(path, "git", "pull", "--rebase")
}

func Push(path string) {
    shell.RunDir(path, "git", "push")
}
