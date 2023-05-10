package terminal

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

func GetFlagBool(cmd *cobra.Command, name string) bool {
    x, err := cmd.Flags().GetBool(name)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    return x
}

func CheckArgs(args []string, required int, message string) {
    if len(args) < required {
        fmt.Println(message)
        os.Exit(1)
    }
}
