package golang

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/git"
    "github.com/d3code/x/internal/gpt"
    "strings"
)

func UpdateGo(directory string) {
    modules, project := GoModule(directory)
    if !project {
        clog.DebugF("%s not a go project", directory)
        return
    }

    dependencies := GoDependencies(directory, modules)
    dependencyVersions := make(map[string]string)
    for _, dependency := range dependencies {
        clog.DebugF("%s", dependency)

        for path, golang := range cfg.Configuration().Golang {
            if golang.Module == dependency {
                commitMessage, changes := gpt.GenerateCommitMessage(path)

                if !changes {
                    clog.InfoF("No changes in dependent {{ go | green }} module {{ %s | blue }}", path)
                } else {
                    clog.InfoF("Updating dependent {{ go | green }} module {{ %s | blue }}", golang.Module)
                    git.StageCommit(path, commitMessage)
                    err := git.Push(path)
                    if err != nil {
                        clog.Error(err.Error())
                        continue
                    }
                }

                commitHash := git.CommitHash(path)
                dependencyVersions[golang.Module] = commitHash.Stdout
            }
        }
    }

    for m, commit := range dependencyVersions {
        shell.RunCmd(directory, false, "go", "get", m+"@"+commit)
    }

    shell.RunCmd(directory, false, "go", "get", "-t", "-u", "./...")
    shell.RunCmd(directory, false, "go", "mod", "tidy")
}

func GoDependencies(directory string, module string) []string {
    graph := shell.RunCmd(directory, false, "go", "mod", "graph")
    dependencies := strings.Split(graph.Stdout, "\n")

    var modules []string
    for _, relationship := range dependencies {
        split := strings.Split(relationship, " ")
        if len(split) == 2 && split[0] == module {
            dependency := split[1]
            dependency = strings.Split(dependency, "@")[0]
            modules = append(modules, dependency)
        }
    }

    return modules
}

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
