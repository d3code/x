package cmd

import (
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/git"
    "github.com/d3code/x/internal/golang"
    "github.com/spf13/cobra"
)

func init() {
    Root.AddCommand(Scan)
}

var Scan = &cobra.Command{
    Use: "scan",
    Run: func(cmd *cobra.Command, args []string) {

        baseDirectory := cfg.BaseDirectory()

        git.Scan(baseDirectory)
        git.VerifyPaths()

        golang.Scan(baseDirectory)
        golang.VerifyPaths()
    },
}
