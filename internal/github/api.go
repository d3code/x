package github

import (
    "encoding/json"
    "fmt"
    "github.com/d3code/clog"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/internal/cfg"
    "io"
    "net/http"
    "strings"
)

func Request(method string, url string, body string, account string, response any) {
    configuration := cfg.Configuration()
    gh, ok := configuration.GitHub[account]

    client := &http.Client{}
    reader := strings.NewReader(body)
    req, _ := http.NewRequest(method, "https://api.github.com"+url, reader)

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
