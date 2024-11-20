package counter

import (
	"os"
	"path/filepath"
)

func FindGitignore(startPath string) (string, error) {
	currentPath, err := filepath.Abs(startPath)
	if err != nil {
		return "", err
	}

	for {
		gitignorePath := filepath.Join(currentPath, ".gitignore")
		if _, err := os.Stat(gitignorePath); err == nil {
			return gitignorePath, nil
		}

		// Move up one directory
		parentPath := filepath.Dir(currentPath)
		if parentPath == currentPath {
			// We've reached the root directory
			return "", nil
		}
		currentPath = parentPath
	}
}

func GetFilesWithGitignore(extension, gitignorePath string) ([]string, error) {
	// Use git ls-files command
	// Implementation here would use os/exec to run git commands
	return nil, nil
}

func GetFilesWithoutGitignore(extension string) ([]string, error) {
	var files []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == "."+extension {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
