package cmd

import (
    _ "embed"
    "github.com/d3code/pkg/shell"
    "github.com/spf13/cobra"
)

func init() {
    RootCmd.AddCommand(version)
}

//go:embed version.txt
var versionString string

var version = &cobra.Command{
    Use: "version",
    Run: func(cmd *cobra.Command, args []string) {
        shell.Println("Version {{" + versionString + "| blue }}")
    },
}
