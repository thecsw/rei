package rei

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Untar takes a reader of tar bytes and untars it to the specified directory.
// It returns an error if any.
func Untar(reader io.Reader, dirName string) error {
	tarReader := tar.NewReader(reader)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("reading tar bytes: %v", err)
		}
		path, err := SanitizeArchivePath(dirName, header.Name)
		if err != nil {
			return fmt.Errorf("sanitizing archive path: %v", err)
		}
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return fmt.Errorf("creating nested directories: %v", err)
			}
			continue
		}

		file, err := os.OpenFile(
			filepath.Clean(path),
			os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
			info.Mode())
		if err != nil {
			return fmt.Errorf("creating files: %v", err)
		}
		defer func() {
			if err := file.Close(); err != nil {
				log.Printf("closing untarred file: %v", err)
			}
		}()
		for {
			_, err = io.CopyN(file, tarReader, 1024)
			if err != nil {
				if err == io.EOF {
					break
				}
				return fmt.Errorf("copying tar files: %v", err)
			}
		}
	}
	return nil
}

// Sanitize archive file pathing from "G305: Zip Slip vulnerability".
// Found in https://github.com/securego/gosec/issues/324
func SanitizeArchivePath(d, t string) (v string, err error) {
	v = filepath.Join(d, t)
	if strings.HasPrefix(v, filepath.Clean(d)) {
		return v, nil
	}
	return "", fmt.Errorf("content filepath is tainted: %s", t)
}
