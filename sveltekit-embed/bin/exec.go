package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

var green = color.New(color.FgGreen)

func main() {
	if len(os.Args) < 2 {
		info()
		return
	}

	command := os.Args[1]

	switch command {
	case "info":
		info()
	case "go-build":
		goBuild()
	case "svelte-build":
		svelteBuild()
	case "all":
		all()
	default:
		fmt.Println("Invalid command.")
	}
}

func all() {
	svelteBuild()
	goBuild()
}

func installTemplateDependencies() {
	// Check if "pnpm" command is available, if not, use "npm"
	if hasCommand("pnpm") {
		execCmd("pnpm", "install", "-C", "./template")
	} else {
		execCmd("npm", "install", "--prefix", "./template")
	}
}

func goBuild() {
	green.Println("=== Building Golang Project ===")
	execCmd("go", "build", "-o", "app", "-v")
}

func svelteBuild() {
	green.Println("=== Building SvelteKit Project ===")
	installTemplateDependencies()
	// Check if "pnpm" command is available, if not, use "npm"
	if hasCommand("pnpm") {
		execCmd("pnpm", "run", "-C", "./template", "build")
	} else {
		execCmd("npm", "run", "build", "--prefix", "./template")
	}
}

func info() {
	green := color.New(color.FgGreen).PrintfFunc()
	bold := color.New(color.Bold).PrintlnFunc()

	SPACES := strings.Repeat(" ", 2)

	green("------------------------------------------\n")
	green("-           SvelteKit Embed App          -\n")
	green("------------------------------------------\n")
	fmt.Println("This Bash helps you manage your projects.")
	fmt.Println()
	fmt.Println("Available commands:")
	bold("- go-build:" + SPACES + "Build the Golang project.")
	bold("- svelte-build:" + SPACES + "Build the SvelteKit project.")
	bold("- all:" + SPACES + "Run all commands (SvelteBuild, GoBuild).")
	fmt.Println()
	fmt.Println("Usage: go run ./bin <command>")
}

// Check if a command exists
func hasCommand(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// Execute a command
func execCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
