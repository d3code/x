package git

import (
    _ "embed"
    "strings"
)

func FormatRepositoryUrl(repository string) string {
    if !strings.HasPrefix(repository, "https://") && !strings.HasPrefix(repository, "git@") {
        return "git@github.com:" + repository + ".git"
    }
    return repository
}
