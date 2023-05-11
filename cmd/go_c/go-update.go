package go_c

import (
    "github.com/d3code/pkg/shell"
    "github.com/spf13/cobra"
)

func init() {
    Go.AddCommand(Update)
}

var Update = &cobra.Command{
    Use:   "update",
    Short: "Update go projects",
    Run: func(cmd *cobra.Command, args []string) {
        shell.Println("{{Scanning for go projects...|green}}")
        //directory := shell.CurrentDirectory()

    },
}
