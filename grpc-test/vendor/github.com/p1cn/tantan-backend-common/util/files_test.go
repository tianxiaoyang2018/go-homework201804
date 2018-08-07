package util

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteToFile(t *testing.T) {
	data := []byte("abcdefg")
	file, err := ioutil.TempFile(os.TempDir(), "tantan")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(file.Name())
	if err := WriteToFile(file.Name(), data, true); err != nil {
		t.Error(err)
	}
	if err := WriteToFile(file.Name(), data, false); err != os.ErrExist {
		t.Error(err)
	}
}
