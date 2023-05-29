package golang

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/git"
    "github.com/d3code/x/internal/gpt"
)

func UpdateGo(directory string) {
    modules, project := GoModule(directory)
    if !project {
        clog.DebugF("%s not a Go project", directory)
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
                    clog.InfoF("No changes in dependent {{ Go | green }} module {{ %s | blue }}", golang.Module)
                } else {
                    clog.InfoF("Updating dependent {{ Go | green }} module {{ %s | blue }}", golang.Module)
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
