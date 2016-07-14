package logit_test

import (
	"github.com/brettallred/go-logit"
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
