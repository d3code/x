package git

import (
    _ "embed"
    "github.com/d3code/pkg/xerr"
    "os"
    "strings"
)

func FormatRepositoryUrl(repository string) string {
    if !strings.HasPrefix(repository, "https://") && !strings.HasPrefix(repository, "git@") {
        return "git@github.com:" + repository + ".git"
    }
    return repository
}

//go:embed gitignore.txt
var gitignoreTemplate string

func GitignoreCreate(directory string) {
    err := os.WriteFile(directory+"/.gitignore", []byte(gitignoreTemplate), 0666)
    xerr.ExitIfError(err)
}
