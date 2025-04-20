package filestream

import "os"

type PipeableFile interface {
	getFile() *os.File
	Close()
}
