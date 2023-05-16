package cfg

import (
    "github.com/d3code/clog"
)

func (c *Config) AddGolang(path string, goConfig Golang) {
    mapMutex.Lock()
    if _, ok := c.Golang[path]; !ok {
        clog.Info("[ go    ] {{ " + path + " | blue }} added")
    }
    c.Golang[path] = goConfig
    c.Save()
    mapMutex.Unlock()
}

func (c *Config) DeleteGolang(path string) {
    mapMutex.Lock()
    if _, ok := c.Golang[path]; ok {
        clog.Info("[ go    ] {{ " + path + " | red }} removed")
        delete(c.Golang, path)
    }
    c.Save()
    mapMutex.Unlock()
}
