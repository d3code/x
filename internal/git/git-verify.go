package git

import "github.com/d3code/x/internal/cfg"

func VerifyPaths() {
    config := cfg.Configuration()
    for path, _ := range config.Git {
        if !Is(path) {
            config.DeleteGitDirectory(path)
        }
    }
}
