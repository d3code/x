package cfg

import (
    "encoding/json"
    "fmt"
    "github.com/d3code/pkg/errors"
    "github.com/d3code/pkg/files"
    "os"
    "sync"
)

var localConfig *Config
var mapMutex sync.RWMutex

func init() {

    config := fmt.Sprintf("%s/config.json", configurationPath())
    if !files.Exist(config) {
        err := files.Save(configurationPath(), "config.json", []byte("{}"), true)
        errors.ExitIfError(err)
    }

    indexByteArray, err := os.ReadFile(config)
    errors.ExitIfError(err)

    err = json.Unmarshal(indexByteArray, &localConfig)
    errors.ExitIfError(err)

    if localConfig.GitHub == nil {
        localConfig.GitHub = make(map[string]GitHub)
    }
    if localConfig.Git == nil {
        localConfig.Git = make(map[string]Git)
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
    Terraform   map[string]Terraform `json:"terraform"`
    Angular     map[string]Angular   `json:"angular"`
    Docker      map[string]Docker    `json:"docker"`
    Environment struct {
        RepositoryFolder string `json:"repository_folder"`
    } `json:"environment"`
}

type Git struct {
    Remote string `json:"remote"`
}

type GitHub struct {
    Organization bool   `json:"organization"`
    Token        string `json:"token"`
}

type Terraform struct {
    Remote string `json:"remote"`
}

type Angular struct {
    Remote string `json:"remote"`
}

type Docker struct {
    Remote string `json:"remote"`
}

func configurationPath() string {
    dir, err := os.UserHomeDir()
    errors.ExitIfError(err)

    dir = dir + "/.x"
    if !files.Exist(dir) {
        err = os.MkdirAll(dir, 0755)
        errors.ExitIfError(err)
    }

    return dir
}
