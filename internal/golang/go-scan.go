package golang

import (
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/internal/cfg"
    "os"
    "strings"
    "sync"
)

func Scan(directory string) {
    var wg sync.WaitGroup
    if name, module := GoModule(directory); module {
        cfg.Configuration().AddGolang(directory, cfg.Golang{
            Module: name,
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
        if name, module := GoModule(directory); module {
            cfg.Configuration().AddGolang(directory, cfg.Golang{
                Module: name,
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
        if !Go(path) {
            config.DeleteGolang(path)
        }
    }
}
