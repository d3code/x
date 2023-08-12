package input

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/slice_utils"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/prompt"
    "github.com/manifoldco/promptui"
    "os"
)

func Account() string {
    configuration := cfg.Configuration()
    accounts := configuration.GitHub
    keys := slice_utils.Keys(accounts)

    if len(accounts) == 0 {
        clog.Error("No accounts configured")
        os.Exit(1)
    } else if len(accounts) == 1 {
        return keys[0]
    }

    prompt := promptui.Select{
        Label:        "Account",
        Items:        keys,
        HideHelp:     true,
        HideSelected: true,
        Templates: &promptui.SelectTemplates{
            Label:    "{{ . }}",
            Active:   "{{ . | green }}",
            Inactive: "  {{ . }}",
            Details:  "",
            Help:     "",
        },
        Stdout: prompt.NoBellStdout,
    }

    run, _, err := prompt.Run()
    xerr.ExitIfError(err)

    clog.InfoF("Selected account: %s", keys[run])
    return keys[run]
}
