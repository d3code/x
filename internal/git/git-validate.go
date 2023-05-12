package git

import "github.com/d3code/x/internal/cfg"

func Validate() {
    config := cfg.Configuration()
    for path, _ := range config.Git {
        if !IsGitDirectory(path) {
            config.DeleteGitDirectory(path)
        }
    }
    config.Save()
}
