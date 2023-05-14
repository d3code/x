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
            _, tfErr := shell.RunDirE(directory, "terraform", "validate")
            return tfErr == nil
        }
    }

    return false
}
