package git

import "github.com/d3code/x/pkg/cfg"

func VerifyPaths() {
    config := cfg.Configuration()
    for path, _ := range config.Git {
        if !Git(path) {
            config.DeleteGitDirectory(path)
        }
    }
    config.Save()
}
