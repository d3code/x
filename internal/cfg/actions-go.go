package cfg

import (
    "github.com/d3code/pkg/shell"
)

func (c *Config) AddGolang(path string, goConfig Golang) {
    mapMutex.Lock()
    if _, ok := c.Golang[path]; !ok {
        shell.Println("[ go    ] {{ " + path + " | blue }} added")
    }
    c.Golang[path] = goConfig
    mapMutex.Unlock()
}

func (c *Config) DeleteGolang(path string) {
    mapMutex.Lock()
    if _, ok := c.Golang[path]; ok {
        shell.Println("[ go    ] {{ " + path + " | red }} removed")
        delete(c.Golang, path)
    }
    mapMutex.Unlock()
}
