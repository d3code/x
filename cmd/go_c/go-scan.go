package go_c

import (
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/internal/golang"
    "github.com/spf13/cobra"
)

func init() {
    Go.AddCommand(Scan)
}

var Scan = &cobra.Command{
    Use:   "scan",
    Short: "Scan for go projects",
    Run: func(cmd *cobra.Command, args []string) {
        shell.Println("{{Scanning for go projects...|green}}")
        directory := shell.CurrentDirectory()

        golang.ScanGoDirectory(directory)
        golang.RemoveNotGoProject()
    },
}
