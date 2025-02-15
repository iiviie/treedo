package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Node struct {
	name     string
	isFile   bool
	children []*Node
}

func parseTreeLine(line string) (int, string, bool) {
	// Count indentation based on tree characters
	indent := 0
	for i, char := range line {
		if char == ' ' || char == '│' || char == '├' || char == '└' || char == '─' {
			indent = i + 1
		} else {
			break
		}
	}

	// Extract clean name
	name := strings.TrimSpace(line[indent:])
	name = strings.TrimPrefix(name, "├── ")
	name = strings.TrimPrefix(name, "└── ")
	
	// Determine if directory or file
	isFile := !strings.HasSuffix(name, "/")
	name = strings.TrimSuffix(name, "/")

	return indent / 4, name, isFile
}

func buildTree(lines []string) *Node {
	root := &Node{name: ".", isFile: false}
	stack := []*Node{root}
	levels := []int{-1}

	for _, line := range lines {
		if strings.TrimSpace(line) == "." || strings.TrimSpace(line) == "" {
			continue
		}

		level, name, isFile := parseTreeLine(line)

		for len(levels) > 1 && levels[len(levels)-1] >= level {
			levels = levels[:len(levels)-1]
			stack = stack[:len(stack)-1]
		}

		node := &Node{
			name:   name,
			isFile: isFile,
		}

		parent := stack[len(stack)-1]
		parent.children = append(parent.children, node)

		stack = append(stack, node)
		levels = append(levels, level)
	}

	return root
}

func createFileSystem(node *Node, basePath string) error {
	path := filepath.Join(basePath, node.name)

	if node.isFile {
		file, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %v", path, err)
		}
		file.Close()
		fmt.Printf("Created file: %s\n", path)
	} else {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %v", path, err)
		}
		fmt.Printf("Created directory: %s\n", path)
	}

	for _, child := range node.children {
		err := createFileSystem(child, path)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	fmt.Println("Enter your tree structure (press Ctrl+D when done):")
	fmt.Println("Example format:")
	fmt.Println(".")
	fmt.Println("├── src/")
	fmt.Println("│   ├── components/")
	fmt.Println("│   │   ├── Header.js")
	fmt.Println("│   │   └── Footer.js")
	fmt.Println("│   └── utils/")
	fmt.Println("└── package.json")

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}

	root := buildTree(lines)
	
	if err := createFileSystem(root, "."); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file system: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Directory structure created successfully!")
}