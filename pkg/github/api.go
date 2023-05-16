package github

import (
    "encoding/json"
    "fmt"
    "github.com/d3code/clog"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/pkg/cfg"
    "io"
    "net/http"
)

func Repositories(account string) []RepoResponse {
    repositories := AccountRepositories(account)

    orgResponse := Org(account)
    for _, org := range orgResponse {
        repoResponse := OrgRepositories(org.Login, account)
        for _, repo := range repoResponse {
            repositories = append(repositories, repo)
        }
    }

    return repositories
}

func Org(account string) []OrgResponse {
    url := fmt.Sprintf("https://api.github.com/users/%s/orgs", account)

    var response []OrgResponse
    GithubRequest(url, account, &response)

    return response
}

func OrgRepositories(org string, account string) []RepoResponse {
    url := fmt.Sprintf("https://api.github.com/orgs/%s/repos", org)

    var response []RepoResponse
    GithubRequest(url, account, &response)

    return response
}

func AccountRepositories(account string) []RepoResponse {
    url := fmt.Sprintf("https://api.github.com/users/%s/repos", account)

    var response []RepoResponse
    GithubRequest(url, account, &response)

    return response
}

func GithubRequest(url string, account string, response any) {
    configuration := cfg.Configuration()
    gh, ok := configuration.GitHub[account]

    client := &http.Client{}
    req, _ := http.NewRequest("GET", url, nil)

    if ok && gh.Token != "" {
        auth := fmt.Sprintf("Bearer %s", gh.Token)
        req.Header.Set("Authorization", auth)
    } else {
        clog.Warn("No token configured for account " + account)
    }

    res, _ := client.Do(req)
    responseBody, _ := io.ReadAll(res.Body)

    clog.Debug(string(responseBody))

    err := json.Unmarshal(responseBody, response)
    xerr.ExitIfError(err)
}
