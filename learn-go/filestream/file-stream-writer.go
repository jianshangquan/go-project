package filestream

import (
	"log"
	"os"
)

type FileStreamWriter struct {
	file *os.File
}

func NewFileStreamWriter(path string) *FileStreamWriter {
	var stream = &FileStreamWriter{}
	stream.load(path)
	return stream
}

func (f *FileStreamWriter) load(path string) *FileStreamWriter {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	f.file = file
	return f
}

func (f *FileStreamWriter) getFile() *os.File {
	return f.file
}

func (f *FileStreamWriter) Close() {
	// Close the file
	if err := f.file.Close(); err != nil {
		log.Fatal(err)
	}
}
