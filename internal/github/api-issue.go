package github

import (
    "encoding/json"
    "fmt"
    "github.com/d3code/pkg/slice_utils"
    "github.com/d3code/x/internal/cfg"
    "sort"
)

func CreateIssue(title string, body string, labels []string) IssueResponse {
    config := cfg.Configuration()
    keys := slice_utils.Keys(config.GitHub)

    var account string
    if len(keys) == 1 {
        account = keys[0]
    } else {
        account = Account()
    }

    repositories := RepositoriesWithIssues(account)
    sort.Sort(RepositoryList(repositories))

    repo := Repo(repositories)

    issue := IssueRequest{
        Title:     title,
        Body:      &body,
        Assignees: []string{account},
        Labels:    labels,
    }

    bodyJson, _ := json.Marshal(issue)

    var response IssueResponse
    url := fmt.Sprintf("/repos/%s/%s/issues", repo.Owner.Login, repo.Name)
    Request("POST", url, string(bodyJson), account, &response)

    return response
}
