package lqfiles

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type memoryFile struct {
	data    []byte
	isDir   bool
	modTime time.Time
}

type memorySystem struct {
	files map[string]memoryFile
	mu    sync.RWMutex
}

type memoryDirEntry struct {
	name  string
	isDir bool
}

func (e *memoryDirEntry) Name() string               { return e.name }
func (e *memoryDirEntry) IsDir() bool                { return e.isDir }
func (e *memoryDirEntry) Type() os.FileMode          { return 0 }
func (e *memoryDirEntry) Info() (os.FileInfo, error) { return nil, nil }

func NewInMemoryFiles() Files {
	return &memorySystem{
		files: make(map[string]memoryFile),
	}
}

func (m *memorySystem) Write(path string, data []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	path = filepath.Clean(path)
	dir := filepath.Dir(path)
	if dir != "." && !m.existsLocked(dir) {
		return errors.New("directory does not exist")
	}

	m.files[path] = memoryFile{
		data:    data,
		modTime: time.Now(),
	}
	return nil
}

func (m *memorySystem) Read(path string) ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	path = filepath.Clean(path)
	if file, ok := m.files[path]; ok && !file.isDir {
		return file.data, nil
	}
	return nil, errors.New("file not found")
}

func (m *memorySystem) CreateDir(path string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	path = filepath.Clean(path)
	m.files[path] = memoryFile{
		isDir:   true,
		modTime: time.Now(),
	}
	return nil
}

func (m *memorySystem) Remove(path string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	path = filepath.Clean(path)
	for k := range m.files {
		if k == path || strings.HasPrefix(k, path+"/") {
			delete(m.files, k)
		}
	}
	return nil
}

func (m *memorySystem) Exists(path string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.existsLocked(path)
}

func (m *memorySystem) existsLocked(path string) bool {
	path = filepath.Clean(path)
	_, ok := m.files[path]
	return ok
}

func (m *memorySystem) ListDir(path string) ([]os.DirEntry, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	path = filepath.Clean(path)
	if !m.existsLocked(path) {
		return nil, errors.New("directory not found")
	}

	seen := make(map[string]bool)
	var entries []os.DirEntry

	for k, v := range m.files {
		if k == path {
			continue
		}
		if strings.HasPrefix(k, path+"/") {
			rel := strings.TrimPrefix(k, path+"/")
			parts := strings.Split(rel, "/")
			if len(parts) > 0 && !seen[parts[0]] {
				seen[parts[0]] = true
				entries = append(entries, &memoryDirEntry{
					name:  parts[0],
					isDir: len(parts) > 1 || v.isDir,
				})
			}
		}
	}
	return entries, nil
}

func (m *memorySystem) Copy(src, dst string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	src, dst = filepath.Clean(src), filepath.Clean(dst)
	if file, ok := m.files[src]; ok {
		m.files[dst] = memoryFile{
			data:    append([]byte{}, file.data...),
			isDir:   file.isDir,
			modTime: time.Now(),
		}
		return nil
	}
	return errors.New("source file not found")
}

func (m *memorySystem) Move(src, dst string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	src, dst = filepath.Clean(src), filepath.Clean(dst)
	if file, ok := m.files[src]; ok {
		m.files[dst] = file
		delete(m.files, src)
		return nil
	}
	return errors.New("source file not found")
}

func (m *memorySystem) IsDir(path string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	path = filepath.Clean(path)
	if file, ok := m.files[path]; ok {
		return file.isDir
	}
	return false
}
