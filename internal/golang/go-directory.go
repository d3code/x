package golang

import (
    "github.com/d3code/pkg/shell"
    "os"
)

func Go(directory string) string {
    path := shell.FullPath(directory)
    if _, err := os.Stat(path + "/go.mod"); err != nil {
        return ""
    }

    out := shell.RunCmd(path, false, "go", "list", "-m")
    return out.Stdout
}
