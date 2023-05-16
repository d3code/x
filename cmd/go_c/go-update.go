package go_c

import (
    "fmt"
    "github.com/d3code/clog"
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
                clog.Underline("Updating", path)
                UpdateGo(path)
            }
        } else {
            path := shell.CurrentDirectory()
            UpdateGo(path)
        }
    },
}

func UpdateGo(directory string) {
    project, err := shell.RunCmdE(directory, false, "go", "list", "-m")
    if err != nil {
        clog.Warn("No go project found")
        return
    }

    list := strings.Split(project.Out, "\n")

    graph := shell.RunCmd(directory, false, "go", "mod", "graph")
    lines := strings.Split(graph.Out, "\n")

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
                clog.InfoF("Updating dependent {{ go | green }} module {{ %s | blue }}", golang.Name)
                git.StageCommitFetchPullPush(path, "")

                commit := shell.RunShell(false, "(cd "+path+";git rev-parse HEAD 2>/dev/null)")
                dependencyVersions[golang.Name] = commit.Out

                fmt.Println()
            }
        }
    }

    for m, commit := range dependencyVersions {
        shell.RunCmd(directory, false, "go", "get", m+"@"+commit)
    }

    shell.RunCmd(directory, false, "go", "get", "-t", "-u", "./...")
    shell.RunCmd(directory, false, "go", "mod", "tidy")
}
