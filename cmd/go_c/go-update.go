package go_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/golang"
    "github.com/spf13/cobra"
)

func init() {
    Go.AddCommand(Update)
    Update.Flags().BoolP("all", "a", false, "update all go projects")
}

var Update = &cobra.Command{
    Use:     "update",
    Aliases: []string{"up"},
    Short:   "Update go project",
    Run: func(cmd *cobra.Command, args []string) {
        all, err := cmd.Flags().GetBool("all")
        xerr.ExitIfError(err)

        if all {
            configuration := cfg.Configuration()
            for path, _ := range configuration.Golang {
                clog.Underline("Checking", path)
                golang.UpdateGo(path)
            }
        }

        path := shell.CurrentDirectory()
        golang.UpdateGo(path)
    },
}
