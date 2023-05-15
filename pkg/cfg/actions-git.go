package cfg

import (
    "github.com/d3code/clog"
)

func (c *Config) AddGitDirectory(path string, git Git) {
    mapMutex.Lock()
    if _, ok := c.Git[path]; !ok {
        clog.Debug("Added git repository " + path + " -> " + git.Remote)
        c.Git[path] = git
    }
    mapMutex.Unlock()
}

func (c *Config) DeleteGitDirectory(path string) {
    mapMutex.Lock()
    if _, ok := c.Git[path]; ok {
        clog.Warn("Removed git repository " + path)
        delete(c.Git, path)
    }
    mapMutex.Unlock()
}
