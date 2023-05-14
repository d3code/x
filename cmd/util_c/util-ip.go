package util_c

import (
    _ "embed"
    "io"
    "net/http"
    "time"

    "github.com/d3code/pkg/xerr"
    "github.com/spf13/cobra"
)

//go:embed util-ip.txt
var ipHelp string

func init() {
    Util.AddCommand(ipCommand)
    ipCommand.Flags().BoolP("verbose", "v", false, "verbose output showing additional information")
}

var ipCommand = &cobra.Command{
    Use:   "ip",
    Short: "Public IP address",
    Long:  ipHelp,
    Run: func(cmd *cobra.Command, args []string) {
        httpRequest, err := http.NewRequest("GET", "https://ipecho.net/plain", nil)
        xerr.ExitIfError(err)

        httpClient := http.Client{
            Timeout: 10 * time.Second,
        }

        response, err := httpClient.Do(httpRequest)
        xerr.ExitIfError(err)

        responseBody, err := io.ReadAll(response.Body)
        xerr.ExitIfError(err)

        cmd.Println(string(responseBody))
    },
}
