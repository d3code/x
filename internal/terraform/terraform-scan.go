package terraform

import (
    "github.com/d3code/pkg/shell"
    "github.com/d3code/x/internal/cfg"
    "os"
    "strings"
    "sync"
)

func Scan(directory string) {
    var wg sync.WaitGroup
    if Terraform(directory) {
        cfg.Configuration().AddTerraform(directory, cfg.Terraform{})
    } else {
        wg.Add(1)
        go scanSubdirectories(&wg, directory)
    }
    wg.Wait()
    cfg.Configuration().Save()
}

func scanSubdirectories(wg *sync.WaitGroup, path string) {
    defer wg.Done()
    home := shell.UserHomeDirectory()

    files, _ := os.ReadDir(path)
    for _, file := range files {
        var directory string
        if strings.HasSuffix(path, "/") {
            directory = path + file.Name()
        } else {
            directory = path + "/" + file.Name()
        }
        if Terraform(directory) {
            cfg.Configuration().AddTerraform(directory, cfg.Terraform{})
        }
        if file.IsDir() &&
            !strings.HasPrefix(directory, home+"/Library/") &&
            !strings.Contains(directory, "/.git/") &&
            !strings.Contains(directory, "/.terraform/modules/") {
            wg.Add(1)
            go scanSubdirectories(wg, directory)
        }
    }
}

func VerifyPaths() {
    config := cfg.Configuration()
    for path := range config.Terraform {
        if !Terraform(path) {
            config.DeleteTerraform(path)
        }
    }
    config.Save()
}
