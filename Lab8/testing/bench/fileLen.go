package bench

import (
	"os"
)

func GetFileLen(f string, bufferSize int) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	size := 0
	for {
		buffer := make([]byte, bufferSize)
		c, err := file.Read(buffer)
		size += c
		if err != nil {
			break
		}
	}
	return size, nil
}
