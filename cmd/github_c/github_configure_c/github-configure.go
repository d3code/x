package github_configure_c

import (
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/pkg/terminal"
    "github.com/spf13/cobra"
)

func init() {

}

var Configure = &cobra.Command{
    Use:     "configure",
    Aliases: []string{"config", "conf", "cfg", "c"},
    Run: func(cmd *cobra.Command, args []string) {
        configuration := cfg.Configuration()
        if len(configuration.GitHub) == 0 {
            runGitHubConfiguration()
            return
        }

        for account, _ := range configuration.GitHub {
            shell.Println("{{Account:|green}} " + account)
        }
    },
}

func runGitHubConfiguration() {
    username := terminal.PromptString("GitHub username", true)
    token := terminal.PromptString("GitHub token", true)

    configuration := cfg.Configuration()
    configuration.SetGitHubUser(username, cfg.GitHub{
        Token: token,
    })
    configuration.Save()
}
