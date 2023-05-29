package git

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/files"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/pkg/slice_utils"
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/gpt"
    "github.com/d3code/x/internal/input"
    "os"
    "strings"
)

func CommitDirectory(repository string, interactive bool) {

    if !Is(repository) {
        clog.Warn("Not a git repository")
        repositories := slice_utils.Keys(cfg.Configuration().Git)
        inConfig := slice_utils.ContainsString(repositories, repository)

        if interactive && inConfig && input.PromptConfirm("Remove directory from config?") {
            cfg.Configuration().DeleteGitDirectory(repository)
        } else if !interactive {
            cfg.Configuration().DeleteGitDirectory(repository)
        }

        return
    }

    commitMessage, changes := gpt.GenerateCommitMessage(repository)
    if !changes {
        clog.Info("No changes detected")
        return
    } else if commitMessage == "" {
        clog.Warn("No commit message provided, changes not committed")
        return
    }

    if interactive {
        clog.Info(commitMessage, "\n")
        if !input.PromptConfirm("Commit changes?") {
            return
        }
    }

    StageCommit(repository, commitMessage)
}

func StageCommit(path string, commitMessage string) {
    path = shell.FullPath(path)
    clog.Debug(path)

    if !Is(path) {
        clog.Error(path, "is not a git repository")
        return
    }

    file := path + "/.gitignore"
    if !files.Exist(file) && input.PromptConfirm("Create {{ .gitignore | green }}?") {
        clog.InfoF("Creating {{ .gitignore | green }} with defaults at {{ %s | blue }}", file)
        GitignoreCreate(path)
    }

    Stage(path)
    Commit(path, commitMessage)
}

func FetchPullPush(path string) {
    path = shell.FullPath(path)
    clog.Debug(path)

    if !Is(path) {
        clog.Error(path, "is not a git repository")
        return
    }

    shell.RunCmd(path, false, "git", "fetch", "-n")

    Pull(path)
    Push(path)
}

func Pull(path string) {
    e, err := shell.RunCmdE(path, false, "git", "pull", "--ff-only")

    if strings.Contains(e.Stderr, "no such ref was fetched") {
        clog.Debug("Branch does not exist on remote")
    } else if err != nil {
        clog.Error(e.Stderr)
        os.Exit(1)
    }
}

func Push(path string) {
    branch := shell.RunCmd(path, false, "git", "branch", "--show-current")
    shell.RunCmd(path, true, "git", "push", "-u", "origin", branch.Stdout)
}

func Stage(path string) (bool, string) {
    status := shell.RunCmd(path, false, "git", "status", "--porcelain")
    if len(status.Stdout) == 0 {
        clog.Debug("No unstaged changes")
        return false, status.Stdout
    }

    shell.RunCmd(path, false, "git", "add", ".")
    return true, status.Stdout
}
