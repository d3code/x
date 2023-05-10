package go_c

import (
    "github.com/d3code/pkg/shell"
    "github.com/spf13/cobra"
)

var Go = &cobra.Command{
    Use: "go",
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
        output := shell.Run("which", "go")
        shell.Println("{{" + output + "|grey}}")
    },
}
