package gpt

import (
    "bytes"
    _ "embed"
    "encoding/json"
    "io"
    "net/http"
    "os"
    "time"

    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/pkg/xerr"
)

//go:embed prompt-commit.txt
var promptCommit string

// GenerateCommitMessage generates a commit message based on the changes (git diff)
// in the current git repository using the OpenAI GPT API
func GenerateCommitMessage(path string) (string, bool) {
    status, e := shell.RunCmdE(path, false, "git", "--no-pager", "diff")
    xerr.ExitIfError(e)

    if status.Stdout == "" {
        clog.Debug("No changes detected")
        return "", false

    } else if len(status.Stdout) > 4096 {
        clog.Warn("Too many changes, truncating diff to 4096 bytes")
        status.Stdout = status.Stdout[:4096]
    }

    openApiKey, exists := os.LookupEnv("OPENAI_API_KEY")
    if !exists {
        clog.Warn("OPENAI_API_KEY not set")
        return "", true
    }

    gpt := Request{
        Model: "gpt-3.5-turbo",
        Messages: []Content{
            {
                Role:    "user",
                Content: promptCommit + status.Stdout,
            },
        },
        Temperature: 0.7,
    }

    body, _ := json.MarshalIndent(gpt, "", "  ")

    httpRequest, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(body))
    xerr.ExitIfError(err)

    httpRequest.Header.Set("Authorization", "Bearer "+openApiKey)
    httpRequest.Header.Set("Content-Type", "application/json")

    httpClient := http.Client{
        Timeout: 30 * time.Second,
    }

    response, err := httpClient.Do(httpRequest)
    xerr.ExitIfError(err)

    responseBody, err := io.ReadAll(response.Body)
    xerr.ExitIfError(err)

    var gptResponse Response
    resErr := json.Unmarshal(responseBody, &gptResponse)
    if resErr != nil {
        clog.Error(resErr.Error())
    }

    if len(gptResponse.Choices) == 0 {
        clog.Warn(response.Status)
        clog.Info(string(responseBody))
        return "", true
    }

    content := gptResponse.Choices[0].Message.Content
    return content, true
}
