package cmd

import (
	_ "embed"

	"github.com/d3code/pkg/shell"
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

	if shell.Installed("go") {
		RootCmd.AddCommand(go_c.Go)
	}

	if shell.Installed("git") {
		RootCmd.AddCommand(git_c.Git)
		RootCmd.AddCommand(github_c.GitHub)
	}

	if shell.Installed("terraform") {
		RootCmd.AddCommand(terraform_c.Terraform)
	}
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
