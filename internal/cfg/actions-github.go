package cfg

import (
    "github.com/d3code/pkg/shell"
)

func (c *Config) SetGitHubUser(user string, github GitHub) {
    mapMutex.Lock()
    shell.Println("[ github ] {{ " + user + " | blue }} set")
    c.GitHub[user] = github
    mapMutex.Unlock()
}

func (c *Config) DeleteGitHubUser(path string) {
    mapMutex.Lock()
    if _, ok := c.GitHub[path]; ok {
        shell.Println("[ github ] {{ " + path + " | red }} removed")
        delete(c.GitHub, path)
    }
    mapMutex.Unlock()
}
