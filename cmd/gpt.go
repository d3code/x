package cmd

import (
    _ "embed"
    "github.com/d3code/x/internal/gpt"
    "github.com/spf13/cobra"
)

func init() {
    Root.AddCommand(GPT)
}

var GPT = &cobra.Command{
    Use: "gpt",
    Run: func(cmd *cobra.Command, args []string) {
        gpt.GenerateCommitMessage(".")
    },
}
