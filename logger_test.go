package lvlogger

import (
	"bytes"
	"log"
	"testing"
)

func TestNewLvLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	prefix := "[prefix]"
	logger := NewLvLogger(buf, prefix, log.LstdFlags|log.Lshortfile, Debug)

	if logger.level != Debug {
		t.Errorf("Don't match level. lvlogger: %d, expected: %d", logger.level, Debug)
	}

	if logger.inner.Prefix() != prefix {
		t.Errorf("Don't match prefix. inner logger: %s, expected: %s", logger.inner.Prefix(), prefix)
	}

	if logger.inner.Flags() != log.LstdFlags|log.Lshortfile {
		t.Errorf("Don't match flags. inner logger: %d, expected: %d", logger.inner.Flags(), log.LstdFlags|log.Lshortfile)
	}
}

func TestOut(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := NewLvLogger(buf, "", 0, Debug)
	logger.Debug("debug")
	if buf.String() != "debug\n" {
		t.Errorf("Don't match msg. lvlogger: %s, expected: %s", buf.String(), "debug")
	}
}

func TestLevel(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := NewLvLogger(buf, "", 0, Info)
	logger.Debug("debug")
	if buf.String() != "" {
		t.Errorf("Don't match msg. lvlogger: %s, expected: blank", buf.String())
	}
	logger.Info("info")
	if buf.String() != "info\n" {
		t.Errorf("Don't match msg. lvlogger: %s, expected: %s", buf.String(), "info")
	}
}
