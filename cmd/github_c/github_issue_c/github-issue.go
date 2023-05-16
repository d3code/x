package github_issue_c

import (
    "github.com/spf13/cobra"
)

var Root = &cobra.Command{
    Use:     "issue",
    Aliases: []string{"i"},
}
