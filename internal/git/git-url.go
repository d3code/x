package git

import (
    _ "embed"
    "strings"
)

func FormatRepositoryUrl(remote string) string {
    remote = strings.TrimSuffix(remote, ".git")
    remote = strings.TrimPrefix(remote, "ssh://")
    remote = strings.ReplaceAll(remote, "git@github.com/", "git@github.com:")

    return remote
}
