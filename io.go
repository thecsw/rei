package rei

import "io"

type WriteCloser struct {
	io.Writer
}

func (WriteCloser) Close() error {
	return nil
}

type ReadCloser struct {
	io.Reader
}

func (ReadCloser) Close() error {
	return nil
}
