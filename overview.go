package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	titleRegex    = regexp.MustCompile(`(?m)^title: (.+)`)
	keywordsRegex = regexp.MustCompile(`(?m)^keywords: \[(.+)\]`)
	headerRegex   = regexp.MustCompile(`(?m)^# .+`)
)

//go:generate go run overview.go
func main() {
	// fetch current file directory
	root, _ := os.Getwd()
	toc := ""
	var missingReadmeDirs []string
	var missingTitleDirs []string
	var missingKeywordsDirs []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != root && filepath.Dir(path) == root && !strings.HasPrefix(info.Name(), ".") {
			readmePath := filepath.Join(path, "README.md")
			relativePath, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			if _, err := os.Stat(readmePath); err == nil {
				title, keywords, err := extractTitleAndKeywords(readmePath)
				if err != nil {
					return err
				}
				if title == "" {
					missingTitleDirs = append(missingTitleDirs, relativePath)
				}
				if len(keywords) == 0 {
					missingKeywordsDirs = append(missingKeywordsDirs, relativePath)
				}
				if title == "" {
					title = "No title"
				}
				toc += fmt.Sprintf("- [%s](./%s/README.md)\n", title, relativePath)
				err = addLinksToReadme(readmePath, info.Name())
				if err != nil {
					return err
				}
			} else {
				missingReadmeDirs = append(missingReadmeDirs, relativePath)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	readmePath := filepath.Join(root, "README.md")
	content, err := os.ReadFile(readmePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	re := regexp.MustCompile(`(?s)<!-- AUTO-GENERATED-CONTENT:START -->(.*?)<!-- AUTO-GENERATED-CONTENT:END -->`)
	newContent := re.ReplaceAllString(string(content), fmt.Sprintf("<!-- AUTO-GENERATED-CONTENT:START -->\n%s<!-- AUTO-GENERATED-CONTENT:END -->", toc))

	err = os.WriteFile(readmePath, []byte(newContent), 0o644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Table of contents updated successfully.")

	if len(missingReadmeDirs) > 0 {
		fmt.Println("Directories without README.md:")
		for _, dir := range missingReadmeDirs {
			fmt.Println("-", dir)
		}
	}

	if len(missingTitleDirs) > 0 {
		fmt.Println("Directories without Docusaurus title:")
		for _, dir := range missingTitleDirs {
			fmt.Println("-", dir)
		}
	}

	if len(missingKeywordsDirs) > 0 {
		fmt.Println("Directories without Docusaurus keywords:")
		for _, dir := range missingKeywordsDirs {
			fmt.Println("-", dir)
		}
	}

	if len(missingReadmeDirs) > 0 || len(missingTitleDirs) > 0 || len(missingKeywordsDirs) > 0 {
		fmt.Println("Error: Some directories are missing README.md files, Docusaurus title, or keywords.")
		os.Exit(1)
	}
}

func extractTitleAndKeywords(readmePath string) (string, []string, error) {
	content, err := os.ReadFile(readmePath)
	if err != nil {
		return "", nil, err
	}

	titleMatches := titleRegex.FindSubmatch(content)
	keywordsMatches := keywordsRegex.FindSubmatch(content)

	var title string
	if len(titleMatches) > 1 {
		title = strings.TrimSpace(string(titleMatches[1]))
	}

	var keywords []string
	if len(keywordsMatches) > 1 {
		keywords = strings.Split(string(keywordsMatches[1]), ",")
		for i := range keywords {
			keywords[i] = strings.TrimSpace(keywords[i])
		}
	}

	return title, keywords, nil
}

func addLinksToReadme(readmePath, dirName string) error {
	content, err := os.ReadFile(readmePath)
	if err != nil {
		return err
	}

	links := fmt.Sprintf(
		"[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/%s) "+
			"[![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/%s)",
		dirName, dirName)

	lines := strings.Split(string(content), "\n")
	headerIndex := -1

	// Find the first headline
	for i, line := range lines {
		if headerRegex.MatchString(line) {
			headerIndex = i
			break
		}
	}

	if headerIndex == -1 {
		return fmt.Errorf("no headline found in %s", readmePath)
	}

	// Find the start and end of the section to be replaced
	start := headerIndex + 1
	end := len(lines)
	for i := start; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "" || strings.Contains(lines[i], "![Github]") || strings.Contains(lines[i], "![StackBlitz]") {
			continue
		}
		end = i
		break
	}

	// Replace the section with the new links and surrounding empty lines
	cleanedLines := append(lines[:start], append([]string{"", links, ""}, lines[end:]...)...)

	content = []byte(strings.Join(cleanedLines, "\n"))
	return os.WriteFile(readmePath, content, 0o644)
}
