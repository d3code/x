package golang

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "strings"
)

func GoModule(directory string) (string, bool) {
    if !Go(directory) {
        return "", false
    }

    project, err := shell.RunCmdE(directory, false, "go", "list", "-m")
    if err != nil {
        clog.Debug(directory, "no go project found")
        return "", false
    }

    list := strings.Split(project.Stdout, "\n")
    if len(list) == 0 {
        clog.Error(directory, "no go module found")
        return "", false
    } else if len(list) > 1 {
        clog.WarnF("%s multiple go modules found %v\n", "returning first module %s", directory, list, list[0])
        return list[0], false
    }

    return list[0], true
}
