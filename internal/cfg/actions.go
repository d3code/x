package cfg

import (
    "encoding/json"
    "fmt"
    "github.com/d3code/pkg/errors"
    "github.com/d3code/pkg/shell"
    "os"
)

func Configuration() *Config {
    return localConfig
}

func (c *Config) Save() {
    file := fmt.Sprintf("%s/config.json", configurationPath())

    configJson, err := json.MarshalIndent(c, "", "  ")
    errors.ExitIfError(err)

    err = os.WriteFile(file, configJson, 0666)
    errors.ExitIfError(err)
}

func (c *Config) AddGitDirectory(path string, git Git) {
    mapMutex.Lock()
    if _, ok := c.Git[path]; !ok {
        shell.Println("[ git   ] {{ " + path + " | blue }} added")
        c.Git[path] = git
    }
    mapMutex.Unlock()
}

func (c *Config) DeleteGitDirectory(path string) {
    mapMutex.Lock()
    if _, ok := c.Git[path]; ok {
        shell.Println("[ git   ] {{ " + path + " | red }} removed")
        delete(c.Git, path)
    }
    mapMutex.Unlock()
}
