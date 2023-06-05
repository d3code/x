package cmd

import (
    _ "embed"
    "github.com/d3code/clog"
    "github.com/d3code/x/internal/embed_text"
    "github.com/spf13/cobra"
)

func init() {
    Root.AddCommand(Version)
}

var Version = &cobra.Command{
    Use: "version",
    Run: func(cmd *cobra.Command, args []string) {
        clog.Info("{{" + embed_text.Version + "| blue }}")
    },
}
