package cfg

import (
    "fmt"
    "github.com/d3code/pkg/shell"
)

func ListGit() {
    shell.PrintHeader("GitHub accounts")
    for name, account := range localConfig.GitHub {
        fmt.Println(name)
        fmt.Println("  org    [", account.Organization, "]")
        fmt.Println("  token  [", account.Token, "]")
        fmt.Println()
    }
}
