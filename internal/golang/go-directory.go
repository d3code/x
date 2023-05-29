package golang

import (
    "github.com/d3code/pkg/shell"
    "os"
)

func Go(directory string) bool {
    path := shell.FullPath(directory)
    if _, err := os.Stat(path + "/go.mod"); err != nil {
        return false
    }

    return true
}
