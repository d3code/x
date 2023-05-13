package terraform

import (
    "github.com/d3code/pkg/shell"
    "os"
    "strings"
)

func Terraform(directory string) bool {
    path := shell.FullPath(directory)
    dir, err := os.ReadDir(path)
    if err != nil {
        return false
    }

    for _, entry := range dir {
        if strings.HasSuffix(entry.Name(), ".tf") {
            return true
        }
    }

    return false
}
