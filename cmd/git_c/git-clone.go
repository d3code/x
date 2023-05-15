package git_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/pkg/cobra_util"
    "github.com/d3code/x/pkg/git"
    "github.com/spf13/cobra"
    "regexp"
)

func init() {
    Git.AddCommand(cloneCmd)
}

var cloneCmd = &cobra.Command{
    Use: "clone",

    Run: func(cmd *cobra.Command, args []string) {
        if git.Git(".") {
            clog.Error("Current directory is already a git repository")
            return
        }

        var repository string
        if len(args) > 0 {
            repository = args[0]
        } else {
            repository = cobra_util.PromptString("Repository to clone", true)
        }

        url := git.FormatRepositoryUrl(repository)
        shell.RunCmd(".", true, "git", "clone", url)
    },
}

func clonedDirectory(message string) string {
    re := regexp.MustCompile(`Cloning into '([^'/].+)'`)
    matches := re.FindAllStringSubmatch(message, -1)

    if len(matches) == 1 && len(matches[0]) == 2 {
        return matches[0][1]
    }

    return ""
}
