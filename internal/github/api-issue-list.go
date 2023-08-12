package github

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/slice_utils"
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/github/input"
    "sort"
    "time"
)

func ListIssue() []IssueList {
    config := cfg.Configuration()
    keys := slice_utils.Keys(config.GitHub)

    var account string
    if len(keys) == 1 {
        account = keys[0]
    } else {
        account = input.Account()
    }

    repositories := RepositoriesWithIssues(account)
    sort.Sort(RepositoryList(repositories))

    var response []IssueList
    //Request("GET", "/issues", "", account, &response)

    var response1 []OrgProjects
    Request("GET", "/orgs/d3code/projects", "", account, &response1)

    for _, projects := range response1 {
        clog.Info(projects.Name)
    }

    return response
}

type OrgProjects struct {
    OwnerUrl   string `json:"owner_url"`
    Url        string `json:"url"`
    HtmlUrl    string `json:"html_url"`
    ColumnsUrl string `json:"columns_url"`
    Id         int    `json:"id"`
    NodeId     string `json:"node_id"`
    Name       string `json:"name"`
    Body       string `json:"body"`
    Number     int    `json:"number"`
    State      string `json:"state"`
    Creator    struct {
        Login             string `json:"login"`
        Id                int    `json:"id"`
        NodeId            string `json:"node_id"`
        AvatarUrl         string `json:"avatar_url"`
        GravatarId        string `json:"gravatar_id"`
        Url               string `json:"url"`
        HtmlUrl           string `json:"html_url"`
        FollowersUrl      string `json:"followers_url"`
        FollowingUrl      string `json:"following_url"`
        GistsUrl          string `json:"gists_url"`
        StarredUrl        string `json:"starred_url"`
        SubscriptionsUrl  string `json:"subscriptions_url"`
        OrganizationsUrl  string `json:"organizations_url"`
        ReposUrl          string `json:"repos_url"`
        EventsUrl         string `json:"events_url"`
        ReceivedEventsUrl string `json:"received_events_url"`
        Type              string `json:"type"`
        SiteAdmin         bool   `json:"site_admin"`
    } `json:"creator"`
    CreatedAt              time.Time `json:"created_at"`
    UpdatedAt              time.Time `json:"updated_at"`
    OrganizationPermission string    `json:"organization_permission"`
    Private                bool      `json:"private"`
}
