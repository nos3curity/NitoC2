package main

import (
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

// Global commands list
var commands = []prompt.Suggest{
	{Text: "initialize", Description: "Create the Nito database"},
	{Text: "beacons", Description: "Control running payloads"},
	{Text: "payloads", Description: "Build payload executables"},
	{Text: "listeners", Description: "Create beacon listeners"},
}

// Complete prompt commands and subcommands
func completer(input prompt.Document) []prompt.Suggest {

	args := input.CurrentLineBeforeCursor()
	argsList := strings.Split(args, " ")

	if len(argsList) <= 1 {
		return prompt.FilterHasPrefix(commands, args, true)
	} else {
		switch argsList[0] {
		case "beacons":
			var subcommands = []prompt.Suggest{
				{Text: "list", Description: "Show all active beacons"},
				{Text: "shell", Description: "Enter interactive shell mode"},
				{Text: "kill", Description: "Ask beacon to self-destruct"},
			}
			return prompt.FilterHasPrefix(subcommands, argsList[1], true)
		case "payloads":
			var subcommands = []prompt.Suggest{
				{Text: "list", Description: "Show all built executables"},
				{Text: "build", Description: "Build a new executable"},
				{Text: "copy", Description: "Move an executable to a folder"},
			}
			return prompt.FilterHasPrefix(subcommands, argsList[1], true)
		case "listeners":
			var subcommands = []prompt.Suggest{
				{Text: "list", Description: "Show all active listeners"},
				{Text: "start", Description: "Create a new listener"},
				{Text: "stop", Description: "Stop a running listener"},
			}
			return prompt.FilterHasPrefix(subcommands, argsList[1], true)
		}
	}
	return []prompt.Suggest{}
}
