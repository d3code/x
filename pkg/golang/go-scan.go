package golang

import (
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/pkg/cfg"
    "os"
    "strings"
    "sync"
)

func Scan(directory string) {
    var wg sync.WaitGroup
    if name := Go(directory); name != "" {
        cfg.Configuration().AddGolang(directory, cfg.Golang{
            Name: name,
        })
    } else {
        wg.Add(1)
        go scanSubdirectories(&wg, directory)
    }
    wg.Wait()
}

func scanSubdirectories(wg *sync.WaitGroup, path string) {
    files, _ := os.ReadDir(path)
    home := shell.UserHomeDirectory()

    for _, file := range files {
        var directory string
        if strings.HasSuffix(path, "/") {
            directory = path + file.Name()
        } else {
            directory = path + "/" + file.Name()
        }
        if name := Go(directory); name != "" {
            cfg.Configuration().AddGolang(directory, cfg.Golang{
                Name: name,
            })
        } else if file.IsDir() &&
            !strings.HasPrefix(directory, home+"/Library/") &&
            !strings.HasPrefix(directory, home+"/go/") &&
            !strings.Contains(directory, "/.git/") {

            wg.Add(1)
            go scanSubdirectories(wg, directory)
        }
    }
    wg.Done()
}

func VerifyPaths() {
    config := cfg.Configuration()
    for path, _ := range config.Golang {
        if name := Go(path); name == "" {
            config.DeleteGolang(path)
        }
    }
}
