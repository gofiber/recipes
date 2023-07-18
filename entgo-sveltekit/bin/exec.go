package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func GenerateEnt() {
	green.Println("=== Generating Entity Files ===")
	execCmd("go", "run", "-mod=mod", "entgo.io/ent/cmd/ent", "generate", "./schema", "--target", "./entity")
}

func All() {
	GenerateEnt()
	SvelteBuild()
	GoTest()
	GoBuild()
}

// Golang Commands
func GoRun() {
	green.Println("=== Running Golang Project ===")
	execCmd("go", "run", "main.go")
}

func GoBuild() {
	green.Println("=== Building Golang Project ===")
	execCmd("go", "build", "-o", "app", "-v")
}

func GoTest() {
	green.Println("=== Running Golang Tests ===")
	execCmd("go", "test", "./handler")
}

func installTemplateDependencies() {
	if hasCommand("pnpm") {
		execCmd("pnpm", "install", "-C", "./template")
	} else {
		execCmd("npm", "install", "--prefix", "./template")
	}
}

// SvelteKit Commands
func SvelteRun() {
	green.Println("=== Running SvelteKit Project ===")
	installTemplateDependencies()
	if hasCommand("pnpm") {
		execCmd("pnpm", "run", "-C", "./template ", "run")
	} else {
		execCmd("npm", "run", "dev", "--prefix ", "./template")
	}
}

func SvelteBuild() {
	green.Println("=== Building SvelteKit Project ===")
	installTemplateDependencies()
	if hasCommand("pnpm") {
		execCmd("pnpm", "run", "-C", "./template ", "build")
	} else {
		execCmd("npm", "run", "build", "--prefix ", "./template")
	}
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

// Check if a command exists
func hasCommand(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func main() {
	if len(os.Args) < 2 {
		Info()
		return
	}

	command := os.Args[1]

	switch command {
	case "info":
		Info()
	case "go-run":
		GoRun()
	case "go-build":
		GoBuild()
	case "go-test":
		GoTest()
	case "svelte-run":
		SvelteRun()
	case "svelte-build":
		SvelteBuild()
	case "generate-ent":
		GenerateEnt()
	case "all":
		All()
	default:
		fmt.Println("Invalid command.")
	}
}

func Info() {
	green := color.New(color.FgGreen).PrintfFunc()
	bold := color.New(color.Bold).PrintlnFunc()

	SPACES := strings.Repeat(" ", 2)

	green("------------------------------------------\n")
	green("-   Entgo and SvelteKit Full Stack App   -\n")
	green("------------------------------------------\n")
	fmt.Println("This Bash helps you manage your projects.")
	fmt.Println()
	fmt.Println("Available commands:")
	bold("- go-run:" + SPACES + "Run the Golang project.")
	bold("- go-build:" + SPACES + "Build the Golang project.")
	bold("- go-test:" + SPACES + "Run tests for the Golang project.")
	bold("- svelte-run:" + SPACES + "Run the SvelteKit project.")
	bold("- svelte-build:" + SPACES + "Build the SvelteKit project.")
	bold("- github:" + SPACES + "Open your GitHub profile in a browser.")
	bold("- generate-ent:" + SPACES + "Generate entity files.")
	bold("- all:" + SPACES + "Run all commands (GenerateEnt, SvelteBuild, GoTest, GoBuild).")
	fmt.Println()
	fmt.Println("Usage: go run ./bin <command>")
}

var green = color.New(color.FgGreen)

// bold  = color.New(color.Bold)
