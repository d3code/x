package terminal

import (
    "fmt"
    "github.com/d3code/pkg/errors"
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
        Invalid: "{{ . | red }} ",
        Success: "{{ . | blue }} ",
    }

    prompt := promptui.Prompt{
        Label:     label,
        Templates: templates,
        Validate:  validate,
        Stdout:    NoBellStdout,
    }

    result, err := prompt.Run()
    errors.ExitIfError(err)

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
    errors.ExitIfError(err)

    return index, result
}

func PromptYesNo(label string) bool {
    items := []string{"yes", "no"}
    prompt := promptui.Select{
        Label:    label,
        Items:    items,
        Stdout:   NoBellStdout,
        HideHelp: true,
    }

    _, result, err := prompt.Run()
    errors.ExitIfError(err)

    return result == "yes"
}
