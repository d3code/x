package github_c

import (
    "github.com/d3code/pkg/slice_utils"
    "github.com/d3code/x/cmd/github_c/github_config_c"
    "github.com/d3code/x/cmd/github_c/github_issue_c"
    "github.com/d3code/x/cmd/github_c/github_repo_c"
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/prompt"
    "github.com/spf13/cobra"
    "os"
)

func init() {
    GitHub.AddCommand(github_repo_c.Root)
    GitHub.AddCommand(github_config_c.Root)
    GitHub.AddCommand(github_issue_c.Root)
}

var GitHub = &cobra.Command{
    Use:     "github",
    Aliases: []string{"gh"},
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
        ConfigureGitHubUser()
    },
}

func GitHubToken() *cfg.GitHub {
    config := cfg.Configuration()
    ConfigureGitHubUser()

    users := slice_utils.Keys(config.GitHub)
    if len(users) == 0 {
        return nil
    }

    if len(users) == 1 {
        user := config.GitHub[users[0]]
        return &user
    }

    _, user := prompt.Select("Select a GitHub user", users)
    selected := config.GitHub[user]
    return &selected
}

func ConfigureGitHubUser() {
    config := cfg.Configuration()
    if len(config.GitHub) == 0 {

        username := prompt.String("GitHub username", true)

        if token, ok := os.LookupEnv("GITHUB_TOKEN"); ok {
            if prompt.Confirm("Do you want to use the GITHUB_TOKEN environment variable?") {
                config.AddGitHubUser(username, cfg.GitHub{
                    Token: token,
                })
            } else {
                token = prompt.String("GitHub token", true)
                config.AddGitHubUser(username, cfg.GitHub{
                    Token: token,
                })
            }
        }
    }
}
