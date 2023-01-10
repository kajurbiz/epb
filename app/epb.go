package app

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func runApp() {
	// Load environment variables from .env
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Please create .env and set the POCKET_DIR environment variable to your pocketbase folder.")
		return
	}

	// Read the working directory from the POCKET_DIR environment variable
	pbDir := os.Getenv("POCKET_DIR")
	if pbDir == "" {
		fmt.Println("Please set the POCKET_DIR environment variable to your pocketbase folder.")
	}

	// Change the current pocketbase directory
	err = os.Chdir(pbDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check if Git Bash is available
	_, err = exec.LookPath("bash")
	if err != nil {
		fmt.Println("Git Bash is not available. Please download it from https://git-scm.com/downloads.")
		return
	}

	// Open a new terminal window and run the "pocketbase serve" command in it
	cmd := exec.Command("bash", "-ic", "./pocketbase serve")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
