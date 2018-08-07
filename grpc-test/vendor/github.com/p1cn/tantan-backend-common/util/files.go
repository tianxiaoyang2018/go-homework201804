package util

import (
	"io"
	"os"
)

func WriteToFile(filePath string, data []byte, force bool) error {
	if !force {
		if _, err := os.Stat(filePath); !os.IsNotExist(err) {
			return os.ErrExist
		}
	}
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil && err != io.ErrUnexpectedEOF {
		return err
	}
	return nil
}
