package cfg

import (
    "github.com/d3code/clog"
)

func (c *Config) AddTerraform(path string, config Terraform) {
    mapMutex.Lock()
    if _, ok := c.Terraform[path]; !ok {
        clog.Info("[ tf ] {{ " + path + " | blue }} added")
    }
    c.Terraform[path] = config
    c.Save()
    mapMutex.Unlock()
}

func (c *Config) DeleteTerraform(path string) {
    mapMutex.Lock()
    if _, ok := c.Terraform[path]; ok {
        clog.Info("[ tf ] {{ " + path + " | red }} removed")
        delete(c.Terraform, path)
    }
    c.Save()
    mapMutex.Unlock()
}
