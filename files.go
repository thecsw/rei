package rei

import (
	"errors"
	"os"
	"path/filepath"
)

// FileMustExist checks if file FileExists. Panics if not.
func FileMustExist(path string) bool {
	exists, err := FileExists(path)
	Try(err)
	return exists
}

// FileExists checks if file FileExists.
// Returns true if FileExists, false if not and error if error happened.
func FileExists(path string) (bool, error) {
	_, err := os.Stat(filepath.Clean(path))
	// exists
	if err == nil {
		return true, nil
	}
	// doesn't exist
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	// error happened
	return false, err
}

// Mkdir creates a directory(ies if needed) and reports fatal errors.
func Mkdir(path string) error {
	// Make sure that the vendor directory exists.
	err := os.MkdirAll(string(path), 0750)
	// If we couldn't create the vendor directory and it doesn't
	// exist, then turn off the vendor option.
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}
