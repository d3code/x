package prompt

import (
    "fmt"
    "strings"

    "github.com/d3code/pkg/xerr"
    "github.com/manifoldco/promptui"
)

func String(label string, required bool) string {
    validate := func(input string) error {
        if len(input) <= 0 && required {
            return fmt.Errorf("please enter a value")
        }
        return nil
    }

    templates := &promptui.PromptTemplates{
        Prompt:  "{{ . }} ",
        Valid:   "{{ . | green }} ",
        Invalid: "{{ . }} ",
        Success: "{{ . | blue }} ",
    }

    prompt := promptui.Prompt{
        Label:     label,
        Templates: templates,
        Validate:  validate,
        Stdout:    NoBellStdout,
    }

    result, err := prompt.Run()
    xerr.ExitIfError(err)

    return result
}

func Select(label string, items []string) (int, string) {
    prompt := promptui.Select{
        Label:  label,
        Items:  items,
        Stdout: NoBellStdout,
        Searcher: func(input string, index int) bool {
            item := items[index]
            if strings.Contains(strings.ToLower(item), strings.ToLower(input)) {
                return true
            }
            return false
        },

        Templates: &promptui.SelectTemplates{
            Label:    "  {{ . }}",
            Active:   "  {{ . | green }}",
            Inactive: "  {{ . }}",
        },
        HideHelp: true,
    }

    index, result, err := prompt.Run()
    xerr.ExitIfError(err)

    return index, result
}

// Confirm prompts the user to confirm a question with a yes or no answer.
// It displays a select prompt with the given label and returns true if the user selects "yes",
// and false if the user selects "no".
func Confirm(label string) bool {
    prompt := promptui.Select{
        Label:  label,
        Items:  []string{"yes", "no"},
        Stdout: NoBellStdout,
        Templates: &promptui.SelectTemplates{
            Label:    "  {{ . }}",
            Active:   "  {{ . | green }}",
            Inactive: "  {{ . }}",
        },
        HideHelp:     true,
        HideSelected: true,
    }

    _, result, err := prompt.Run()
    xerr.ExitIfError(err)

    return result == "yes"
}
