package filestream

import "time"

func WaitOnChunk(second int) func(data []byte) {
	return func(data []byte) {
		time.Sleep(time.Duration(second) * time.Second)
	}
}
