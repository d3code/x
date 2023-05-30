package prompt

import "github.com/chzyer/readline"

type stdoutQuiet struct{}

func (n *stdoutQuiet) Write(p []byte) (int, error) {
    if len(p) == 1 && p[0] == readline.CharBell {
        return 0, nil
    }
    return readline.Stdout.Write(p)
}

func (n *stdoutQuiet) Close() error {
    return readline.Stdout.Close()
}

var NoBellStdout = &stdoutQuiet{}
