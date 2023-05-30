package cfg

import (
    "github.com/d3code/clog"
)

func (c *Config) AddGitHubUser(user string, github GitHub) {
    mapMutex.Lock()
    clog.Info("[ github ] {{ " + user + " | blue }} set")
    c.GitHub[user] = github
    c.Save()
    mapMutex.Unlock()
}

func (c *Config) DeleteGitHubUser(path string) {
    mapMutex.Lock()
    if _, ok := c.GitHub[path]; ok {
        clog.Info("[ github ] {{ " + path + " | red }} removed")
        delete(c.GitHub, path)
        c.Save()
    }
    mapMutex.Unlock()
}
