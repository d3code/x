package git

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "os"
)

func Commit(path string, commitMessage string) bool {
    status := shell.RunCmd(path, false, "git", "status", "--porcelain")
    if len(status.Stdout) == 0 {
        clog.Debug("Nothing to commit")
        return false
    }

    if len(commitMessage) == 0 {
        clog.Error("No commit message provided")
        os.Exit(1)
    }

    shell.RunCmd(path, false, "git", "commit", "-m", commitMessage)
    return true
}

func CommitHash(path string) shell.CommandResponse {
    commit := shell.RunShell(false, "(cd "+path+";git rev-parse HEAD 2>/dev/null)")
    return commit
}
