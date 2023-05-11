package cfg

import (
    "github.com/d3code/pkg/shell"
)

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
