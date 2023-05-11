package git

import (
    "github.com/d3code/x/internal/cfg"
    "os"
    "strings"
    "sync"
)

func ScanGitDirectory(directory string) {
    var wg sync.WaitGroup
    if IsGitDirectory(directory) {
        remote, _ := Remote(directory)
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
        if IsGitDirectory(directory) {
            remote, _ := Remote(directory)
            cfg.Configuration().AddGitDirectory(directory, cfg.Git{Remote: remote})
        } else if file.IsDir() {
            wg.Add(1)
            go scanSubdirectories(wg, directory)
        }
    }
    wg.Done()
}

func RemoveNotGitRepo() {
    config := cfg.Configuration()
    for path, _ := range config.Git {
        if !IsGitDirectory(path) {
            config.DeleteGitDirectory(path)
        }
    }
    config.Save()
}
