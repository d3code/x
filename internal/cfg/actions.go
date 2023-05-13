package cfg

import (
    "encoding/json"
    "fmt"
    "github.com/d3code/pkg/xerr"
    "os"
)

func Configuration() *Config {
    return localConfig
}

func (c *Config) Save() {
    file := fmt.Sprintf("%s/config.json", configurationPath())

    configJson, err := json.MarshalIndent(c, "", "  ")
    xerr.ExitIfError(err)

    err = os.WriteFile(file, configJson, 0666)
    xerr.ExitIfError(err)
}
