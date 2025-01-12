package ui

import (
	"github.com/charmbracelet/huh"
)

func TextInput() (string, error) {
	var result string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Placeholder("Ask a question").
				Value(&result).
				WithAccessible(true).
				WithHeight(3),
		),
	)

	err := form.Run()
	return result, err
}
