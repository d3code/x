package cmd

import (
    _ "embed"

    "github.com/d3code/clog"
    "github.com/d3code/clog/color"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/cmd/git_c"
    "github.com/d3code/x/cmd/github_c"
    "github.com/d3code/x/cmd/go_c"
    "github.com/d3code/x/cmd/terraform_c"
    "github.com/d3code/x/cmd/util_c"
    "github.com/d3code/x/internal/help"
    "github.com/spf13/cobra"
)

func init() {

    Root.AddCommand(util_c.Util)

    Root.AddCommand(git_c.Git)
    Root.AddCommand(github_c.GitHub)

    Root.AddCommand(go_c.Go)
    Root.AddCommand(terraform_c.Terraform)

    Root.PersistentFlags().BoolP("verbose", "v", false, "show additional information about command execution")

    cobra.OnInitialize(func() {
        if verbose, err := Root.PersistentFlags().GetBool("verbose"); err == nil {
            clog.ShowDebugLogs = verbose
        }
    })
}

var Root = &cobra.Command{
    Use:  "x",
    Long: color.Template(help.Root),
}

func Execute() {
    err := Root.Execute()
    xerr.ExitIfError(err)
}
