package ioadapters

import "os"

type ReaderAdapter struct{}

func (r *ReaderAdapter) ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

type WriterAdapter struct{}

func (w *WriterAdapter) WriteFile(filePath string, data []byte) {
	if err := os.WriteFile(filePath, data, os.ModeAppend); err != nil {
		panic(err)
	}
}
