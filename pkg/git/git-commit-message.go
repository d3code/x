package git

import (
    "strings"
)

var (
    messageModified = []string{"🚀 Updated", "🚧 Work in progress"}
    messageRefactor = []string{"📝 Refactor code", "👌 Improve code structure", "⚡ Improve performance"}
    messageRemove   = []string{"🗑️ Remove unused code", "🔥 Remove code or files"}
    messageFix      = []string{"🐛 Fix bugs", "🚑 Fix critical bug", "🔒 Fix security issues", "🚨 Fix linter warnings"}
    messageDocs     = []string{"📝 Add or update documentation", "📚 Add or update documentation"}
    messageTest     = []string{"✅ Add or update tests", "🚨 Fix linter warnings"}
)

func GenerateCommitMessage(status string) string {
    lines := strings.Split(status, "\n")

    var x map[string][]string
    for _, line := range lines {
        x[line[:1]] = append(x[line[:1]], line[3:])
    }

    return ""
}
