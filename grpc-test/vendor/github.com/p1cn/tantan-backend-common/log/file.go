package log

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	ErrInvalidFD = errors.New("fd is nil")
)

// RotateFile : 非线程安全
// 现在仅支持按照时间切分。
type RotateFile struct {
	fd *os.File

	pFileName  string
	baseName   string
	interval   int64
	suffix     string
	rolloverAt int64
	writeData  int64
}

func NewRotateFile(baseName string, when string, interval int) (*RotateFile, error) {

	h := new(RotateFile)

	h.baseName = baseName

	switch when {
	case "second":
		h.interval = 1
		h.suffix = "2006-01-02_15-04-05"
	case "minute":
		h.interval = 60
		h.suffix = "2006-01-02_15-04"
	case "hour":
		h.interval = 3600
		h.suffix = "2006-01-02_15"
	case "day":
		h.interval = 3600 * 24
		h.suffix = "2006-01-02"
	default:
		return nil, fmt.Errorf("invalid when_rotate: %d", when)
	}

	h.interval = h.interval * int64(interval)
	fn := h.fileName()

	var err error
	h.fd, err = os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	h.pFileName = fn

	h.reSetRollover()

	return h, nil
}

func (h *RotateFile) fileName() string {
	if strings.Contains(h.baseName, "%s") {
		return fmt.Sprintf(h.baseName, time.Now().Format(h.suffix))
	}
	return h.baseName + time.Now().Format(h.suffix)
}

func (h *RotateFile) reSetRollover() {
	h.rolloverAt = time.Now().Unix() + h.interval
}

func (h *RotateFile) doRollover() (err error) {
	if h.rolloverAt <= time.Now().Unix() {
		if h.writeData > 0 {
			fName := h.fileName()
			h.fd.Close()

			h.fd, err = os.OpenFile(fName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				return err
			}
			h.pFileName = fName
		}
		h.reSetRollover()
		h.writeData = 0
	}

	return nil
}

// it is not thread safe
func (h *RotateFile) Write(b []byte) (n int, err error) {
	if h.fd == nil {
		return 0, ErrInvalidFD
	}
	h.doRollover()
	n, err = h.fd.Write(b)
	h.writeData += int64(n)
	return n, err
}

func (h *RotateFile) Close() error {
	return h.fd.Close()
}

func (h *RotateFile) Flush() error {
	return h.fd.Sync()
}
