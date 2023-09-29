package github

import (
    "bytes"
    _ "embed"
    "encoding/json"
    "fmt"
    "github.com/d3code/clog"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/github/graphql/model"
    "io"
    "net/http"
)

//go:embed graphql/queries/projects.graphql
var query string

//go:embed graphql/queries/repos.graphql
var queryRepos string

func GetProjects() {
    variablesMap := map[string]string{
        "org": "d3code",
    }

    js, err := json.Marshal(variablesMap)
    xerr.ExitIfError(err)

    jsonMapInstance := map[string]string{
        "query":     query,
        "variables": string(js),
    }

    jsonResult, err := json.Marshal(jsonMapInstance)
    xerr.ExitIfError(err)

    var response ProjectsResponse
    RequestGraph(jsonResult, "luk3sands", &response)

    for _, node := range response.Data.Organization.ProjectsV2.Nodes {
        clog.Infof("%s: %s", node.ID, node.Title)
    }
}

func RequestGraph(body []byte, account string, response any) {
    configuration := cfg.Configuration()
    gh, ok := configuration.GitHub[account]

    client := &http.Client{}
    req, _ := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewReader(body))

    if ok && gh.Token != "" {
        auth := fmt.Sprintf("Bearer %s", gh.Token)
        req.Header.Set("Authorization", auth)
    } else {
        clog.Warn("No token configured for account " + account)
    }

    res, _ := client.Do(req)
    responseBody, _ := io.ReadAll(res.Body)

    err := json.Unmarshal(responseBody, response)
    xerr.ExitIfError(err)
}

type ProjectsResponse struct {
    Data struct {
        Organization model.Organization `json:"organization"`
    } `json:"data"`
}
