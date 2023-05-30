package cfg

import (
    "bytes"
    "fmt"
    "github.com/d3code/clog"
    "github.com/d3code/pkg/files"
    "github.com/d3code/pkg/xerr"
    "gopkg.in/yaml.v3"
    "os"
    "sync"
)

var (
    localConfig *Config
    mapMutex    sync.RWMutex

    configFilePath = fmt.Sprintf("%s/config.yaml", ConfigPath())
)

func init() {
    if !files.Exist(configFilePath) {
        localConfig = &Config{}
        localConfig.Save()
    }

    indexByteArray, err := os.ReadFile(configFilePath)
    xerr.ExitIfError(err)

    err = yaml.Unmarshal(indexByteArray, &localConfig)
    xerr.ExitIfError(err)

    localConfig.Save()
}

type Config struct {
    GitHub      map[string]GitHub    `yaml:"github"`
    Git         map[string]Git       `yaml:"git"`
    Golang      map[string]Golang    `yaml:"go"`
    Terraform   map[string]Terraform `yaml:"terraform"`
    Angular     map[string]Angular   `yaml:"angular"`
    Environment map[string]string    `yaml:"environment"`
}

type Git struct {
    Remote string `yaml:"remote"`
}

type GitHub struct {
    Token string `yaml:"token"`
}

type Golang struct {
    Module string `yaml:"module"`
}

type Terraform struct {
}

type Angular struct {
}

func ConfigPath() string {
    dir, err := os.UserHomeDir()
    xerr.ExitIfError(err)

    dir = dir + "/.x"
    if !files.Exist(dir) {
        err = os.MkdirAll(dir, 0755)
        xerr.ExitIfError(err)
    }

    return dir
}
func Configuration() *Config {
    return localConfig
}

func (c *Config) Save() {
    var buffer bytes.Buffer
    yamlEncoder := yaml.NewEncoder(&buffer)
    yamlEncoder.SetIndent(4)

    err := yamlEncoder.Encode(c)
    xerr.ExitIfError(err)

    err = os.WriteFile(configFilePath, buffer.Bytes(), 0666)
    xerr.ExitIfError(err)

    clog.Debug("Saved configuration")
}
