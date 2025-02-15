package main

import (
	"fmt"
	"io"
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
	indent := 0
	for i, char := range line {
		if char == ' ' || char == '│' || char == '├' || char == '└' || char == '─' {
			indent = i + 1
		} else {
			break
		}
	}

	name := strings.TrimSpace(line[indent:])
	name = strings.TrimPrefix(name, "├── ")
	name = strings.TrimPrefix(name, "└── ")
	
	name = strings.TrimLeft(name, "│ ")
	
	isFile := true
	if strings.HasSuffix(name, "/") {
		isFile = false
		name = strings.TrimSuffix(name, "/")
	}
	
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

func markDirectories(node *Node) {
	if len(node.children) > 0 {
		node.isFile = false
		for _, child := range node.children {
			markDirectories(child)
		}
	}
}

func main() {
	fmt.Println("\n=== Tree Structure Generator ===")
	fmt.Println("\nCreate directories and files from a tree structure.")
	fmt.Println("\nInstructions:")
	fmt.Println("1. Enter your tree structure below")
	fmt.Println("2. Press Enter after each line")
	fmt.Println("3. Press Ctrl+D (Unix) or Ctrl+Z (Windows) once when done")
	fmt.Println("   Note: If using Ctrl+D, you may need to press it twice if")
	fmt.Println("         you're at the beginning of a new line")
	fmt.Println("\nExample structure:")
	fmt.Println(".")
	fmt.Println("├── project/")
	fmt.Println("│   ├── src/")
	fmt.Println("│   │   ├── main.go")
	fmt.Println("│   │   └── utils/")
	fmt.Println("│   │       └── helper.go")
	fmt.Println("│   ├── tests/")
	fmt.Println("│   │   └── main_test.go")
	fmt.Println("│   └── README.md")
	fmt.Println("└── .gitignore")
	fmt.Println("\nEnter your tree structure below:")

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	var lines []string
	for _, line := range strings.Split(string(data), "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			lines = append(lines, line)
		}
	}

	if len(lines) == 0 {
		fmt.Println("No input provided. Exiting...")
		os.Exit(1)
	}

	root := buildTree(lines)
	markDirectories(root)
	
	if err := createFileSystem(root, "."); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file system: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n✨ Directory structure created successfully!")
}