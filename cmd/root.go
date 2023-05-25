package cmd

import (
    _ "embed"
    "github.com/d3code/clog"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/cmd/git_c"
    "github.com/d3code/x/cmd/github_c"
    "github.com/d3code/x/cmd/go_c"
    "github.com/d3code/x/cmd/terraform_c"
    "github.com/d3code/x/cmd/util_c"
    "github.com/spf13/cobra"
)

func init() {

    RootCmd.AddCommand(util_c.Util)

    RootCmd.AddCommand(git_c.Git)
    RootCmd.AddCommand(github_c.GitHub)

    RootCmd.AddCommand(go_c.Go)
    RootCmd.AddCommand(terraform_c.Terraform)

    RootCmd.PersistentFlags().BoolP("verbose", "v", false, "show additional information about command execution")

    cobra.OnInitialize(func() {
        if verbose, err := RootCmd.PersistentFlags().GetBool("verbose"); err == nil {
            clog.ShowDebugLogs = verbose
        }
    })
}

//go:embed root.txt
var welcome string

var RootCmd = &cobra.Command{
    Use:  "x",
    Long: welcome,
}

func Execute() {
    err := RootCmd.Execute()
    xerr.ExitIfError(err)
}
