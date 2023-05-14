package go_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/pkg/golang"
    "github.com/spf13/cobra"
)

func init() {
    Go.AddCommand(Scan)
}

var Scan = &cobra.Command{
    Use:   "scan",
    Short: "Scan for go projects",
    Run: func(cmd *cobra.Command, args []string) {
        directory := shell.CurrentDirectory()
        clog.UnderlineF("Scanning {{%s|green}} projects under {{%s|blue}}", "go", directory)
        golang.Scan(directory)
    },
}
