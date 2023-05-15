package git_c

import "testing"

func TestClonedDirectory(t *testing.T) {
    type args struct {
        message string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {name: "Success", args: struct{ message string }{message: "Cloning into 'xauth'..."}, want: "xauth"},
        {name: "Empty directory", args: struct{ message string }{message: "Cloning into ''..."}, want: ""},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := clonedDirectory(tt.args.message); got != tt.want {
                t.Errorf("ClonedDirectory() = %v, want %v", got, tt.want)
            }
        })
    }
}
