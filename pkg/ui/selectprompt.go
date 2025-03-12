package ui

import "github.com/charmbracelet/huh"

func SelectPrompt(promptsList []string) (string, error) {
	options := make([]huh.Option[string], len(promptsList))
	for i, prompt := range promptsList {
		options[i] = huh.NewOption("/"+prompt, prompt)
	}

	var response string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select prompt").
				Options(
					options...,
				).
				Value(&response),
		),
	)

	err := form.Run()
	if err != nil {
		if err == huh.ErrUserAborted {
			return "", nil
		}
		return "", err
	}

	return response, nil
}
