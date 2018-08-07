package log

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// 文件
// @todo
// 1. 支持批量写入
// 2. 支持按照size切
type RotateFileHandler struct {
	sync.WaitGroup
	closed      int32
	buffers     sync.Pool
	file        *RotateFile
	cfg         RotateFileHandlerConfig
	writeBuffer chan *bytes.Buffer
}

type RotateFileHandlerConfig struct {
	BaseName   string
	When       string
	Interval   int
	BufferSize uint16
	LineBreak  bool

	Flags LFlag
	Level LogLevel
	Fh    FileHeaderFunc
}

func NewRotateFileHandler(cfg RotateFileHandlerConfig) (LoggerHandler, error) {
	f, err := NewRotateFile(cfg.BaseName, cfg.When, cfg.Interval)
	if err != nil {
		return nil, err
	}

	h := &RotateFileHandler{
		file: f,
		cfg:  cfg,
	}

	h.buffers.New = func() interface{} {
		return new(bytes.Buffer)
	}
	h.writeBuffer = make(chan *bytes.Buffer, cfg.BufferSize)

	go h.keepWrite()
	return h, err
}

func (self *RotateFileHandler) Clone() (LoggerHandler, error) {
	h2 := &RotateFileHandler{
		file:        self.file,
		cfg:         self.cfg,
		writeBuffer: make(chan *bytes.Buffer, self.cfg.BufferSize),
	}
	h2.buffers.New = func() interface{} {
		return new(bytes.Buffer)
	}
	return h2, nil
}

// 不可重复调用，非线程安全
func (self *RotateFileHandler) Close() error {
	atomic.CompareAndSwapInt32(&self.closed, 0, 1)
	close(self.writeBuffer)
	self.Wait()
	return self.file.Close()
}

func (self *RotateFileHandler) Flush() error {
	return self.file.Flush()
}

func (self *RotateFileHandler) SetFileHeader(fh FileHeaderFunc) {
	self.cfg.Fh = fh
}

func (self *RotateFileHandler) Logf(level LogLevel, format string, v ...interface{}) error {
	if !self.cfg.Level.Contains(level) {
		return ErrInvalidLevel
	}

	buf := self.getBuffer()

	if self.cfg.Flags&LFlagDate > 0 {
		buf.WriteString(time.Now().Format(LogTimeFormat))
	}

	if self.cfg.Flags&LFlagLevel > 0 {
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString(strings.ToUpper(levelMap[level]))
	}

	if self.cfg.Flags&LFlagFile > 0 {
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString(self.cfg.Fh())
	}

	if buf.Len() > 0 {
		buf.WriteByte(' ')
	}

	buf.WriteString(fmt.Sprintf(format, v...))
	if self.cfg.LineBreak {
		buf.WriteByte('\n')
	}

	timer := time.NewTimer(1 * time.Second)
	self.Add(1)
	select {
	case self.writeBuffer <- buf:
	case <-timer.C:
		self.Done()
		if os.Stderr != nil {
			os.Stderr.WriteString("file_handler.go : write log timeout")
		}
	}

	return nil
}

func (self *RotateFileHandler) keepWrite() {
	for {
		if atomic.LoadInt32(&self.closed) == 1 {
			break
		}
		ch := make(chan struct{}, 0)
		go self.write(ch)
		<-ch
		time.Sleep(1 * time.Second)
	}
}

func (self *RotateFileHandler) write(ch chan struct{}) {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 1<<18)
			n := runtime.Stack(buf, false)
			str := fmt.Sprintf("%v, STACK: %s", r, buf[0:n])

			if os.Stderr != nil {
				os.Stderr.WriteString(str)
			}
		}
	}()
	defer func() {
		ch <- struct{}{}
	}()

	for bb := range self.writeBuffer {
		_, err := self.file.Write(bb.Bytes())
		if err != nil {
			fmt.Println(err)
		}
		self.Done()
		self.putBuffer(bb)
	}
}

func (self *RotateFileHandler) getBuffer() *bytes.Buffer {
	return self.buffers.Get().(*bytes.Buffer)
}

func (self *RotateFileHandler) putBuffer(buf *bytes.Buffer) {
	buf.Reset()
	self.buffers.Put(buf)
}
