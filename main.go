package main

import (
	"fmt"
	"os"
	"strings"
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
}

func handleGenerateSubCommand(cmds []string) {
	subcmd := cmds[0]
	switch subcmd {
	case "a", "all":
		handleAll(cmds[1])
	case "i", "interface":
		fmt.Println("interface")
		handleInterface(cmds[1:])
	case "r", "repo", "repository":
		fmt.Println("repository")
		handleRepository(cmds[1])
	case "c", "controller":
		fmt.Println("controller")
	}
}

func handleAll(fileName string) {
	handleRepository(fileName)
	handleController(fileName)
}

func handleInterface(cmds []string) {
	interfaceType := cmds[0]
	switch interfaceType {
	case "r", "repo", "repository":
		fmt.Println("repository")
		generateRepositoryInterface(cmds[1])
	case "c", "controller":
		fmt.Println("controller")
	}
}

func handleRepository(fileName string) {
	generateRepositoryInterface(fileName)
	generateRepository(fileName)
}

func handleController(fileName string) {
	generateControllerInterface(fileName)
	generateController(fileName)
}

func generateRepositoryInterface(fileName string) {
	generateFile("app/interfaces/repositories/", fileName+".repository.go", "package interfaces\n"+"\ntype I"+strings.Title(fileName)+"Repository interface {}\n")
}
func generateRepository(fileName string) {
	generateFile("repositories/", fileName+".repository.go", "package repositories\n")
}

func generateControllerInterface(fileName string) {
	generateFile("app/interfaces/controllers/", fileName+".controller.go", "package controllers\n"+"\ntype I"+strings.Title(fileName)+"Controller interface {}\n")
}
func generateController(fileName string) {
	generateFile("controllers/", fileName+".controller.go", "package contollers\n")
}

func generateFile(dir string, file string, body string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0700)
	}
	f, err := os.Create(dir + file)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	l, err := f.WriteString(body)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
}
