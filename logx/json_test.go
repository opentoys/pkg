package logx_test

import (
	"easyadmin/pkg/github/logx"
	"log/slog"
	"os"
	"testing"
)

func TestXxx(t *testing.T) {
	log := slog.New(logx.New(os.Stdout, logx.WithErrorWriter(os.Stderr), logx.WithLevel(slog.LevelDebug)))

	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")
}
