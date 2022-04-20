package main

import (
	prompt "github.com/c-bata/go-prompt"
)

func main() {
	// Prompt Configuration
	p := prompt.New(
		executor,
		completer,

		// Prompt text and title configuration
		prompt.OptionTitle("Nito C2"),
		prompt.OptionPrefix("nito >>> "),
		prompt.OptionPrefixTextColor(prompt.Red),

		// Prompt suggestion configuration
		prompt.OptionPreviewSuggestionTextColor(prompt.Yellow),
		prompt.OptionSuggestionTextColor(prompt.Black),
		prompt.OptionSuggestionBGColor(prompt.LightGray),
		prompt.OptionSelectedSuggestionTextColor(prompt.White),
		prompt.OptionSelectedSuggestionBGColor(prompt.DarkGray),

		// Prompt description configuration
		prompt.OptionDescriptionTextColor(prompt.White),
		prompt.OptionDescriptionBGColor(prompt.Red),
		prompt.OptionSelectedDescriptionBGColor(prompt.DarkRed),
	)
	p.Run()
}
