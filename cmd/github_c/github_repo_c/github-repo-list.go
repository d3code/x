package github_repo_c

import (
    "encoding/json"
    "fmt"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/pkg/cfg"
    "github.com/d3code/x/pkg/cobra_util"
    "github.com/spf13/cobra"
    "io"
    "net/http"
    "os"
)

func init() {

    Repo.AddCommand(List)
}

var List = &cobra.Command{
    Use: "list",

    Run: func(cmd *cobra.Command, args []string) {
        owner := getAccount(args)

        repoResponse := getRepositories(owner)
        for _, repo := range repoResponse {
            fmt.Println(owner + "/" + repo.Name)
        }
    },
}

func getAccount(args []string) string {
    configuration := cfg.Configuration()

    var account string
    if len(args) > 0 {
        account = args[0]
    } else if len(configuration.GitHub) > 0 {
        var owners []string
        for name, _ := range configuration.GitHub {
            owners = append(owners, name)
        }
        if len(owners) > 0 {
            _, account = cobra_util.PromptSelect("Please select account", owners)
        }
    }

    if account == "" {
        account = cobra_util.PromptString("Please select account", true)
    }

    return account
}

func getRepoUrl(owner string) string {
    configuration := cfg.Configuration()
    if _, ok := configuration.GitHub[owner]; ok {
        return fmt.Sprintf("https://api.github.com/orgs/%s/repos", owner)
    }

    return fmt.Sprintf("https://api.github.com/orgs/%s/repos", owner)
}

func getToken(owner string) string {
    configuration := cfg.Configuration()
    if val, ok := configuration.GitHub[owner]; ok {
        return val.Token
    }

    return os.Getenv("GITHUB_TOKEN")
}

func getRepositories(owner string) []RepoResponse {

    client := &http.Client{}
    url := getRepoUrl(owner)
    req, _ := http.NewRequest("GET", url, nil)

    token := getToken(owner)
    if token != "" {
        auth := fmt.Sprintf("Bearer %s", token)
        req.Header.Set("Authorization", auth)
    }

    res, _ := client.Do(req)
    responseBody, _ := io.ReadAll(res.Body)

    var repoResponse []RepoResponse
    err := json.Unmarshal(responseBody, &repoResponse)
    xerr.ExitIfError(err)

    return repoResponse
}
