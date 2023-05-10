package config_c

import (
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/terminal"
    "github.com/spf13/cobra"
)

func init() {
    Config.AddCommand(GitHub)
    GitHub.Flags().BoolP("set", "s", false, "set config")
}

var GitHub = &cobra.Command{
    Use:     "github",
    Aliases: []string{"gh"},
    Run: func(cmd *cobra.Command, args []string) {
        if terminal.GetFlagBool(cmd, "set") {
            configItems := []string{
                "GitHub accounts",
            }
            index, _ := terminal.PromptSelect("Configure", configItems)
            if index == 0 {
                configurationGitHubAccount()
            }
        } else {
            cfg.ListGit()
        }
    },
}

func configurationGitHubAccount() {
    configuration := cfg.Configuration()
    accounts := getAccountList(configuration)

    // Select account
    index, name := terminal.PromptSelect("Select", accounts)
    if index == 0 {
        name = terminal.PromptString("User or organization name", true)
    } else {
        name = accounts[index]
    }

    configuration.GitHub[name] = cfg.GitHub{
        Organization: terminal.PromptYesNo("Organization"),
        Token:        terminal.PromptString("Token", true),
    }

    configuration.Save()
}

func getAccountList(configuration *cfg.Config) []string {
    var accounts []string
    accounts = append(accounts, "Create new account")
    for name, _ := range configuration.GitHub {
        accounts = append(accounts, name)
    }
    return accounts
}
