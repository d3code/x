package terraform_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/pkg/cfg"
    "github.com/d3code/x/pkg/terraform"
    "github.com/spf13/cobra"
)

func init() {
    Terraform.AddCommand(Update)
    Update.Flags().BoolP("all", "a", false, "apply all terraform projects")
}

var Update = &cobra.Command{
    Use:   "apply",
    Short: "Apply terraform project",
    Run: func(cmd *cobra.Command, args []string) {

        if all, err := cmd.Flags().GetBool("all"); err == nil && all {
            terraform.VerifyPaths()

            configuration := cfg.Configuration()
            for path, _ := range configuration.Terraform {
                clog.Info("Applying terraform project {{" + path + "|blue}}")

                shell.RunCmd(path, true, "terraform", "init")
                shell.RunCmd(path, true, "terraform", "apply", "--auto-approve")
            }
        } else {
            clog.Info("To implement")
        }
    },
}
