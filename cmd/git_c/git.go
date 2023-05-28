package git_c

import (
	"github.com/spf13/cobra"
)

var Git = &cobra.Command{
	Use:     "git",
	Aliases: []string{"g"},
}
