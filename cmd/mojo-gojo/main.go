package main

import (
	"fmt"
	"github.com/EdmondTabaku/mojo-gojo/translator"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mojo-gojo <command> [directory/file]")
		os.Exit(1)
	}

	cmd := os.Args[1]
	dir := "."

	if len(os.Args) > 2 {
		dir = os.Args[2]
	}

	switch cmd {
	case "init":
		// Initialize a new Go module for package management
		modCmd := exec.Command("go", "mod", "init", dir)
		modCmd.Stdout = os.Stdout
		modCmd.Stderr = os.Stderr
		if err := modCmd.Run(); err != nil {
			fmt.Printf("Error initializing Go module: %s\n", err)
			os.Exit(1)
		}
		fmt.Println("Initialized a new amazing mojo-gojo project:", dir)

	case "build":
		generatedFiles, err := translator.TranslateDirectory(dir)
		if err != nil {
			fmt.Printf("Error translating the awesome mojo-gojo files: %s\n", err)
			os.Exit(1)
		}
		defer deleteGeneratedFiles(generatedFiles)

	case "run":
		generatedFiles, err := translator.TranslateDirectory(dir)
		if err != nil {
			fmt.Printf("Error translating the awesome mojo-gojo files: %s\n", err)
			os.Exit(1)
		}
		defer deleteGeneratedFiles(generatedFiles)

		// Now, run the Go code
		cmd := exec.Command("go", "run", dir)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error running the Go code: %s\n", err)
			os.Exit(1)
		}

	default:
		fmt.Println("Unknown command. Valid commands are: init, build, run")
		os.Exit(1)
	}
}

// deleteGeneratedFiles deletes the provided list of .go files.
func deleteGeneratedFiles(files []string) {
	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			fmt.Printf("Error deleting generated file %s: %s\n", file, err)
		}
	}
}
