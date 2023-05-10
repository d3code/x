package config_c

import (
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/git"
    "github.com/spf13/cobra"
    "os"
    "strings"
    "sync"
)

func init() {
    Config.AddCommand(Scan)
}

var Scan = &cobra.Command{
    Use:   "scan",
    Short: "Scan for git repositories",
    Run: func(cmd *cobra.Command, args []string) {
        shell.Println("{{Scanning for git repositories...|green}}")
        directory := shell.CurrentDirectory()
        scanGitDirectory(directory)
        removeNotGitRepo()
    },
}

func removeNotGitRepo() {
    config := cfg.Configuration()
    for path, _ := range config.Git {
        if !git.IsGitDirectory(path) {
            config.DeleteGitDirectory(path)
        }
    }
    config.Save()
}

func scanGitDirectory(directory string) {
    var wg sync.WaitGroup
    if git.IsGitDirectory(directory) {
        remote, _ := git.Remote(directory)
        cfg.Configuration().AddGitDirectory(directory, cfg.Git{Remote: remote})
    } else {
        wg.Add(1)
        go scanSubdirectories(&wg, directory)
    }
    wg.Wait()
    cfg.Configuration().Save()
}

func scanSubdirectories(wg *sync.WaitGroup, path string) {
    files, _ := os.ReadDir(path)
    for _, file := range files {
        var directory string
        if strings.HasSuffix(path, "/") {
            directory = path + file.Name()
        } else {
            directory = path + "/" + file.Name()
        }
        if git.IsGitDirectory(directory) {
            remote, _ := git.Remote(directory)
            cfg.Configuration().AddGitDirectory(directory, cfg.Git{Remote: remote})
        } else if file.IsDir() {
            wg.Add(1)
            go scanSubdirectories(wg, directory)
        }
    }
    wg.Done()
}
