package go_c

import (
    "github.com/d3code/x/pkg/golang"
    "github.com/spf13/cobra"
)

var Go = &cobra.Command{
    Use: "go",
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
        golang.VerifyPaths()
    },
}
