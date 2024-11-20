# countlines

A fast and efficient command-line tool to count lines of code in your projects. Respects `.gitignore` rules and supports various file extensions.

## Features

- 📂 Counts lines in files with specific extensions
- 🚫 Respects `.gitignore` rules automatically
- ⚡️ Fast and lightweight
- 🔍 Recursive directory search
- 📊 Clear, formatted output

## Installation

```bash
go install github.com/richardamare/countlines@latest
```

## Usage

Basic usage:
```bash
countlines [extension]
```

Examples:
```bash
countlines go     # Count lines in Go files
countlines js     # Count lines in JavaScript files
countlines py     # Count lines in Python files
```

Output example:
```bash
$ countlines go
     124    cmd/root.go
      89    internal/counter/counter.go
      67    internal/counter/gitignore.go
      12    main.go

Total: 292 lines in 4 files
```

## How It Works

1. Searches for a `.gitignore` file in the current directory and parent directories
2. If found, uses git's ignore rules to exclude files
3. Counts lines in all matching files
4. Displays individual file counts and total

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see LICENSE for details