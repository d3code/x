package github

import (
    "encoding/json"
    "fmt"
    "github.com/d3code/clog"
    "github.com/d3code/pkg/xerr"
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

func AccRepositories(user string) Repos {
    variablesMap := map[string]string{
        "user": user,
    }

    js, err := json.Marshal(variablesMap)
    xerr.ExitIfError(err)

    jsonMapInstance := map[string]string{
        "query":     queryRepos,
        "variables": string(js),
    }

    jsonResult, err := json.Marshal(jsonMapInstance)
    xerr.ExitIfError(err)

    var response Repos
    RequestGraph(jsonResult, "luk3sands", &response)

    for _, node := range response.Data.User.Repositories.Nodes {
        clog.InfoF("%s: %s", node.Name, node.Url)
    }

    for _, node := range response.Data.User.Organizations.Nodes {
        for _, s := range node.Repositories.Nodes {
            clog.InfoF("%s: %s", s.Name, s.Url)
        }
    }

    return response
}

type Repos struct {
    Data struct {
        User struct {
            Organizations struct {
                Nodes []struct {
                    Repositories struct {
                        Nodes []struct {
                            Name            string  `json:"name"`
                            Url             string  `json:"url"`
                            Description     *string `json:"description"`
                            PrimaryLanguage *struct {
                                Name string `json:"name"`
                            } `json:"primaryLanguage"`
                            Languages struct {
                                Nodes []struct {
                                    Name string `json:"name"`
                                } `json:"nodes"`
                            } `json:"languages"`
                        } `json:"nodes"`
                    } `json:"repositories"`
                } `json:"nodes"`
            } `json:"organizations"`
            Repositories struct {
                Nodes []struct {
                    Name            string `json:"name"`
                    Url             string `json:"url"`
                    Description     string `json:"description"`
                    PrimaryLanguage *struct {
                        Name string `json:"name"`
                    } `json:"primaryLanguage"`
                    Languages struct {
                        Nodes []struct {
                            Name string `json:"name"`
                        } `json:"nodes"`
                    } `json:"languages"`
                } `json:"nodes"`
            } `json:"repositories"`
        } `json:"user"`
    } `json:"data"`
}
