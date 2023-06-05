package github

import (
    "fmt"
    "github.com/d3code/clog"
)

func Repositories(account string) []RepoResponse {
    repositories := AccountRepositories(account)

    orgResponse := Org(account)
    for _, org := range orgResponse {
        repoResponse := OrgRepositories(org.Login, account)
        clog.InfoF("Repositories: %v", len(repoResponse))
        for _, repo := range repoResponse {
            repositories = append(repositories, repo)
        }
    }

    return repositories
}

func RepositoriesWithIssues(account string) []RepoResponse {
    repositories := AccountRepositories(account)

    orgResponse := Org(account)
    for _, org := range orgResponse {
        repoResponse := OrgRepositories(org.Login, account)
        for _, repo := range repoResponse {
            repositories = append(repositories, repo)
        }
    }

    var repositoriesWithIssues []RepoResponse
    for _, repo := range repositories {
        if repo.HasIssues {
            repositoriesWithIssues = append(repositoriesWithIssues, repo)
        }
    }

    return repositoriesWithIssues
}

func Org(account string) []OrgResponse {
    url := fmt.Sprintf("/users/%s/orgs", account)

    var response []OrgResponse
    Request("GET", url, "", account, &response)

    return response
}

func OrgRepositories(org string, account string) []RepoResponse {
    url := fmt.Sprintf("/orgs/%s/repos", org)

    var response []RepoResponse
    Request("GET", url, "", account, &response)

    return response
}

func AccountRepositories(account string) []RepoResponse {
    url := fmt.Sprintf("/users/%s/repos", account)

    var response []RepoResponse
    Request("GET", url, "", account, &response)

    return response
}
