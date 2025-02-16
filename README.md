# Tree Structure Generator 🌳

A command-line utility written in Go that creates directory structures from tree-like input. Perfect for quickly setting up project templates!

## Features ✨
- 🖥️ Cross-platform support (Windows/macOS/Linux)
- 📂 Creates nested directories and files in one go
- 🌲 Works with standard tree structure notation
- 🚫 No external dependencies
- ✅ Progress feedback during creation

## Getting Started

### Prerequisites
- [Go 1.16+](https://golang.org/dl/) (only required for building from source)

### Installation 📦

#### Using pre-built binary (recommended):
```bash
# Install directly using Go
go install github.com/iiviie/treedo@latest
```

#### Manual build:
```bash
git clone https://github.com/iiviie/treedo.git
cd treedo
go build -o treedo main.go

# Move binary to PATH (optional)
sudo mv treedo /usr/local/bin/
```

This setup will let users:
1. Install with `make install` for local development
2. Use pre-built binaries from GitHub releases
3. Run as `treedo` command globally

For immediate use without installation:
```bash
go run github.com/iiviie/treedo@latest
```

4. Verify installation:
```bash
treedo --version
```

> **Note**: If you can't run the `treedo` command after installation, it's likely because your Go bin directory isn't in your PATH. Follow steps 1-3 above to fix this.

#### Troubleshooting 🔧
- Command not found? Make sure your Go bin directory is in your PATH
- Not sure which shell you're using? Run `echo $SHELL`
- Still having issues? Run `echo $PATH` to check your current PATH

3. Reload your shell configuration:
```bash
source ~/.bashrc  # or source ~/.zshrc if using zsh
```

## Usage 🛠️

### Basic usage:
```bash
treedo
```

1. Enter your tree structure when prompted
2. Press Enter after your last line
3. Submit with:
   - **Unix/macOS:** `Ctrl+D` (press once if at end of line, twice if on new line)
   - **Windows:** `Ctrl+Z` then Enter

### Example Input:
```
.
├── my_project/
│   ├── src/
│   │   ├── main.go
│   │   └── utils/
│   │       └── helper.go
│   └── README.md
└── .gitignore
```

### Expected Output:
```
Created directory: ./my_project
Created directory: ./my_project/src
Created file: ./my_project/src/main.go
Created directory: ./my_project/src/utils
Created file: ./my_project/src/utils/helper.go
Created file: ./my_project/README.md
Created file: ./.gitignore

✨ Directory structure created successfully!
```

## How to Format Your Tree 🌴

### Directory Notation:
```bash
directory/      # With trailing slash (recommended)
another-dir     # Without trailing slash (must have children)
```

### File Notation:
```bash
file.txt        # Simple file
.config         # Hidden file
```

### Tips 💡
- Start with `.` as root
- Use standard tree symbols: `├──`, `└──`, `│`
- Indent with spaces (4 spaces per level recommended)
- Add trailing `/` to explicitly mark directories
- Empty directories need at least one child or trailing slash

## Examples 🚀

### Simple Structure:
```bash
.
├── docs/
│   └── README.md
└── src/
    ├── main.py
    └── utils/
```


## Troubleshooting 🔧

### Common Issues:
- **Permission denied**: Run with `sudo` (Unix) or Admin privileges (Windows)
- **Empty directories**: Add trailing `/` or create a dummy file
- **Incorrect symbols**: Use only `├──`, `└──`, and `│` characters
- **Partial creation**: Check parent directory permissions

### Error Messages:
- `Error reading input`: Invalid input format
- `Failed to create directory`: Permission or path issues
- `No input provided`: Submit empty input

After installation, make sure the Go bin directory is in your PATH:

1. Find your Go bin directory:
```bash
go env GOPATH
```

2. Add this line to your shell configuration file (`~/.bashrc`, `~/.zshrc`, or similar):
```bash
# Replace /path/to/go with the output from 'go env GOPATH'
export PATH="$PATH:$(go env GOPATH)/bin"
```

3. Reload your shell configuration:
```bash
source ~/.bashrc  # or source ~/.zshrc if using zsh
```

4. Verify installation:
```bash
treedo --version
```

> **Note**: If you can't run the `treedo` command after installation, it's likely because your Go bin directory isn't in your PATH. Follow steps 1-3 above to fix this.

#### Troubleshooting 🔧
- Command not found? Make sure your Go bin directory is in your PATH
- Not sure which shell you're using? Run `echo $SHELL`
- Still having issues? Run `echo $PATH` to check your current PATH


