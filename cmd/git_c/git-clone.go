package git_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/git"
    "github.com/d3code/x/internal/prompt"
    "github.com/spf13/cobra"
    "os"
    "regexp"
)

func init() {
    Git.AddCommand(cloneCmd)
}

var cloneCmd = &cobra.Command{
    Use:   "clone",
    Short: "Clone a git repository",
    Run: func(cmd *cobra.Command, args []string) {
        if git.Git(".") {
            clog.Error("Current directory is already a git repository")
            return
        }

        var repository string
        if len(args) > 0 {
            repository = args[0]
        } else {
            repository = prompt.String("Repository to clone", true)
        }

        url := git.FormatRepositoryUrl(repository)
        e := shell.RunCmd(".", false, "git", "clone", url)

        directory := clonedDirectory(e.Stderr)

        clog.Info("Cloned into {{ " + directory + " | blue }}")
        git.GitignoreCreate(directory)

        remote, _ := git.Remote(directory)
        cfg.Configuration().AddGitDirectory(directory, cfg.Git{Remote: remote})
    },
}

func clonedDirectory(message string) string {
    re := regexp.MustCompile(`Cloning into '([^'/].+)'`)
    matches := re.FindAllStringSubmatch(message, -1)

    if len(matches) == 1 && len(matches[0]) == 2 {
        return shell.FullPath(matches[0][1])
    }

    clog.Info(message)
    clog.Error("Could not determine cloned directory")
    os.Exit(1)

    return ""
}
