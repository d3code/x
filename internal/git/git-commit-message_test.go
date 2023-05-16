package git

import (
    "github.com/d3code/clog"
    "github.com/d3code/pkg/slice_utils"
    "testing"
)

func TestGenerateCommitMessage(t *testing.T) {
    type args struct {
        status string
    }
    tests := []struct {
        name string
        args args
        want []string
    }{
        {name: "Delete", args: struct{ status string }{status: "D  cmd/test.go"}, want: messageDelete},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            message := GenerateCommitMessage(tt.args.status)
            clog.InfoF("  [ %s ]\n", message)
            if got := message; !slice_utils.ContainsString(tt.want, got) {
                t.Errorf("GenerateCommitMessage() = %v, want %v", got, tt.want)
            }
        })
    }
}
