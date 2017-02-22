package logit_test

import (
	"github.com/brettallred/go-logit"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOutputTypes(t *testing.T) {
	os.Setenv("LOG_LEVEL", "DEBUG")

	logit.Debug(" == %s", "DEBUG")
	logit.Info(" == %s", "INFO")
	logit.Warn(" == %s", "WARN")
	logit.Error(" == %s", "ERROR")
	logit.Fatal(" == %s", "FATAL")
}

func TestChangeLogLevelByName(t *testing.T) {
	logit.SetLogLevel(logit.DEBUG)
	assert.Equal(t, 0, logit.LogLevel())
	logit.Warn("Should print")
	logit.SetLogLevel(logit.INFO)
	assert.Equal(t, 1, logit.LogLevel())
	logit.Warn("Should print")
	logit.SetLogLevel(logit.WARN)
	assert.Equal(t, 2, logit.LogLevel())
	logit.Warn("Should print")
	logit.SetLogLevel(logit.ERROR)
	assert.Equal(t, 3, logit.LogLevel())
	logit.Warn("Should not print")
	logit.SetLogLevel(logit.FATAL)
	assert.Equal(t, 4, logit.LogLevel())
	logit.Warn("Should not print")
}

func TestChangeLogLevel(t *testing.T) {
	logit.SetLogLevel(0)
	assert.Equal(t, 0, logit.LogLevel())
	logit.Warn("Should print")
	logit.SetLogLevel(1)
	assert.Equal(t, 1, logit.LogLevel())
	logit.Warn("Should print")
	logit.SetLogLevel(2)
	assert.Equal(t, 2, logit.LogLevel())
	logit.Warn("Should print")
	logit.SetLogLevel(3)
	assert.Equal(t, 3, logit.LogLevel())
	logit.Warn("Should not print")
	logit.SetLogLevel(4)
	assert.Equal(t, 4, logit.LogLevel())
	logit.Warn("Should not print")
}
