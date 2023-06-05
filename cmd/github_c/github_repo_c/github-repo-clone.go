package github_repo_c

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/files"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/git"
    "github.com/d3code/x/internal/github"
    "github.com/d3code/x/internal/golang"
    "github.com/spf13/cobra"
    "path"
    "sort"
)

func init() {
    Root.AddCommand(Clone)
}

var Clone = &cobra.Command{
    Use: "clone",

    Run: func(cmd *cobra.Command, args []string) {
        account := github.Account()
        repositories := github.Repositories(account)
        sort.Sort(github.RepositoryList(repositories))

        baseDir := cfg.BaseDirectory()
        for _, repo := range repositories {

            pathName := path.Join(baseDir, repo.FullName)

            if files.Exist(pathName) {
                //clog.InfoL("{{ Repository | green }} "+repo.Owner.Login+"/"+repo.Name, "{{ Clone URL | grey }}  "+repo.SshUrl, "{{ URL | grey }}        "+repo.HtmlUrl)
                //clog.InfoF("{{%s | green}} %s", repo.Name, " exists")
                continue
            }

            clog.InfoL("{{ Repository | blue }} "+repo.Owner.Login+"/"+repo.Name, "{{ Clone URL | grey }}  "+repo.SshUrl, "{{ URL | grey }}        "+repo.HtmlUrl)

            ownerBase := path.Join(baseDir, repo.Owner.Login)
            shell.RunCmd(ownerBase, true, "git", "clone", repo.SshUrl)
        }

        git.Scan(baseDir)
        git.VerifyPaths()

        golang.Scan(baseDir)
        golang.VerifyPaths()
    },
}
