package cmd

import (
    _ "embed"
    "github.com/d3code/clog"
    "github.com/d3code/x/internal/help"
    "github.com/spf13/cobra"
)

func init() {
    Root.AddCommand(Version)
}

var Version = &cobra.Command{
    Use: "version",
    Run: func(cmd *cobra.Command, args []string) {
        clog.Info("{{" + help.Version + "| blue }}")
    },
}
