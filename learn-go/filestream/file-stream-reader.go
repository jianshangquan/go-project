package filestream

import (
	"io"
	"log"
	"os"
)

type fileStreamReader struct {
	file *os.File
}

func NewFileStreamReader(path string) *fileStreamReader {
	var stream = &fileStreamReader{}
	stream.load(path)
	return stream
}

func (f *fileStreamReader) load(path string) *fileStreamReader {
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	f.file = file
	return f
}

func (f *fileStreamReader) pipe(pipeable PipeableFile, size int, onChunk func(data []byte)) *fileStreamReader {
	var dest = pipeable.getFile()
	var src = f.file

	if size <= 0 {
		size = 1024 * 1024 // Default chunk size of 1MB
	}
	buffer := make([]byte, size)

	for {
		n, err := src.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		if n == 0 {
			break
		}

		var data = buffer[:n]
		_, err = dest.Write(data)

		if onChunk != nil {
			onChunk(data)
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	return f
}

func (f *fileStreamReader) PipeInChunk(pipeable PipeableFile, size int) *fileStreamReader {
	return f.pipe(pipeable, size, nil)
}

func (f *fileStreamReader) PipeInChunkWithCallback(pipeable PipeableFile, size int, onChunk func(data []byte)) *fileStreamReader {
	return f.pipe(pipeable, size, onChunk)
}

func (f *fileStreamReader) Pipe(pipeable PipeableFile) *fileStreamReader {
	return f.pipe(pipeable, -1, nil)
}

func (f *fileStreamReader) Close() {
	// Close the file
	if err := f.file.Close(); err != nil {
		log.Fatal(err)
	}
}
