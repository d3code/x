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
    modules, project := GoModules(directory)
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
                }

                git.Push(path)
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

func GoDependencies(directory string, list []string) []string {
    graph := shell.RunCmd(directory, false, "go", "mod", "graph")
    lines := strings.Split(graph.Stdout, "\n")

    var modules []string
    for _, line := range lines {
        if line == "" {
            continue
        }

        split := strings.Split(line, " ")
        if len(split) == 2 && split[0] == list[0] {
            dependency := split[1]
            dependency = strings.Split(dependency, "@")[0]
            modules = append(modules, dependency)
        }
    }
    return modules
}

func GoModules(directory string) ([]string, bool) {
    project, err := shell.RunCmdE(directory, false, "go", "list", "-m")
    if err != nil {
        clog.Warn("No go project found")
        return nil, true
    }

    list := strings.Split(project.Stdout, "\n")
    return list, false
}
