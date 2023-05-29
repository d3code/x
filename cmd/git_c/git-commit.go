package git_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/pkg/slice_utils"
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/git"
    "github.com/d3code/x/internal/golang"
    "github.com/spf13/cobra"
)

func init() {
    Git.AddCommand(Commit)
    Commit.Flags().BoolP("all", "a", false, "all configured repositories")
    Commit.Flags().BoolP("interactive", "i", false, "interactive mode")
    Commit.Flags().BoolP("push", "p", false, "push to remote")
}

var Commit = &cobra.Command{
    Use:     "commit",
    Aliases: []string{"c"},
    Short:   "Commit and push changes to the remote repository",
    Run: func(cmd *cobra.Command, args []string) {
        all, err := cmd.Flags().GetBool("all")
        xerr.ExitIfError(err)

        push, err := cmd.Flags().GetBool("push")
        xerr.ExitIfError(err)

        interactive, err := cmd.Flags().GetBool("interactive")
        xerr.ExitIfError(err)

        if all {
            repositories := slice_utils.Keys(cfg.Configuration().Git)

            for _, repository := range repositories {

                clog.UnderlineF("Checking {{ %s | blue }}", repository)

                golang.UpdateGo(repository)
                git.CommitDirectory(repository, interactive)

                if push {
                    git.FetchPullPush(repository)
                }
            }
        }

        directory := shell.CurrentDirectory()
        golang.UpdateGo(directory)
        git.CommitDirectory(directory, interactive)

        if push {
            git.FetchPullPush(directory)
        }
    },
}
