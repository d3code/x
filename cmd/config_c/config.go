package config_c

import (
    "github.com/spf13/cobra"
)

var Config = &cobra.Command{
    Use:     "config",
    Aliases: []string{"cfg"},
}
