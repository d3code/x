package cfg

import (
    "github.com/d3code/pkg/xerr"
    "github.com/d3code/x/internal/prompt"
    "os"
    "path"
)

func BaseDirectory() string {
    configuration := Configuration()
    baseDirectory := configuration.Environment["basedir"]

    dir, err := os.UserHomeDir()
    xerr.ExitIfError(err)

    if baseDirectory == "" {
        defaultPath := path.Join(dir, "Development", "repo")
        if prompt.Confirm("Use default base directory [" + defaultPath + "]") {
            baseDirectory = defaultPath

            configuration.Environment["basedir"] = defaultPath
            configuration.Save()
        } else {
            baseDirectory = prompt.String("Base directory", true)
        }
    }

    return baseDirectory
}
