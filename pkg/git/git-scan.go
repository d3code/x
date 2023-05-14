package git

import (
    "github.com/d3code/x/pkg/cfg"
    "os"
    "strings"
    "sync"
)

func Scan(directory string) {
    var wg sync.WaitGroup
    if Git(directory) {
        remote, _ := Remote(directory)
        cfg.Configuration().AddGitDirectory(directory, cfg.Git{Remote: remote})
    } else {
        wg.Add(1)
        go ScanSubdirectories(&wg, directory)
    }
    wg.Wait()
    cfg.Configuration().Save()
}

func ScanSubdirectories(wg *sync.WaitGroup, path string) {
    files, _ := os.ReadDir(path)
    for _, file := range files {
        var directory string
        if strings.HasSuffix(path, "/") {
            directory = path + file.Name()
        } else {
            directory = path + "/" + file.Name()
        }
        if Git(directory) {
            remote, _ := Remote(directory)
            cfg.Configuration().AddGitDirectory(directory, cfg.Git{Remote: remote})
        } else if file.IsDir() {
            wg.Add(1)
            go ScanSubdirectories(wg, directory)
        }
    }
    wg.Done()
}
