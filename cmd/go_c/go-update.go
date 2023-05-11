package go_c

import (
    "fmt"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/cmd/git_c"
    "github.com/d3code/x/internal/cfg"
    "github.com/spf13/cobra"
    "strings"
)

func init() {
    Go.AddCommand(Update)
}

var Update = &cobra.Command{
    Use:   "update",
    Short: "Update go project",
    Run: func(cmd *cobra.Command, args []string) {
        shell.Println("{{Updating go project...|green}}")
        //directory := shell.CurrentDirectory()

        pro := shell.Run("go", "list")
        list := strings.Split(pro, "\n")

        graph := shell.Run("go", "mod", "graph")
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
            //shell.RunOut("go", "get", "-u", module)
            for path, golang := range configuration.Golang {
                if golang.Name == module {
                    git_c.Commit(path, false, "Update "+golang.Name)
                    commit := shell.RunShell("(cd " + path + ";git rev-parse HEAD 2>/dev/null)")
                    fmt.Println(path, golang.Name, commit)
                    dependencyVersions[golang.Name] = commit
                }
            }
        }

        for m, commit := range dependencyVersions {
            fmt.Println(m, commit)
            shell.RunOut("go", "get", m+"@"+commit)
        }

        shell.RunOut("go", "mod", "tidy")

        //shell.RunOut("go", "get", "-u", "./...")
    },
}
