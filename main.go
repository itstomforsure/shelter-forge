package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "g", "generate":
		fmt.Println("generate")
		handleGenerateSubCommand(os.Args[2:])
	default:
		return
	}

	fmt.Println("args", os.Args)
}

func handleGenerateSubCommand(cmds []string) {
	subcmd := cmds[0]
	switch subcmd {
	case "i", "interface":
		fmt.Println("interface")
		handleInterface(cmds[1:])
	case "repo":
		fmt.Println("repository")
	}
}

func handleInterface(cmds []string) {
	interfaceType := cmds[0]
	switch interfaceType {
	case "c", "controller":

	}
	fmt.Println(cmds)
}
