package config_c

import (
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/internal/git"
    "github.com/d3code/x/internal/golang"
    "github.com/spf13/cobra"
)

func init() {
    Config.AddCommand(Scan)
}

var Scan = &cobra.Command{
    Use:   "scan",
    Short: "Scan for git repositories",
    Run: func(cmd *cobra.Command, args []string) {
        shell.Println("{{Scanning for git repositories...|green}}")
        directory := shell.CurrentDirectory()
        git.Scan(directory)
        git.Validate()

        shell.Println("{{Scanning for go projects...|green}}")
        golang.ScanGoDirectory(directory)
    },
}
