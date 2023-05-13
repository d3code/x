package cfg

import (
    "github.com/d3code/pkg/shell"
)

func (c *Config) AddTerraform(path string, config Terraform) {
    mapMutex.Lock()
    if _, ok := c.Terraform[path]; !ok {
        shell.Println("[ tf ] {{ " + path + " | blue }} added")
    }
    c.Terraform[path] = config
    mapMutex.Unlock()
}

func (c *Config) DeleteTerraform(path string) {
    mapMutex.Lock()
    if _, ok := c.Terraform[path]; ok {
        shell.Println("[ tf ] {{ " + path + " | red }} removed")
        delete(c.Terraform, path)
    }
    mapMutex.Unlock()
}
