package util_c

import (
    _ "embed"
    "github.com/d3code/clog"
    "github.com/d3code/x/internal/git"
    "github.com/spf13/cobra"
)

func init() {
    Util.AddCommand(GPTCommand)
}

var GPTCommand = &cobra.Command{
    Use: "gpt",
    Run: func(cmd *cobra.Command, args []string) {

        resp := git.ChatGPT(".")
        clog.InfoF(resp)
    },
}
