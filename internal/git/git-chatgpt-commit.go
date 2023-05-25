package git

import (
    "bytes"
    "encoding/json"
    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/pkg/xerr"
    "io"
    "net/http"
    "os"
    "time"
)

func ChatGPT(path string) string {
    status, e := shell.RunCmdE(path, false, "git", "--no-pager", "diff")
    xerr.ExitIfError(e)

    if status.Out == "" {
        clog.Warn("No changes")
        return ""
    }

    gpt := GPT{
        Model: "gpt-3.5-turbo",
        Messages: []GPTContent{
            {
                Role:    "user",
                Content: "What is a good git commit message based on the following changes? If you arent able to determine a commit message, simply reply with a very general commit message such as 'Update project', but more verbose.\n" + status.Out,
            },
        },
        Temperature: 0.7,
    }

    body, _ := json.MarshalIndent(gpt, "", "  ")

    httpRequest, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(body))
    xerr.ExitIfError(err)

    openApiKey := os.Getenv("OPENAI_API_KEY")
    httpRequest.Header.Set("Authorization", "Bearer "+openApiKey)
    httpRequest.Header.Set("Content-Type", "application/json")

    httpClient := http.Client{
        Timeout: 30 * time.Second,
    }

    response, err := httpClient.Do(httpRequest)
    xerr.ExitIfError(err)

    responseBody, err := io.ReadAll(response.Body)
    xerr.ExitIfError(err)

    var gptResponse GPTResponse
    resErr := json.Unmarshal(responseBody, &gptResponse)
    if resErr != nil {
        clog.Error(resErr.Error())
    }

    if len(gptResponse.Choices) == 0 {
        clog.Warn("No response from GPT")
        return ""
    }

    content := gptResponse.Choices[0].Message.Content
    clog.InfoF("{{ %s | grey }} {{ %s | blue }}", "[commit message]", content)

    return content
}

type GPTResponse struct {
    Id      string `json:"id"`
    Object  string `json:"object"`
    Created int    `json:"created"`
    Model   string `json:"model"`
    Usage   struct {
        PromptTokens     int `json:"prompt_tokens"`
        CompletionTokens int `json:"completion_tokens"`
        TotalTokens      int `json:"total_tokens"`
    } `json:"usage"`
    Choices []struct {
        Message struct {
            Role    string `json:"role"`
            Content string `json:"content"`
        } `json:"message"`
        FinishReason string `json:"finish_reason"`
        Index        int    `json:"index"`
    } `json:"choices"`
}

type GPTContent struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type GPT struct {
    Model       string       `json:"model"`
    Messages    []GPTContent `json:"messages"`
    Temperature float64      `json:"temperature"`
}
