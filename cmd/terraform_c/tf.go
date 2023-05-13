package terraform_c

import (
    "github.com/spf13/cobra"
)

var Terraform = &cobra.Command{
    Use:     "terraform",
    Aliases: []string{"tf"},
}
