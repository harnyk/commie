package ui

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

type ConsentResponse string

const (
	ConsentResponseYes      ConsentResponse = "y"
	ConsentResponseNo       ConsentResponse = "n"
	ConsentResponseFollowUp ConsentResponse = "f"
)

func ShowConsent(message string) (ConsentResponse, string, error) {
	var response ConsentResponse
	var followup string

	fmt.Println(RenderMarkdown(message))

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[ConsentResponse]().
				Title("Do you consent?").
				Options(
					huh.NewOption("Yes", ConsentResponseYes),
					huh.NewOption("No (ctrl+c)", ConsentResponseNo),
					huh.NewOption("Follow up", ConsentResponseFollowUp),
				).
				Value(&response),
		),
	)

	err := form.Run()
	if err != nil {
		if err == huh.ErrUserAborted {
			return ConsentResponseNo, followup, nil
		}
		return "", followup, err
	}

	if response == ConsentResponseFollowUp {
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Follow up").
					Value(&followup),
			),
		)

		err := form.Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				return ConsentResponseNo, followup, nil
			}
			return "", followup, err
		}
	}

	return response, followup, nil
}
