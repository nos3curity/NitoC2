package main

import (
	"fmt"
	"os"
	"strings"
)

// Execute prompt commands
func executor(in string) {

	in = strings.TrimSpace(in)
	blocks := strings.Split(in, " ")

	switch blocks[0] {
	case "exit":
		os.Exit(0)
	case "initialize":
		initializeDatabase()
	case "beacons":
		if len(blocks) > 1 {
			switch blocks[1] {
			case "list":
				fmt.Println("[TODO] Not Implemented")
			case "shell":
				fmt.Println("[TODO] Not Implemented")
			case "kill":
				fmt.Println("[TODO] Not Implemented")
			}
		}
	case "payloads":
		if len(blocks) > 1 {
			switch blocks[1] {
			case "list":
				fmt.Println("[TODO] Not Implemented")
			case "build":
				fmt.Println("[TODO] Not Implemented")
			case "copy":
				fmt.Println("[TODO] Not Implemented")
			}
		}
	case "listeners":
		if len(blocks) > 1 {
			switch blocks[1] {
			case "list":
				getListeners()
			case "start":
				fmt.Println("[TODO] Not Implemented")
			case "stop":
				fmt.Println("[TODO] Not Implemented")
			}
		}
	}
}
