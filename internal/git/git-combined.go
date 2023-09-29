package git

import (
    "fmt"
    "github.com/d3code/clog"
    "github.com/d3code/pkg/files"
    "github.com/d3code/pkg/shell"
    "github.com/d3code/pkg/slice_utils"
    "github.com/d3code/x/internal/cfg"
    "github.com/d3code/x/internal/gpt"
    "github.com/d3code/x/internal/prompt"
    "strings"
)

func CommitDirectory(repository string, interactive bool) {
    if !Git(repository) {
        clog.WarnF("Not a git repository %s", repository)
        repositories := slice_utils.Keys(cfg.Configuration().Git)
        inConfig := slice_utils.ContainsString(repositories, repository)

        if interactive && inConfig && prompt.Confirm("Remove directory from config?") {
            cfg.Configuration().DeleteGitDirectory(repository)
        } else if !interactive {
            cfg.Configuration().DeleteGitDirectory(repository)
        }

        return
    }

    commitMessage, changes := gpt.GenerateCommitMessage(repository)
    if !changes {
        clog.Debug("No changes detected")
        return
    } else if commitMessage == "" {
        clog.Warn("No commit message provided, changes not committed")
        return
    }

    clog.Info(commitMessage, "\n")

    if interactive {
        if !prompt.Confirm("Commit changes?") {
            return
        }
    }

    StageCommit(repository, commitMessage)
}

func StageCommit(path string, commitMessage string) {
    path = shell.FullPath(path)
    clog.Debug(path)

    if !Git(path) {
        clog.Error(path, "is not a git repository")
        return
    }

    file := path + "/.gitignore"
    if !files.Exist(file) && prompt.Confirm("Create {{ .gitignore | green }}?") {
        clog.Infof("Creating {{ .gitignore | green }} with defaults at {{ %s | blue }}", file)
        GitignoreCreate(path)
    }

    Stage(path)
    Commit(path, commitMessage)
}

func FetchPullPush(path string) error {
    path = shell.FullPath(path)
    clog.Debug(path)

    if !Git(path) {
        return fmt.Errorf(path, "is not a git repository")
    }

    shell.RunCmd(path, false, "git", "fetch", "-n")

    err := Pull(path)
    if err != nil {
        return err
    }

    err = Push(path)
    return err
}

func Pull(path string) error {
    e, err := shell.RunCmdE(path, false, "git", "pull", "--ff-only")

    if strings.Contains(e.Stderr, "no such ref was fetched") {
        clog.Debug("Branch does not exist on remote")
    } else if err != nil {
        return fmt.Errorf(e.Stderr)
    }

    return nil
}

func Push(path string) error {
    branch, err := shell.RunCmdE(path, false, "git", "branch", "--show-current")
    if err != nil {
        return err
    }
    out, err := shell.RunCmdE(path, false, "git", "push", "-u", "origin", branch.Stdout)
    if err != nil {
        return fmt.Errorf(out.Stderr)
    }

    return nil
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
