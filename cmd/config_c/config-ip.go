package config_c

import (
    "github.com/spf13/cobra"
    "io"
    "net/http"
    "time"
)

func init() {
    Config.AddCommand(ipCmd)
}

var ipCmd = &cobra.Command{
    Use:   "ip",
    Short: "Display your public IP address",
    Long:  `The public IP address of the machine as seen by the internet`,
    RunE:  ip,
}

func ip(cmd *cobra.Command, args []string) error {
    httpRequest, err := http.NewRequest("GET", "https://ipecho.net/plain", nil)
    if err != nil {
        cmd.Println(err)
        return err
    }

    httpClient := http.Client{
        Timeout: 10 * time.Second,
    }

    response, requestError := httpClient.Do(httpRequest)
    if requestError != nil {
        cmd.Println(requestError)
        return err
    }

    responseBody, responseError := io.ReadAll(response.Body)
    if responseError != nil {
        cmd.Println(responseError)
        return err
    }

    cmd.Println(string(responseBody))

    return nil
}
