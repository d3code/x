package util_c

import (
    _ "embed"
    "io"
    "net/http"
    "time"

    "github.com/d3code/clog"
    "github.com/d3code/x/internal/embed_text"

    "github.com/d3code/pkg/xerr"
    "github.com/spf13/cobra"
)

func init() {
    Util.AddCommand(ipCommand)
}

var ipCommand = &cobra.Command{
    Use:   "ip",
    Short: "Public IP address",
    Long:  embed_text.UtilIp,
    Run: func(cmd *cobra.Command, args []string) {

        ip := GetIp()
        clog.Info(ip)
    },
}

func GetIp() string {
    httpRequest, err := http.NewRequest("GET", "https://ipecho.net/plain", nil)
    xerr.ExitIfError(err)

    httpClient := http.Client{
        Timeout: 10 * time.Second,
    }

    response, err := httpClient.Do(httpRequest)
    xerr.ExitIfError(err)

    responseBody, err := io.ReadAll(response.Body)
    xerr.ExitIfError(err)

    return string(responseBody)
}
