package github_config_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/clog/color"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/pkg/cfg"
    "github.com/d3code/x/pkg/cobra_util"
    "github.com/spf13/cobra"
)

func init() {

}

var Root = &cobra.Command{
    Use:     "config",
    Aliases: []string{"conf", "cfg"},
    Run: func(cmd *cobra.Command, args []string) {
        configuration := cfg.Configuration()
        if configuration.Environment["github_default_clone_directory"] == "" {
            curDir := shell.CurrentDirectory()
            clog.UnderlineF("Set current directory as default clone directory?")
            clog.InfoL("This will set the current directory as the default clone directory for GitHub repositories.",
                "This can be changed later in the configuration file.", color.ColorString(curDir, "blue"), "")

            if cobra_util.PromptConfirm("Set current directory") {

            }
        }

        if len(configuration.GitHub) == 0 {
            GitHubConfiguration()
        }

        for account, _ := range configuration.GitHub {
            clog.Info("{{ Account: | green }} " + account)
        }
    },
}

func GitHubConfiguration() {
    username := cobra_util.PromptString("GitHub username", true)
    token := cobra_util.PromptString("GitHub token", true)

    configuration := cfg.Configuration()
    configuration.SetGitHubUser(username, cfg.GitHub{
        Token: token,
    })
    configuration.Save()
}
