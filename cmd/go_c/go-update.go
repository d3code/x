package go_c

import (
    "github.com/d3code/pkg/clog"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/git"
    "github.com/spf13/cobra"
    "strings"
)

func init() {
    Go.AddCommand(Update)
    Update.Flags().BoolP("all", "a", false, "update all go projects")
}

var Update = &cobra.Command{
    Use:     "update",
    Aliases: []string{"up"},
    Short:   "Update go project",
    Run: func(cmd *cobra.Command, args []string) {

        all, err := cmd.Flags().GetBool("all")
        xerr.ExitIfError(err)
        if all {
            configuration := cfg.Configuration()
            for path, _ := range configuration.Golang {
                clog.Underline("Updating {{", path, "|blue}}")
                update(path)
            }
        } else {
            path := shell.CurrentDirectory()
            update(path)
        }
    },
}

func update(directory string) {
    clog.Info("{{Updating go project...|green}}")

    pro, err := shell.RunDirE(directory, "go", "list", "-m")
    if err != nil {
        clog.Info("{{No go project found|yellow}}")
        return
    }
    list := strings.Split(pro, "\n")

    graph := shell.RunDir(directory, "go", "mod", "graph")
    lines := strings.Split(graph, "\n")

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

    dependencyVersions := make(map[string]string)
    configuration := cfg.Configuration()

    for _, module := range modules {
        for path, golang := range configuration.Golang {
            if golang.Name == module {

                message := "Update " + golang.Name
                git.StageCommitFetchPullPush(path, message)

                commit := shell.RunShell("(cd " + path + ";git rev-parse HEAD 2>/dev/null)")
                dependencyVersions[golang.Name] = commit
            }
        }
    }

    for m, commit := range dependencyVersions {
        shell.RunOutDir(directory, "go", "get", m+"@"+commit)
    }

    shell.RunOutDir(directory, "go", "get", "-u", "./...")
    shell.RunOutDir(directory, "go", "mod", "tidy")
}
