package input

import (
    "fmt"

    "github.com/d3code/pkg/xerr"
    "github.com/manifoldco/promptui"
)

func PromptString(label string, required bool) string {
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

func PromptSelect(label string, items []string) (int, string) {
    prompt := promptui.Select{
        Label:    label,
        Items:    items,
        Stdout:   NoBellStdout,
        HideHelp: true,
    }

    index, result, err := prompt.Run()
    xerr.ExitIfError(err)

    return index, result
}

// PromptConfirm prompts the user to confirm a question with a yes or no answer.
// It displays a select prompt with the given label and returns true if the user selects "yes",
// and false if the user selects "no".
func PromptConfirm(label string) bool {
    prompt := promptui.Select{
        Label:        label,
        Items:        []string{"yes", "no"},
        Stdout:       NoBellStdout,
        HideHelp:     true,
        HideSelected: true,
    }

    _, result, err := prompt.Run()
    xerr.ExitIfError(err)
    return result == "yes"
}
