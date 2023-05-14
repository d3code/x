package git

import (
    _ "embed"
    "github.com/d3code/pkg/xerr"
    "os"
)

//go:embed gitignore.txt
var gitignoreTemplate string

func GitignoreCreate(directory string) {
    err := os.WriteFile(directory+"/.gitignore", []byte(gitignoreTemplate), 0666)
    xerr.ExitIfError(err)
}
