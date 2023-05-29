package golang

import (
    "github.com/d3code/pkg/shell"
    "strings"
)

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
