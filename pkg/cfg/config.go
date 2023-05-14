package cfg

import (
    "encoding/json"
    "fmt"
    "github.com/d3code/pkg/files"
    "github.com/d3code/pkg/xerr"
    "os"
    "sync"
)

var localConfig *Config
var mapMutex sync.RWMutex

func init() {

    config := fmt.Sprintf("%s/config.json", configurationPath())
    if !files.Exist(config) {
        err := files.Save(configurationPath(), "config.json", []byte("{}"), true)
        xerr.ExitIfError(err)
    }

    indexByteArray, err := os.ReadFile(config)
    xerr.ExitIfError(err)

    err = json.Unmarshal(indexByteArray, &localConfig)
    xerr.ExitIfError(err)

    if localConfig.GitHub == nil {
        localConfig.GitHub = make(map[string]GitHub)
    }
    if localConfig.Git == nil {
        localConfig.Git = make(map[string]Git)
    }
    if localConfig.Golang == nil {
        localConfig.Golang = make(map[string]Golang)
    }
    if localConfig.Terraform == nil {
        localConfig.Terraform = make(map[string]Terraform)
    }
    if localConfig.Angular == nil {
        localConfig.Angular = make(map[string]Angular)
    }
    if localConfig.Docker == nil {
        localConfig.Docker = make(map[string]Docker)
    }
}

type Config struct {
    GitHub      map[string]GitHub    `json:"github"`
    Git         map[string]Git       `json:"git"`
    Golang      map[string]Golang    `json:"go"`
    Terraform   map[string]Terraform `json:"terraform"`
    Angular     map[string]Angular   `json:"angular"`
    Docker      map[string]Docker    `json:"docker"`
    Environment struct {
    } `json:"environment"`
}

type Git struct {
    Remote string `json:"remote"`
}

type GitHub struct {
    Token string `json:"token"`
}

type Golang struct {
    Name string `json:"name"`
}

type Terraform struct {
}

type Angular struct {
}

type Docker struct {
}

func configurationPath() string {
    dir, err := os.UserHomeDir()
    xerr.ExitIfError(err)

    dir = dir + "/.x"
    if !files.Exist(dir) {
        err = os.MkdirAll(dir, 0755)
        xerr.ExitIfError(err)
    }

    return dir
}
