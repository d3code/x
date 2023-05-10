package config_c

import (
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/terminal"
    "github.com/spf13/cobra"
)

func init() {
    Config.AddCommand(Git)
    Git.Flags().BoolP("set", "s", false, "set config")
}

var Git = &cobra.Command{
    Use:   "git",
    Short: "Configure git settings",
    Run: func(cmd *cobra.Command, args []string) {
        configItems := []string{
            "GitHub accounts",
        }

        if terminal.GetFlagBool(cmd, "set") {
            index, _ := terminal.PromptSelect("Configure", configItems)
            if index == 0 {
                configurationGitHubAccount()
            }
        } else {
            cfg.ListGit()
        }
    },
}
