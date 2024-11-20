package counter

import (
	"bufio"
	"os"
)

func CountLines(extension string) (map[string]int, error) {
	// Initialize result map
	counts := make(map[string]int)

	// Get .gitignore path
	gitignorePath, err := FindGitignore(".")
	if err != nil {
		return nil, err
	}

	// Get files to process
	var files []string
	if gitignorePath != "" {
		files, err = GetFilesWithGitignore(extension, gitignorePath)
	} else {
		files, err = GetFilesWithoutGitignore(extension)
	}
	if err != nil {
		return nil, err
	}

	// Count lines in each file
	for _, file := range files {
		count, err := countLinesInFile(file)
		if err != nil {
			return nil, err
		}
		counts[file] = count
	}

	return counts, nil
}

func countLinesInFile(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	return lineCount, scanner.Err()
}
