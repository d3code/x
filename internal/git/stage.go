package git

import (
    "github.com/d3code/pkg/shell"
)

func Stage(path string) bool {
    status := shell.RunDir(path, "git", "status", "--porcelain")
    if len(status) == 0 {
        shell.Println("No changes to stage")
        return false
    }

    shell.RunOutDir(path, "git", "add", ".")
    return true
}
