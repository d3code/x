package terraform_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/pkg/terraform"
    "github.com/spf13/cobra"
)

func init() {
    Terraform.AddCommand(Scan)
}

var Scan = &cobra.Command{
    Use:   "scan",
    Short: "Scan for go projects",
    Run: func(cmd *cobra.Command, args []string) {
        clog.Info("{{Scanning for terraform projects...|green}}")
        directory := shell.CurrentDirectory()

        terraform.Scan(directory)
        terraform.VerifyPaths()
    },
}
