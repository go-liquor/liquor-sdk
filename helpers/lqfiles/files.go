package lqfiles

import (
	"io"
	"os"
)

// Files interface provides common file system operations
type Files interface {
	// Write writes data to a file, creating it if it doesn't exist
	Write(path string, data []byte) error

	// Read reads the entire contents of a file
	Read(path string) ([]byte, error)

	// CreateDir creates a directory and all necessary parent directories
	CreateDir(path string) error

	// Remove removes a file or directory
	Remove(path string) error

	// Exists checks if a file or directory exists
	Exists(path string) bool

	// ListDir lists all files and directories in the specified path
	ListDir(path string) ([]os.DirEntry, error)

	// Copy copies a file from source to destination
	Copy(src, dst string) error

	// Move moves/renames a file from source to destination
	Move(src, dst string) error

	// IsDir checks if the path is a directory
	IsDir(path string) bool
}

type system struct{}

func NewFiles() Files {
	return &system{}
}

func (s *system) Write(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func (s *system) Read(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (s *system) CreateDir(path string) error {
	return os.MkdirAll(path, 0755)
}

func (s *system) Remove(path string) error {
	return os.RemoveAll(path)
}

func (s *system) Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (s *system) ListDir(path string) ([]os.DirEntry, error) {
	return os.ReadDir(path)
}

func (s *system) Copy(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

func (s *system) Move(src, dst string) error {
	return os.Rename(src, dst)
}

func (s *system) IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
