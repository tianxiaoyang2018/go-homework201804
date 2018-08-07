package log

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func testListFile(t *testing.T, dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	var fs []string
	for _, f := range files {
		fs = append(fs, f.Name())
	}
	return fs
}

func TestRotateFileHandler(t *testing.T) {

	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}

	baseName := fmt.Sprintf("%s/test-%s.log", tempDir, "%s")

	hh, err := NewRotateFileHandler(RotateFileHandlerConfig{
		BaseName: baseName,
		When:     "second",
		Interval: 2,

		Flags: LFlagDate | LFlagFile | LFlagLevel,
		Level: LevelInfo,
		Fh:    fileHeader(2),
	})
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		err = hh.Logf(LevelInfo, "test : %d", i)
		if err != nil {
			t.Fatal(err)
		}
		time.Sleep(320 * time.Millisecond)
	}

	fs := testListFile(t, tempDir)
	if len(fs) != 2 {
		t.Fatal("error", fs)
	}
	fmt.Println(fs)
}
