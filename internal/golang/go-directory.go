package golang

import (
    "github.com/d3code/pkg/shell"
    "os"
)

func GoProject(directory string) string {
    path := shell.FullPath(directory)
    if _, err := os.Stat(path + "/go.mod"); err != nil {
        return ""
    }

    out := shell.RunDir(path, "go", "list", "-m")
    return out
}
