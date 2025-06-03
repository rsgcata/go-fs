package fs

import (
	"github.com/rsgcata/go-fs/filelock"
	"github.com/rsgcata/go-fs/filelock/unix"
	"github.com/rsgcata/go-fs/filelock/windows"
	"runtime"
)

// New creates a new FileLock for the specified file path
func New(path string) filelock.FileLock {
	if runtime.GOOS == "windows" {
		return windows.New(path)
	}
	return unix.New(path)
}
