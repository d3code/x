package cfg

import (
    "github.com/d3code/clog"
)

func (c *Config) AddEnvironment(variable string, value string) {
    mapMutex.Lock()
    if _, ok := c.Environment[variable]; !ok {
        clog.Info("[ environment ] {{ " + variable + " | blue }} added")
    }
    c.Environment[variable] = value
    c.Save()
    mapMutex.Unlock()
}

func (c *Config) DeleteEnvironment(variable string) {
    mapMutex.Lock()
    if _, ok := c.Environment[variable]; ok {
        clog.Info("[ environment ] {{ " + variable + " | red }} removed")
        delete(c.Environment, variable)
    }
    c.Save()
    mapMutex.Unlock()
}
