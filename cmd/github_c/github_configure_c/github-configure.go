package github_configure_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/x/pkg/cfg"
    "github.com/d3code/x/pkg/cobra_util"
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
            clog.Info("{{Account:|green}} " + account)
        }
    },
}

func runGitHubConfiguration() {
    username := cobra_util.PromptString("GitHub username", true)
    token := cobra_util.PromptString("GitHub token", true)

    configuration := cfg.Configuration()
    configuration.SetGitHubUser(username, cfg.GitHub{
        Token: token,
    })
    configuration.Save()
}
