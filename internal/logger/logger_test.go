package logger

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"
)

func TestStdLogger(t *testing.T) {
	logger := NewStdLogger(false, false, false, INFO)

	flags := logger.logger.Flags()
	if flags != 0 {
		t.Fatalf("Expected %q, received %q\n", 0, flags)
	}

	if logger.logLevel != INFO {
		t.Fatalf("Expected %d, received %d\n", INFO, logger.logLevel)
	}
}

func TestStdLoggerWithTime(t *testing.T) {
	logger := NewStdLogger(true, false, false, DEBUG)

	flags := logger.logger.Flags()
	if flags != log.LstdFlags|log.Lmicroseconds {
		t.Fatalf("Expected %d, received %d\n", log.LstdFlags, flags)
	}
}

func TestStdLoggerInfo(t *testing.T) {
	expectOutput(t, func() {
		logger := NewStdLogger(false, false, false, INFO)
		logger.Logf(INFO, "foo")
	}, "[INF] foo\n")
}

func TestStdLoggerInfoWithColor(t *testing.T) {
	expectOutput(t, func() {
		logger := NewStdLogger(false, true, false, INFO)
		logger.Logf(INFO, "foo")
	}, "[\x1b[32mINF\x1b[0m] foo\n")
}

func TestStdLoggerDebug(t *testing.T) {
	expectOutput(t, func() {
		logger := NewStdLogger(false, false, false, DEBUG)
		logger.Logf(DEBUG, "foo %s", "bar")
	}, "[DBG] foo bar\n")
}

func TestStdLoggerDebugWithINFO(t *testing.T) {
	expectOutput(t, func() {
		logger := NewStdLogger(false, false, false, INFO)
		logger.Logf(DEBUG, "foo")
	}, "")
}

func expectOutput(t *testing.T, f func(), expected string) {
	old := os.Stderr // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stderr = w

	f()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	os.Stderr.Close()
	os.Stderr = old // restoring the real stdout
	out := <-outC
	if out != expected {
		t.Fatalf("Expected '%s', received '%s'\n", expected, out)
	}
}
