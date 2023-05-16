package github

import "time"

type IssueRequest struct {
    Title     string   `json:"title"`
    Body      *string  `json:"body"`
    Assignees []string `json:"assignees"`
    Labels    []string `json:"labels"`
    Milestone *int     `json:"milestone"`
}

type IssueResponse struct {
    Id            int    `json:"id"`
    NodeId        string `json:"node_id"`
    Url           string `json:"url"`
    RepositoryUrl string `json:"repository_url"`
    LabelsUrl     string `json:"labels_url"`
    CommentsUrl   string `json:"comments_url"`
    EventsUrl     string `json:"events_url"`
    HtmlUrl       string `json:"html_url"`
    Number        int    `json:"number"`
    State         string `json:"state"`
    Title         string `json:"title"`
    Body          string `json:"body"`
    User          struct {
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
    } `json:"user"`
    Labels []struct {
        Id          int    `json:"id"`
        NodeId      string `json:"node_id"`
        Url         string `json:"url"`
        Name        string `json:"name"`
        Description string `json:"description"`
        Color       string `json:"color"`
        Default     bool   `json:"default"`
    } `json:"labels"`
    Assignee struct {
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
    } `json:"assignee"`
    Assignees []struct {
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
    } `json:"assignees"`
    Milestone struct {
        Url         string `json:"url"`
        HtmlUrl     string `json:"html_url"`
        LabelsUrl   string `json:"labels_url"`
        Id          int    `json:"id"`
        NodeId      string `json:"node_id"`
        Number      int    `json:"number"`
        State       string `json:"state"`
        Title       string `json:"title"`
        Description string `json:"description"`
        Creator     struct {
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
        OpenIssues   int       `json:"open_issues"`
        ClosedIssues int       `json:"closed_issues"`
        CreatedAt    time.Time `json:"created_at"`
        UpdatedAt    time.Time `json:"updated_at"`
        ClosedAt     time.Time `json:"closed_at"`
        DueOn        time.Time `json:"due_on"`
    } `json:"milestone"`
    Locked           bool   `json:"locked"`
    ActiveLockReason string `json:"active_lock_reason"`
    Comments         int    `json:"comments"`
    PullRequest      struct {
        Url      string `json:"url"`
        HtmlUrl  string `json:"html_url"`
        DiffUrl  string `json:"diff_url"`
        PatchUrl string `json:"patch_url"`
    } `json:"pull_request"`
    ClosedAt  interface{} `json:"closed_at"`
    CreatedAt time.Time   `json:"created_at"`
    UpdatedAt time.Time   `json:"updated_at"`
    ClosedBy  struct {
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
    } `json:"closed_by"`
    AuthorAssociation string `json:"author_association"`
    StateReason       string `json:"state_reason"`
}
