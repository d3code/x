package git

import (
    "github.com/d3code/x/internal/cfg"
    "os"
    "path"
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
}

func ScanSubdirectories(wg *sync.WaitGroup, newPath string) {
    files, _ := os.ReadDir(newPath)

    for _, file := range files {
        directory := path.Join(newPath, file.Name())

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
