package main

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	athena "main.go/artefacts" // My  own packages or functions
)

func main() {
	banner()
	// If no arguments, show
	if len(os.Args) < 3 {
		fmt.Println("No arguments provided")
		fmt.Println("Use -h to show help menu")
		os.Exit(0)
	}

	comandos := os.Args[1]

	switch comandos {
	case "-h", "--heelp":
		help()
	case "-r", "--read":
		if len(os.Args) < 3 {
			fmt.Println("No file provided")
			os.Exit(0)
		}
		path := os.Args[2]
		headers, err := athena.ReadFile(path)
		if err != nil {
			fmt.Println("[!]Error, exiting", err)
			os.Exit(1)
		}
		athena.PrintHeaders(headers)
	}

}

func banner() {
	myFigure := figure.NewColorFigure("EML Checker", "doom", "green", true)
	myFigure.Print()
	fmt.Println("\nBy: @AreiaNight")
}

func help() {
	fmt.Println("Usage: emlcheck -r [file.eml]")
}
